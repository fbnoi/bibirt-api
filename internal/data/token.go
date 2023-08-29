package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
)

const (
	TOKEN_BLACK_LIST_NAMESPACE = "black_list:token"
)

type TokenRepo struct {
	data *Data
	log  *log.Helper
}

func (tr *TokenRepo) BlockToken(t *jwt.Token) error {
	tr.data.rdb.SAdd()
}
func (tr *TokenRepo) IsTokenBlocked(t *jwt.Token) bool {

}
