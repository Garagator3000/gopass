package main

import (
	"log"
	"os"
)

func main() {
	app := initApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
