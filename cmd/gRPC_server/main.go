package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/neidersalgado/go-camp-grpc/cmd/gRPC_server/pb"
	"github.com/neidersalgado/go-camp-grpc/cmd/gRPC_server/repository"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db := dbConn()
	defer db.Close()
	userRepo := repository.NewMySQLUserRepository(db)
	user := pb.UserRequest{
		Id:                    "134",
		Name:                  "Juan Bedoya",
		PwdHash:               "fgdfgerFGDrWErwerWErWE435RFW",
		Age:                   23,
		AdditionalInformation: "none",
	}

	//err := userRepo.Create(user)
	//err := userRepo.Delete(user.Id)
	userResponse, err := userRepo.Get(user.Id)
	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println(fmt.Errorf("error in save user", err))
	}
	fmt.Println(userResponse.String())
	// be careful deferring Queries if you are using transactions
	defer db.Close()

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "secret"
	dbName := "users"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:33060)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
