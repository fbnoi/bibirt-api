package data

import (
	"bibirt-api/internal/biz"
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

const (
	TOKEN_BLACK_LIST_NAMESPACE = "token:black_list:"
	TOKEN_USER_NAMESPACE       = "token:user"
)

type MyClaimsInterface interface {
	jwt.Claims
	GetID() (string, error)
	GetUUID() (string, error)
}

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
	claim := tok.Claims.(MyClaimsInterface)
	tr.must(
		tr.rdb.HSet(
			context.Background(),
			fmt.Sprintf("%s:%s", TOKEN_USER_NAMESPACE, claimUUID(claim)),
			claimSubject(claim),
			claimID(claim),
		),
	)
	key := fmt.Sprintf("token:%s:%s", claimSubject(claim), claimID(claim))
	tr.must(tr.rdb.Set(context.Background(), key, 0, time.Until(claimExpireAt(claim))))
}

func (tr *TokenRepo) BlockToken(tok *jwt.Token) {
	claim := tok.Claims.(MyClaimsInterface)
	key := fmt.Sprintf("token:%s:%s", claimSubject(claim), claimID(claim))
	tr.must(tr.rdb.Set(context.Background(), key, 1, 0))
}

func (tr *TokenRepo) BlockTokenByUUIDAndSubject(uuid, subject string) {
	id := tr.getUserTokenID(uuid, subject)
	key := fmt.Sprintf("token:%s:%s", subject, id)
	tr.must(tr.rdb.Set(context.Background(), key, 1, 0))
}

func (tr *TokenRepo) getUserTokenID(UUID, field string) string {
	key := fmt.Sprintf("%s:%s", TOKEN_USER_NAMESPACE, UUID)
	id, err := tr.rdb.HGet(context.Background(), key, field).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}

	return id
}

func (tr *TokenRepo) IsTokenBlocked(tok *jwt.Token) bool {
	claim := tok.Claims.(MyClaimsInterface)
	key := fmt.Sprintf("token:%s:%s", claimSubject(claim), claimID(claim))
	res, err := tr.rdb.Exists(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	if res <= 0 {
		return false
	}
	i, err := tr.rdb.Get(context.Background(), key).Int()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return i == 1
}

func (tr *TokenRepo) must(errAble ErrAble) {
	if err := errAble.Err(); err != nil && err != redis.Nil {
		panic(err)
	}
}

func claimUUID(claim MyClaimsInterface) string {
	str, _ := claim.GetUUID()
	return str
}

func claimID(claim MyClaimsInterface) string {
	str, _ := claim.GetID()
	return str
}

func claimSubject(claim MyClaimsInterface) string {
	str, _ := claim.GetSubject()
	return str
}

func claimExpireAt(claim MyClaimsInterface) time.Time {
	t, _ := claim.GetExpirationTime()
	return t.Time
}
