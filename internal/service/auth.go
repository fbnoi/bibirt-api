package service

import (
	"context"
	"time"

	pb "bibirt-api/api/user/v1"
	"bibirt-api/internal/biz"
	"bibirt-api/internal/conf"
	"bibirt-api/pkg/util"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UUID string
}

type AuthService struct {
	pb.UnimplementedAuthServer

	uc *biz.UserUseCase
	tc *biz.TokenUseCase
	c  *conf.Server
}

func NewAuthService(uc *biz.UserUseCase, conf *conf.Server) *AuthService {
	return &AuthService{uc: uc, c: conf}
}

func (s *AuthService) RegisterAsAnonymous(ctx context.Context, req *pb.RegisterAsAnonymousRequest) (*pb.RegisterAsAnonymousReply, error) {

	user := newAnonymousUser()
	err := s.uc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	refreshToken := s.tc.NewRefreshToken(user)
	token := s.tc.NewToken(refreshToken)
	tokenStr, _ := token.SignedString(s.c.)

	return &pb.RegisterAsAnonymousReply{
		Token:        token.SignedString(conf.Auth.Jwt.Secret),
		RefreshToken: refreshToken.SignedString(conf.Auth.Jwt.Secret),
	}, nil
}

func (s *AuthService) WSToken(ctx context.Context, req *pb.WSTokenRequest) (*pb.WSTokenReply, error) {
	tokenStr := req.Token
	token, err := s.tc.ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(Claims)
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	return &pb.RefreshTokenReply{}, nil
}

func (s *AuthService) ParseToken(ctx context.Context, req *pb.ParseTokenRequest) (*pb.ParseTokenReply, error) {
	tokenStr := req.Token
	token, err := s.tc.ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(Claims)

	return &pb.ParseTokenReply{Uuid: claims.UUID}, nil
}

func (s *AuthService) newToken(user *biz.User) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		UUID: user.Uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.conf.String(),
			Subject:   "client_auth",
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
}

func newAnonymousUser() *biz.User {
	uuid4 := uuid.Must(uuid.NewV4())
	return &biz.User{
		Uuid:      uuid4.String(),
		Type:      biz.USER_TYPE_TEMP,
		Name:      util.GetRandomStr(6),
		Status:    biz.USER_STATUS_PENDING_TMP,
		CreatedAt: time.Now(),
	}
}
