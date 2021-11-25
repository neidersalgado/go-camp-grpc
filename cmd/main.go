package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"

	"github.com/neidersalgado/go-camp-grpc/pkg/user"
)

func main() {
	conf := Config{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}

	muxRouter := mux.NewRouter()

	SetUpRouter(muxRouter)

	server := &http.Server{
		Handler:      muxRouter,
		Addr:         conf.Hosts + strconv.Itoa(conf.Port),
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func SetUpRouter(router *mux.Router) {
	repository := user.NewProxyRepository()
	service := user.NewDefaultUserService(repository)
	userHandler := user.NewUserHandler(service)
	router.HandleFunc("/users", userHandler.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.Get).Methods("GET")
	router.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", userHandler.SetUserParents).Methods("PUT")
}

type Config struct {
	Port         int    `env:"RESTSERVER_PORT" envDefault:"8080"`
	Env          string `env:"RESTSERVER_ENV" envDefault:"TEST"`
	Hosts        string `env:"RESTSERVER_HOSTS" envDefault:"127.0.0.1:"`
	WriteTimeout int    `env:"RESTSERVER_WRITETIMEOUT" envDefault:"15"`
	ReadTimeout  int    `env:"RESTSERVER_WRITETIMEOUT" envDefault:"15"`
}
