package biz

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	UUID string
}

type TokenRepo interface {
	BlockToken(t *jwt.Token) error
	IsTokenBlocked(t *jwt.Token) bool
}

type TokenUseCase struct {
	repo *TokenRepo
}

func NewTokenUseCase(repo *TokenRepo) *TokenUseCase {
	return &TokenUseCase{repo}
}

func (tr *TokenUseCase) NewToken(u *User, refreshToken *jwt.Token) (*jwt.Token, error)
func (tr *TokenUseCase) NewRefreshToken(u *User) (*jwt.Token, error)
func (tr *TokenUseCase) ExpireToken(*jwt.Token)
