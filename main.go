package main

import (
	"os"

	"github.com/shakilbd009/go-items-api/app"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApp()
}
