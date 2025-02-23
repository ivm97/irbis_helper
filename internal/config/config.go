package config

import (
	"encoding/json"
	"os"
)

type Params struct {
	AppPort       string `json:"app_port"`
	PassPhrase    string `json:"pass_phrase"`
	ClientPath    string `json:"client_path"`
	ClientVersion string `json:"client_version"`
}

func Get() *Params {
	b, err := os.ReadFile("settings/settings.json")
	if err != nil {
		panic(err)
	}

	var p Params
	err = json.Unmarshal(b, &p)
	if err != nil {
		panic(err)
	}
	return &p
}
