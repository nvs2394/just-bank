package main

import (
	"github.com/nvs2394/just-bank/app"
	"github.com/nvs2394/just-bank/logger"
)

func main() {
	logger.Info("Starting just bank service")
	app.Start()
}
