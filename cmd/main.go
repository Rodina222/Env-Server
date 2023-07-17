package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/EnvServer-Rodina/internal"
)

func main() {
	var port int

	flag.IntVar(&port, "p", 8080, "port setted by the to run the app")

	flag.Parse()

	app, err := internal.CreateApp(port)
	if err != nil {
		log.Fatal(err)
	}

}
