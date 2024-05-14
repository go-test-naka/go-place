package repository

import "github.com/go-test-naka/go-place/model"

type PersonRepository interface {
	Save(person model.Person) error
	FindAll() []*model.Person
	FindByName(name string) *model.Person
}
