package main

import (
	"flag"
	"log"

	"github.com/gouravk3/ntt-golang-coding-test/config"
	"github.com/gouravk3/ntt-golang-coding-test/internal/router"
	"github.com/gouravk3/ntt-golang-coding-test/internal/server"
)

var configFilePath = flag.String("configPath", "/home/gourav.prasad_win@npci.org.in/Downloads/ntt-golang-coding-test/config/config.fuelestimation.yaml", "app configurations")

func main() {
	flag.Parse()

	appConfig, err := config.Init(configFilePath)
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	ginEngine, err := router.FuelEstimationRouter(appConfig)
	if err != nil {
		log.Fatalf("failed to initialize router: %v", err)
	}
	server.New(appConfig, ginEngine.Handler()).Start()
}
