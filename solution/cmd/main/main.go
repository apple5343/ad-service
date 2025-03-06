package main

import "server/internal/app"

func main() {
	app := app.NewApp()
	app.Run()
}
