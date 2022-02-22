package service

import (
	"encoding/json"
	"fmt"
	"io"
	"lol-data-cli-tool/logger"
	"net/http"
)

type DataDragonChampionName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DataDragonResponse struct {
	Version string `json:"version"`
	Data    map[string]DataDragonChampionName
}

func FetchChampionNames(v string) DataDragonResponse {
	URL := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion.json", v)
	logger.Info("Fetching champion names for patch: " + v + " ...")

	res, err := http.Get(URL)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Fatal(err.Error())
	}
	var result DataDragonResponse
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Fatal(err.Error())
	}

	if result.Version != v {
		logger.Fatal("Fetched data version doesn't match!")
	}

	logger.Info("Successfully fetched champion names.")
	return result
}
