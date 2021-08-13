package main

import (
	"fmt"
	"log"
	"net-catcher/app"
	"os"
)

func main() {
	fmt.Printf("Get info about website\n")

	application := app.Catch()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
