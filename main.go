package main

import (
	"fmt"
	"io"
	"lol-data-cli-tool/logger"
	"net/http"
)

func fetchChampionNames(v string) {
	URL := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion.json", v)

	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	// version := "12.4.1"
	// fetchChampionNames(version)
	// c := color.New(color.FgCyan).Add(color.Underline)
	// test := c.Sprint("[prefix]")
	// InfoLogger = log.New(os.Stderr, test+" ", log.Lmsgprefix)
	// InfoLogger.Println(version)
	logger.Info("info")
	logger.Warning("warning")
	logger.Error("error")
}
