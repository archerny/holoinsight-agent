/*
 * Copyright 2022 Holoinsight Project Authors. Licensed under Apache-2.0.
 */

package main

import (
	"encoding/json"
	"github.com/traas-stack/holoinsight-agent/pkg/appconfig"
	"github.com/traas-stack/holoinsight-agent/pkg/collecttask"
	"github.com/traas-stack/holoinsight-agent/pkg/pipeline/standard"
	_ "github.com/traas-stack/holoinsight-agent/pkg/plugin/output/console"
	"github.com/traas-stack/holoinsight-agent/pkg/util"
	"os"
	"time"
)

func main() {
	appconfig.SetDev(true)
	//data, err := os.ReadFile("cmd/plugintest/telegraf_redis.json")
	data, err := os.ReadFile("cmd/plugintest/telegraf_nginx.json")
	if err != nil {
		panic(err)
	}
	task := &collecttask.CollectTask{}
	if err := json.Unmarshal(data, task); err != nil {
		panic(err)
	}
	task.Config.Content = util.ToJsonBytes(task.Config.ContentObj)

	pipeline, err := standard.ParsePipeline(task)
	if err != nil {
		panic(err)
	}

	pipeline.Start()
	time.Sleep(time.Hour)
}
