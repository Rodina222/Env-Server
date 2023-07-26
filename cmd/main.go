package main

import (
	"flag"
	"log"

	server "github.com/codescalersinternships/EnvServer-Rodina/internal"
)

func main() {

	var port int

	flag.IntVar(&port, "p", 8080, "port setted by the user to run the app")

	flag.Parse()

	app, err := server.NewApp(port)

	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run() ; err!=nil {
    		log.Fatal(err)
	}

}
