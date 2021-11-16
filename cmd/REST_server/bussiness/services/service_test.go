package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	service "github.com/neidersalgado/go-camp-grpc/cmd/REST_server/bussiness/services"
)

func Test_NewServiceFunction(t *testing.T) {
	userService := service.NewUserService(nil)

	assert.IsType(t, service.DefaultUserService{}, userService)
}
