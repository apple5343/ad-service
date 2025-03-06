package main

import "server/internal/app"

func main() {
	app := app.NewTestApp()
	app.Run()
}
