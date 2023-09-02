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

	TOKEN_BLOCKED_FIELD = "blocked"
)

type ErrAble interface {
	Err() error
}

type TokenRepo struct {
	rdb *redis.Client
	log *log.Helper
}

func NewTokenRepo(data *Data, logger log.Logger) biz.TokenRepo {
	return &TokenRepo{
		rdb: data.rdb,
		log: log.NewHelper(logger),
	}
}

func (tr *TokenRepo) RegisterToken(tok *jwt.Token) {
	claims := tok.Claims.(jwt.RegisteredClaims)
	UUID := claims.Audience[0]
	tr.logErr(
		tr.rdb.HSet(
			fmt.Sprintf("%s:%s", TOKEN_USER_NAMESPACE, UUID),
			claims.Subject,
			claims.ID,
		),
	)
	key := fmt.Sprintf("token:%s:%s", claims.Subject, claims.ID)
	tr.logErr(tr.rdb.Set(key, 0, time.Until(claims.ExpiresAt.Time)))
}

func (tr *TokenRepo) BlockToken(tok *jwt.Token) {
	claims := tok.Claims.(jwt.RegisteredClaims)
	key := fmt.Sprintf("token:%s:%s", claims.Subject, claims.ID)
	tr.logErr(tr.rdb.Set(key, 1, 0))
}

func (tr *TokenRepo) BlockTokenByUUIDAndSubject(uuid, subject string) {
	if id, err := tr.getUserTokenID(uuid, subject); err == nil {
		key := fmt.Sprintf("token:%s:%s", subject, id)
		tr.logErr(tr.rdb.Set(key, 1, 0))
	} else {
		tr.log.Error(err)
	}
}

func (tr *TokenRepo) getUserTokenID(UUID, field string) (string, error) {
	key := fmt.Sprintf("%s:%s", TOKEN_USER_NAMESPACE, UUID)
	return tr.rdb.HGet(key, field).Result()
}

func (tr *TokenRepo) IsTokenBlocked(tok *jwt.Token) bool {
	claims := tok.Claims.(jwt.RegisteredClaims)
	key := fmt.Sprintf("token:%s:%s", claims.Subject, claims.ID)
	if res, err := tr.rdb.Exists(key).Result(); err == nil && res > 0 {
		if i, err := tr.rdb.Get(key).Int(); err == nil {
			return i == 1
		} else {
			tr.log.Error(err)
		}
	} else {
		tr.log.Error(err)
	}
	return true
}

func (tr *TokenRepo) logErr(errAble ErrAble) {
	if err := errAble.Err(); err != nil {
		tr.log.Error(err)
	}
}
