package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"

	"github.com/neidersalgado/go-camp-grpc/cmd/grpc_server/config"
	"github.com/neidersalgado/go-camp-grpc/cmd/grpc_server/web"
)

func main() {
	conf := config.Config{}

	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}

	muxRouter := mux.NewRouter()

	web.SetUpRouter(muxRouter)

	server := &http.Server{
		Handler: muxRouter,
		Addr:    conf.Hosts + strconv.Itoa(conf.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
