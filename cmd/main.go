package main

import (
	"fmt"
	"log"

	"github.com/gouravk3/ntt-golang-coding-test/config"
)

func main() {
	appConfig, err := config.Init("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	fmt.Printf("appConfig: %v\n", appConfig)

	// // Initialize logger
	// logger, err := lgr.NewLogger(appConfig.GetLoggerConfig().Level)
	// check(err)

	// logger.Info(fmt.Sprintf("Environment : %s", appConfig.GetServerConfig().Environment))

	// // Initialize database connection
	// _, err = db.NewDatabase(appConfig.GetDatabaseConfig(), logger)
	// check(err)
	// logger.Info("Database initialised")

	// // Initialize gin
	// ginEngine, err := router.InitRouter(logger)
	// check(err)
	// logger.Info("ginEngine initialised")

	// logger.Info("Starting server...")
	// done := make(chan os.Signal, 1)
	// signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	err := ginEngine.Run(appConfig.GetServerConfig().Host + ":" + appConfig.GetServerConfig().Port)
	// 	check(err)
	// }()

	// appServices := services.InitServices()
	// logger.Info("services initialised")

	// <-done
	// log.Info("Server has stopped.")
}
