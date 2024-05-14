package person

import (
	"time"

	"github.com/go-test-naka/go-place/gateway"
	"github.com/go-test-naka/go-place/log"
	"github.com/go-test-naka/go-place/model"
)

// Aqui está usando um ponteiro de gateway
type Create struct {
	personGateway *gateway.PersonGateway
}

func NewCreate(personGateway *gateway.PersonGateway) *Create {
	return &Create{personGateway}
}

func (cp *Create) Create() {
	log.Info("Create person")
	person := model.Person{
		Id:       1,
		Name:     "Kakaroto",
		Birthday: time.Date(1900, 1, 25, 0, 0, 0, 0, time.UTC),
		Country:  "Brazil",
	}

	cp.personGateway.Insert(person)
}
