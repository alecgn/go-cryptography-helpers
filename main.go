package main

import (
	"go-cryptography-helpers/app"
	"log"
	"os"
)

func main() {
	application := app.Create()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}