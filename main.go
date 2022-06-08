package main

import (
	app "banking/app"
	"banking/logger"
)

func main() {
	// log.Println("Starting our application...")
	logger.Info("Starting our application...")
	app.Start()
}
