package biz

import (
	v1 "bibirt-api/api/user/v1"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrCreateAnonymousUserFailed = errors.InternalServer(v1.ErrorCode_CREATE_ANONYMOUS_USER_FAILED.String(), "create anonymous user failed")
	ErrTokenInvalid              = errors.BadRequest(v1.ErrorCode_TOKEN_INVALID.String(), "token invalid")
	ErrUserNotFound              = errors.NotFound(v1.ErrorCode_USER_NOT_FOUND.String(), "user not found")
	ErrRefreshTokenInvalid       = errors.BadRequest(v1.ErrorCode_REFRESH_TOKEN_INVALID.String(), "refresh token invalid")
	ErrRefreshTokenMissMatch     = errors.BadRequest(v1.ErrorCode_REFRESH_TOKEN_MISS_MATCH.String(), "refresh token miss match")
	ErrGenerateTokenFailed       = errors.InternalServer(v1.ErrorCode_GENERATE_TOKEN_FAILED.String(), "generate token failed")
)

const (
	USER_TYPE_TEMP   = 1
	USER_TYPE_WECHAT = 2

	USER_STATUS_PENDING_VALID = 1
	USER_STATUS_PENDING_TMP   = 2
	USER_STATUS_ENABLE        = 3
)

type User struct {
	Id        int64  `json:"id"`
	Uuid      string `json:"uuid"`
	Type      int    `json:"type"`
	Name      string `json:"name"`
	Pwd       sql.NullString
	Salt      sql.NullString
	Email     sql.NullString `json:"email,omitempty"`
	Phone     sql.NullString `json:"phone,omitempty"`
	Status    int            `json:"status,omitempty"`
	CreatedAt int64          `json:"created_at,omitempty"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	FindByUuid(context.Context, string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*User, error) {
	user, err := uc.repo.Save(ctx, u)
	if err != nil {
		uJson, _ := json.Marshal(u)
		uc.log.WithContext(ctx).Errorf("CreateUser: %v error: \n%v", string(uJson), err)
	}
	return user, nil
}

func (uc *UserUseCase) FindUserByUuid(ctx context.Context, uuid string) (*User, bool) {
	user, err := uc.repo.FindByUuid(ctx, uuid)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("FindUserByUuid: %v error: \n%v", uuid, err)
		return nil, false
	}
	return user, true
}
