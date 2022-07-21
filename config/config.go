package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/edrank/edrank_backend/types"
	"github.com/edrank/edrank_backend/utils"
)

func LoadConfig() *types.DatabaseConfig {
	var mainConfig *types.DatabaseConfig
	jsonFile, err := ioutil.ReadFile("./config/database.json")
	if err != nil {
		utils.PrintToConsole(err.Error(), "error")
		return nil
	}
	err = json.Unmarshal([]byte(jsonFile), &mainConfig)

	if err != nil {
		utils.PrintToConsole(err.Error(), "error")
		return nil
	}
	return mainConfig
}
