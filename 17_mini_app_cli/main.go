package main

import (
	"fmt"
	"log"
	"mini-cli/app"
	"os"
)

func main() {
	fmt.Println("App Start")

	application := app.Create()

	appError := application.Run(os.Args)
	if appError != nil {
		log.Fatal(appError)
	}
}
