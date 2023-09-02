package biz

import (
	"bibirt-api/internal/conf"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TOKEN_SUBJECT         = "auth_token"
	WS_TOKEN_SUBJECT      = "ws_token"
	REFRESH_TOKEN_SUBJECT = "refresh_token"
)

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
	RegisterToken(t *jwt.Token)
	BlockToken(t *jwt.Token)
	BlockTokenByUUIDAndSubject(uuid, subject string)
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

func (tr *TokenUseCase) NewWSToken(token *jwt.Token) *jwt.Token {
	return tr.newToken(
		token.Claims.(*Claims).UUID,
		WS_TOKEN_SUBJECT,
		time.Now().Add(tr.cj.WsPeriod.AsDuration()),
	)
}

func (tr *TokenUseCase) NewToken(refreshToken *jwt.Token) *jwt.Token {
	return tr.newToken(
		refreshToken.Claims.(*Claims).UUID,
		TOKEN_SUBJECT,
		time.Now().Add(tr.cj.WsPeriod.AsDuration()),
	)
}

func (tr *TokenUseCase) NewRefreshToken(u *User) *jwt.Token {
	return tr.newToken(
		u.Uuid,
		REFRESH_TOKEN_SUBJECT,
		time.Now().Add(tr.cj.RefreshPeriod.AsDuration()),
	)
}

func (tr *TokenUseCase) BlockToken(tok *jwt.Token) {
	if !tr.repo.IsTokenBlocked(tok) {
		tr.repo.BlockToken(tok)
	}
}

func (tr *TokenUseCase) IsTokenValid(tok *jwt.Token) bool {
	return tok.Valid && !tr.repo.IsTokenBlocked(tok)
}

func (tr *TokenUseCase) ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return tr.cj.Secret, nil
	})
}

func (tr *TokenUseCase) SignedString(token *jwt.Token) string {
	str, err := token.SignedString(tr.cj.Secret)
	if err != nil {
		panic(err)
	}
	return str
}

func (tr *TokenUseCase) newToken(UUID, subject string, expiresAt time.Time) *jwt.Token {
	now := time.Now()
	uuid4 := uuid.Must(uuid.NewV4())
	tok := jwt.NewWithClaims(encrypt_methods[tr.cj.EncryptMethod], Claims{
		UUID: UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tr.issuer,
			Subject:   subject,
			Audience:  []string{UUID},
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			ID:        uuid4.String(),
		},
	})
	tr.repo.BlockTokenByUUIDAndSubject(UUID, subject)
	tr.repo.RegisterToken(tok)
	return tok
}
