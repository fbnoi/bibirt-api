package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	pb "bibirt-api/api/user/v1"
	"bibirt-api/internal/biz"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	uc *biz.UserUseCase
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) RegisterAsAnonymous(ctx context.Context, req *pb.RegisterAsAnonymousRequest) (*pb.RegisterAsAnonymousReply, error) {
	user := newAnonymousUser()
	user, err := s.uc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterAsAnonymousReply{}, nil
}
func (s *AuthService) LoginFromToken(ctx context.Context, req *pb.LoginFromTokenRequest) (*pb.LoginFromTokenReply, error) {
	return &pb.LoginFromTokenReply{}, nil
}
func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	return &pb.RefreshTokenReply{}, nil
}

func newAnonymousUser() *biz.User {
	uuid4 := uuid.Must(uuid.NewV4())
	return &biz.User{
		Uuid:      uuid4.String(),
		Type:      biz.USER_TYPE_TEMP,
		Name:      genRandomStr(6),
		Status:    biz.USER_STATUS_PENDING_TMP,
		CreatedAt: time.Now().Unix(),
	}
}

func newToken(user *biz.User) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		Uid: user.Uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "client_auth",
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
}

func genRandomStr(len int) string {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
