// cmd/metanebula/main.go
package main

import (
	"flag"
	"log"
	"os"

	"metanebula/internal/metanebula"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := metanebula.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
