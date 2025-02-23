package main

import (
	"github.com/irbis_helper/internal/api"
	"github.com/irbis_helper/internal/config"
	"github.com/irbis_helper/internal/files"
)

func main() {
	cfg := config.Get()
	fls := files.NewFiles(cfg.ClientPath, cfg.ClientVersion)
	e := api.Initiate(cfg.PassPhrase, fls)
	if err := e.Start(cfg.AppPort); err != nil {
		panic(err)
	}
}
