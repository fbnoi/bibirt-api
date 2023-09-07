package data

import (
	"bibirt-api/internal/biz"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TOKEN_BLACK_LIST_NAMESPACE = "token:black_list:"
	TOKEN_USER_NAMESPACE       = "token:user"
)

type ErrAble interface {
	Err() error
}

type TokenRepo struct {
	rdb *redis.Client
}

func NewTokenRepo(data *Data, logger log.Logger) biz.TokenRepo {
	return &TokenRepo{
		rdb: data.rdb,
	}
}

func (tr *TokenRepo) RegisterToken(tok *jwt.Token) {
	claims := tok.Claims.(jwt.RegisteredClaims)
	UUID := claims.Audience[0]
	tr.must(
		tr.rdb.HSet(
			fmt.Sprintf("%s:%s", TOKEN_USER_NAMESPACE, UUID),
			claims.Subject,
			claims.ID,
		),
	)
	key := fmt.Sprintf("token:%s:%s", claims.Subject, claims.ID)
	tr.must(tr.rdb.Set(key, 0, time.Until(claims.ExpiresAt.Time)))
}

func (tr *TokenRepo) BlockToken(tok *jwt.Token) {
	claims := tok.Claims.(jwt.RegisteredClaims)
	key := fmt.Sprintf("token:%s:%s", claims.Subject, claims.ID)
	tr.must(tr.rdb.Set(key, 1, 0))
}

func (tr *TokenRepo) BlockTokenByUUIDAndSubject(uuid, subject string) {
	id := tr.getUserTokenID(uuid, subject)
	key := fmt.Sprintf("token:%s:%s", subject, id)
	tr.must(tr.rdb.Set(key, 1, 0))
}

func (tr *TokenRepo) getUserTokenID(UUID, field string) string {
	key := fmt.Sprintf("%s:%s", TOKEN_USER_NAMESPACE, UUID)
	id, err := tr.rdb.HGet(key, field).Result()
	if err != nil {
		panic(err)
	}

	return id
}

func (tr *TokenRepo) IsTokenBlocked(tok *jwt.Token) bool {
	claims := tok.Claims.(jwt.RegisteredClaims)
	key := fmt.Sprintf("token:%s:%s", claims.Subject, claims.ID)
	res, err := tr.rdb.Exists(key).Result()
	if err != nil {
		panic(err)
	}
	if res <= 0 {
		return false
	}
	i, err := tr.rdb.Get(key).Int()
	if err != nil {
		panic(err)
	}
	return i == 1
}

func (tr *TokenRepo) must(errAble ErrAble) {
	if err := errAble.Err(); err != nil {
		panic(err)
	}
}
