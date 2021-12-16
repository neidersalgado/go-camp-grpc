package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"

	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/controller"
	"github.com/neidersalgado/go-camp-grpc/cmd/rest_gokit_server/service"
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
	repository := service.NewUserProxy()
	usersController := controller.NewUserController(repository)
	router.HandleFunc("/users/{id}", usersController.GetByID).Methods("GET")
	//router.HandleFunc("/users", usersController.GetAll).Methods("GET")
	router.HandleFunc("/users", usersController.Create).Methods("POST")
	//router.HandleFunc("/users/{id}", usersController.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", usersController.Delete).Methods("DELETE")
}

type Config struct {
	Port         int    `env:"RESTSERVER_PORT" envDefault:"8000"`
	Env          string `env:"RESTSERVER_ENV" envDefault:"TEST"`
	Hosts        string `env:"RESTSERVER_HOSTS" envDefault:"127.0.0.1:"`
	WriteTimeout int    `env:"RESTSERVER_WRITETIMEOUT" envDefault:"15"`
	ReadTimeout  int    `env:"RESTSERVER_WRITETIMEOUT" envDefault:"15"`
}
