package service

import (
	"context"
	"time"

	pb "bibirt-api/api/user/v1"
	"bibirt-api/internal/biz"
	"bibirt-api/internal/conf"
	"bibirt-api/pkg/util"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UUID string
}

type AuthService struct {
	pb.UnimplementedAuthServer

	uc   *biz.UserUseCase
	tc   *biz.TokenUseCase
	conf *conf.Server
}

func NewAuthService(uc *biz.UserUseCase, tc *biz.TokenUseCase, conf *conf.Server) *AuthService {
	return &AuthService{uc: uc, tc: tc, conf: conf}
}

func (s *AuthService) RegisterAsAnonymous(ctx context.Context, req *pb.RegisterAsAnonymousRequest) (*pb.RegisterAsAnonymousReply, error) {
	var err error
	user := newAnonymousUser()
	err = s.uc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	refreshToken := s.tc.NewRefreshToken(user)
	token := s.tc.NewToken(refreshToken)
	return &pb.RegisterAsAnonymousReply{
		Token:        s.tc.SignedString(token),
		RefreshToken: s.tc.SignedString(refreshToken),
	}, nil
}

func (s *AuthService) WSToken(ctx context.Context, req *pb.WSTokenRequest) (*pb.WSTokenReply, error) {
	token, err := s.parseAndValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	wsToken := s.tc.NewWSToken(token)
	return &pb.WSTokenReply{
		Token: s.tc.SignedString(wsToken),
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	refreshToken, err := s.parseAndValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	token := s.tc.NewToken(refreshToken)

	return &pb.RefreshTokenReply{
		Token: s.tc.SignedString(token),
	}, nil
}

func (s *AuthService) ValidateWSToken(ctx context.Context, req *pb.ValidateWSTokenRequest) (*pb.ValidateWSTokenReply, error) {
	token, err := s.parseAndValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(Claims)

	return &pb.ValidateWSTokenReply{Uuid: claims.UUID}, nil
}

func (s *AuthService) parseAndValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := s.tc.ParseToken(tokenStr)
	if err != nil {
		return nil, errors.BadRequest(err.Error(), "token invalid")
	}
	if !s.tc.IsTokenValid(token) {
		return nil, errors.BadRequest("token invalid", "token invalid")
	}
	return token, nil
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
