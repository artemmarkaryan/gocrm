package server

import (
	"fmt"
	"log"
	"os"
)

const address = "0.0.0.0"

func Run() {
	r := getEngine()
	setRouter(r)
	setMiddleware(r)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("No port declared in env")
	}
	log.Printf("Running on http://%v:%v", address, port)
	err := r.Run(fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		log.Fatal(err.Error())
	}
}
