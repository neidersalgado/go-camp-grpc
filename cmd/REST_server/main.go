package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"

	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/config"
	"github.com/neidersalgado/go-camp-grpc/cmd/REST_server/endpoint"
)

func main() {
	conf := config.Config{}

	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}

	muxRouter := mux.NewRouter()

	endpoint.SetUpRouter(muxRouter)

	server := &http.Server{
		Handler:      muxRouter,
		Addr:         conf.Hosts + strconv.Itoa(conf.Port),
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
