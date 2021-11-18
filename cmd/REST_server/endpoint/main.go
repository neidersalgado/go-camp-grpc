package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	service "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/services"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	r := mux.NewRouter()

	userService := service.NewUserService(nil)
	CreateAccountHandler := httptransport.NewServer(
		makeCreateUserEndpoint(userService),
		decodeCreateUserRequest,
		encodeResponse,
	)
	GetByCustomerIdHandler := httptransport.NewServer(
		makeGetUserEndpoint(userService),
		decodeGetUserRequest,
		encodeResponse,
	)
	DeleteCustomerHandler := httptransport.NewServer(
		makeDeleteUserEndpoint(userService),
		decodeDeleteUserequest,
		encodeResponse,
	)
	http.Handle("/", r)
	http.Handle("/users", CreateAccountHandler)

	r.Handle("/user/{user_id}", GetByCustomerIdHandler).Methods("GET")
	r.Handle("/user/{user_id}", DeleteCustomerHandler).Methods("DELETE")
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
