package main

import (
	"log"

	"github.com/gouravk3/ntt-golang-coding-test/config"
	"github.com/gouravk3/ntt-golang-coding-test/internal/router"
	"github.com/gouravk3/ntt-golang-coding-test/internal/server"
)

func main() {
	appConfig, err := config.Init()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	ginEngine, err := router.FuelEstimationRouter(appConfig)
	if err != nil {
		log.Fatalf("failed to initialize router: %v", err)
	}
	server.New(appConfig, ginEngine.Handler()).Start()
}
