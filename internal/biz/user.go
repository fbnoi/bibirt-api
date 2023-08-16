package biz

import (
	v1 "bibirt-api/api/user/v1"
	"context"
	"database/sql"
	"encoding/json"
	"time"

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
	Type      uint8  `json:"type"`
	Name      string `json:"name"`
	Pwd       sql.NullString
	Salt      sql.NullString
	Email     sql.NullString `json:"email,omitempty"`
	Phone     sql.NullString `json:"phone,omitempty"`
	Status    uint8          `json:"status,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
}

type UserRepo interface {
	Save(context.Context, *User) error
	FindByUUID(context.Context, string, *User) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUseCase new a Greeter usecase.
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) error {
	err := uc.repo.Save(ctx, u)
	if err != nil {
		uJson, _ := json.Marshal(u)
		uc.log.WithContext(ctx).Errorf("CreateUser: %v error: \n%v", string(uJson), err)
	}
	return nil
}

func (uc *UserUseCase) FindUserByUuid(ctx context.Context, uuid string, dist *User) bool {
	err := uc.repo.FindByUUID(ctx, uuid, dist)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("FindUserByUuid: %v error: \n%v", uuid, err)
		return false
	}
	return true
}
