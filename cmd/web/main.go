package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type appConfig struct {
	useCache bool
}

type application struct {
	config appConfig
}

func main() {
	app := application{}

	server := &http.Server{
		Addr:              ":3456",
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           app.routes(),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("http://localhost:3456")
	fmt.Println("Press <ENTER> to quit")
	fmt.Scanln()
}
