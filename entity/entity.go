package entity

import "github.com/google/uuid"

var NilID = uuid.Nil

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}
