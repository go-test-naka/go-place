package model

import "time"

type Person struct {
	Id       uint
	Name     string
	Birthday time.Time
	Country  string
}
