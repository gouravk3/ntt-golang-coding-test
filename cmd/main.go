package main

import (
	"log"

	"github.com/gouravk3/ntt-golang-coding-test/config"
	"github.com/gouravk3/ntt-golang-coding-test/internal/server"
)

func main() {
	appConfig, err := config.Init("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	server.New(appConfig).Start()
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
