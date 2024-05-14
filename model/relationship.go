package model

type Relationship struct {
	Id   uint
	Type RelationshipType
}

type RelationshipType int

const (
	PARENT RelationshipType = iota
	SIBLING
	SPOUSE
	PARTNER
	FRIEND
)
