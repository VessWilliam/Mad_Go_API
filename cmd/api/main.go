package main

import "log"

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
