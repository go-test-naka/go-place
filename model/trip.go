package model

import "time"

type Trip struct {
	Id       uint
	PersonId uint
	PlaceId  uint
	Review   string
	Rating   uint
	Date     time.Time
}
