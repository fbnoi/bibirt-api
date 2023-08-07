package model

import (
	"bibirt-api/api/front/model/trait"
)

type User struct {
	*trait.TimeTrait

	ID uint64
	Uuid string 
	Type uint8
	Phone string
	Password string
	Name string
	Salt string
	Status uint8
}
