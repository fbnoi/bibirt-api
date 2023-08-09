package dao

import (
	"bibirt-api/api/front/model"
	"bibirt-api/api/front/model/trait"
	"bibirt-api/api/util"
	"time"

	"github.com/gofrs/uuid"
)

var (
	_sql_new_user = `
	INSERT INTO users (uuid, type, name, status, created_at, updated_at) 
	VALUES (?,?,?,?,?,?)
	`

	_sql_search_user_by_Uuid = `
	SELECT * FROM users limit 1 WHERE uuid = ?
	`
)

func NewTmpUser() *model.User {
	uuid4 := uuid.Must(uuid.NewV4())
	user := &model.User{
		Uuid:   uuid4.String(),
		Type:   model.USER_TYPE_TEMP,
		Name:   util.GenRandomStr(6),
		Status: model.USER_STATUS_PENDING_TMP,
		TimeTrait: &trait.TimeTrait{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}
	ret := db.MustExec(_sql_new_user, user.Uuid, user.Type, user.Name, user.Status, user.CreatedAt, user.UpdatedAt)
	id, _ := ret.LastInsertId()
	user.ID = uint64(id)

	return user
}

func FindUserByUuid(uuid4 string) (*model.User, error) {
	var user *model.User
	err := db.Get(user, _sql_search_user_by_Uuid, uuid4)

	return user, err
}
