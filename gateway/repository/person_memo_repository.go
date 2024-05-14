package repository

import (
	"github.com/go-test-naka/go-place/model"
)

var persons = []model.Person{}

type PersonMemoRepository struct{}

func CreatePersonMemoRepository() *PersonMemoRepository {
	return &PersonMemoRepository{}
}

func (pmr *PersonMemoRepository) Save(person model.Person) error {
	persons = append(persons, person)
	return nil
}

func (pmr *PersonMemoRepository) FindAll() []model.Person {
	return persons
}

func (pmr *PersonMemoRepository) FindByName(name string) *model.Person {
	for _, p := range persons {
		if p.Name == name {
			return &p
		}
	}
	return nil
}
