package biz

import (
	"bibirt-api/internal/conf"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TOKEN_SUBJECT = "auth_token"
const REFRESH_TOKEN_SUBJECT = "refresh_token"

var (
	encrypt_methods = map[conf.Auth_Jwt_EncryptMethods]jwt.SigningMethod{
		conf.Auth_Jwt_ES256: jwt.SigningMethodES256,
		conf.Auth_Jwt_ES384: jwt.SigningMethodES384,
		conf.Auth_Jwt_ES512: jwt.SigningMethodES512,
		conf.Auth_Jwt_RS256: jwt.SigningMethodRS256,
		conf.Auth_Jwt_RS384: jwt.SigningMethodRS384,
		conf.Auth_Jwt_RS512: jwt.SigningMethodRS512,
		conf.Auth_Jwt_HS256: jwt.SigningMethodHS256,
		conf.Auth_Jwt_HS384: jwt.SigningMethodHS384,
		conf.Auth_Jwt_HS512: jwt.SigningMethodHS512,
		conf.Auth_Jwt_PS256: jwt.SigningMethodPS256,
		conf.Auth_Jwt_PS384: jwt.SigningMethodPS384,
		conf.Auth_Jwt_PS512: jwt.SigningMethodPS512,
		conf.Auth_Jwt_EdDSA: jwt.SigningMethodEdDSA,
	}
)

type Claims struct {
	jwt.RegisteredClaims
	UUID string `json:"uid,omitempty"`
}

type TokenRepo interface {
	BlockToken(t *jwt.Token) error
	IsTokenBlocked(t *jwt.Token) bool
}

type TokenUseCase struct {
	repo   TokenRepo
	issuer string
	cj     *conf.Auth_Jwt
}

func NewTokenUseCase(ce *conf.Endpoint, ca *conf.Auth, repo TokenRepo) *TokenUseCase {
	return &TokenUseCase{repo, ce.Id, ca.Jwt}
}

func (tr *TokenUseCase) NewToken(refreshToken *jwt.Token) *jwt.Token {
	now := time.Now()
	expiresAt := now.Add(tr.cj.Period.AsDuration())

	return jwt.NewWithClaims(encrypt_methods[tr.cj.EncryptMethod], Claims{
		UUID: refreshToken.Claims.(*Claims).UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tr.issuer,
			Subject:   TOKEN_SUBJECT,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
}
func (tr *TokenUseCase) NewRefreshToken(u *User) *jwt.Token {
	now := time.Now()
	expiresAt := now.Add(tr.cj.RefreshPeriod.AsDuration())
	return jwt.NewWithClaims(encrypt_methods[tr.cj.EncryptMethod], Claims{
		UUID: u.Uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tr.issuer,
			Subject:   REFRESH_TOKEN_SUBJECT,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
}
func (tr *TokenUseCase) BlockToken(tok *jwt.Token) {
	t, _ := tok.Claims.GetExpirationTime()
	if t.After(time.Now()) && !tr.repo.IsTokenBlocked(tok) {
		tr.repo.BlockToken(tok)
	}
}

func (tr *TokenUseCase) ParseToken(tokenStr string) (*jwt.Token, error)
