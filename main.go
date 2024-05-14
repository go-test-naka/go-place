package main

import (
	"github.com/go-test-naka/go-place/config/mysql"
	"github.com/go-test-naka/go-place/controller"
	"github.com/go-test-naka/go-place/gateway"
	"github.com/go-test-naka/go-place/gateway/repository"
	"github.com/go-test-naka/go-place/log"
	"github.com/go-test-naka/go-place/usecase/person"
)

func main() {
	log.Info("Starting app")

	//rPerson := repository.CreatePersonMemoRepository()

	dbConnection := mysql.CreateConnection()
	defer dbConnection.Close()

	rPerson := repository.CreatePersonMySQLRepository(dbConnection)
	gPerson := gateway.NewPersonGateway(rPerson)
	uCreatePerson := person.NewCreate(gPerson)
	uSearchPerson := person.NewSearch(*gPerson)
	c := controller.CreateNewController(*uCreatePerson, *uSearchPerson)
	c.Execute()
}
