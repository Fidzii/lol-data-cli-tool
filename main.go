package main

import (
	"lol-data-cli-tool/logger"
	"lol-data-cli-tool/service"
)

func main() {
	version := "12.4.1"
	data := service.FetchChampionNames(version)
	logger.Info(data)
}
