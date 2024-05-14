package gateway

import (
	"fmt"

	"github.com/go-test-naka/go-place/gateway/repository"
	"github.com/go-test-naka/go-place/log"
	"github.com/go-test-naka/go-place/model"
)

type PersonGateway struct {
	repository repository.PersonRepository
}

func NewPersonGateway(repository repository.PersonRepository) *PersonGateway {
	return &PersonGateway{repository}
}

func (pg *PersonGateway) Insert(person model.Person) {
	err := pg.repository.Save(person)
	if err != nil {
		log.Error(fmt.Sprintf("Error to save Person: %v", person), err)
	}
}

func (pg *PersonGateway) SearchAll() []*model.Person {
	return pg.repository.FindAll()
}

func (pg *PersonGateway) SearchByName(name string) *model.Person {
	return pg.repository.FindByName(name)
}
