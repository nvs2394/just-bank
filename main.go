package main

import (
	"github.com/nvs2394/just-bank-lib/logger"
	"github.com/nvs2394/just-bank/app"
)

func main() {
	logger.Info("Starting just bank service")
	app.Start()
}
