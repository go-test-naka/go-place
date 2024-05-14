package person

import (
	"fmt"

	"github.com/go-test-naka/go-place/gateway"
	"github.com/go-test-naka/go-place/log"
)

// Aqui não está usando um ponteiro de gateway
type Search struct {
	personGateway gateway.PersonGateway
}

func NewSearch(personGateway gateway.PersonGateway) *Search {
	return &Search{personGateway}
}

func (s *Search) SearchAll() {
	log.Info("Find all persons")
	persons := s.personGateway.SearchAll()
	log.Info(fmt.Sprintf("Persons quantity: %v", len(persons)))
	for _, v := range persons {
		log.Info(fmt.Sprintf("Person: %v", v))
	}
}

func (s *Search) SearchById(id uint) {
	persons := s.personGateway.SearchAll()
	for _, v := range persons {
		if v.Id == id {
			log.Info(fmt.Sprintf("Person: %v", v))
			break
		}
	}
}

func (s *Search) SearchByName(name string) {
	person := s.personGateway.SearchByName(name)
	if person != nil {
		log.Info(fmt.Sprintf("Person found: %v", person))
		return
	}
	log.Info("Person not found")
}
