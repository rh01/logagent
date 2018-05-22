package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "trace":
		return logs.LevelTrace
	default:
		return logs.LevelDebug
	}
}


func InitLogger() (err error){


	config := make(map[string]interface{})
	config["filename"] = AppConfig.LogPath
	config["level"] = convertLogLevel(AppConfig.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}