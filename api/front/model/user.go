package model

import (
	"bibirt-api/api/front/model/trait"
	"database/sql"
)

const (
	USER_TYPE_TEMP   = uint8(1)
	USER_TYPE_WECHAT = uint8(2)

	USER_STATUS_PENDING_VALID = uint8(1)
	USER_STATUS_PENDING_TMP   = uint8(2)
	USER_STATUS_ENABLE        = uint8(3)
)

type User struct {
	*trait.TimeTrait

	ID       uint64
	Uuid     string
	Type     uint8
	Phone    sql.NullString
	Password sql.NullString
	Name     string
	Salt     sql.NullString
	Status   uint8
}
