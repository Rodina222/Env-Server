package main

import (
	"flag"
	"fmt"
	"log"

	server "github.com/codescalersinternships/EnvServer-Rodina/internal"
)

func main() {

	var port int

	flag.IntVar(&port, "p", 8080, "port setted by the user to run the app")

	fmt.Println("port detected:", port)

	flag.Parse()

	app, err := server.NewApp(port)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}

}
