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
	token := s.tc.NewAuthToken(refreshToken)
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
	if !s.tc.IsAuthToken(token) {
		return nil, pb.ErrorTokenInvalid("token invalid")
	}
	wsToken := s.tc.NewWSToken(token)
	return &pb.WSTokenReply{
		Token: s.tc.SignedString(wsToken),
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	refreshToken, err := s.parseAndValidateToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if !s.tc.IsRefreshToken(refreshToken) {
		return nil, pb.ErrorTokenInvalid("token's subject miss match")
	}
	token := s.tc.NewAuthToken(refreshToken)
	return &pb.RefreshTokenReply{
		Token: s.tc.SignedString(token),
	}, nil
}

func (s AuthService) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	token, err := s.parseAndValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	if !s.tc.IsAuthToken(token) {
		return nil, pb.ErrorTokenInvalid("token's subject miss match")
	}
	claims := s.tc.Claims(token)
	var user biz.User
	if s.uc.FindUserByUuid(ctx, claims.UUID, &user) {
		return &pb.UserInfoReply{
			Uuid:  claims.UUID,
			Name:  user.Name,
			Score: user.Score,
		}, nil
	}
	return nil, pb.ErrorUserNotFound("user not found")
}

func (s *AuthService) WSUserInfo(ctx context.Context, req *pb.WSUserInfoRequest) (*pb.WSUserInfoReply, error) {
	token, err := s.parseAndValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	if !s.tc.IsWSToken(token) {
		return nil, pb.ErrorTokenInvalid("token's subject miss match")
	}
	claims := s.tc.Claims(token)
	var user biz.User
	if s.uc.FindUserByUuid(ctx, claims.UUID, &user) {
		return &pb.WSUserInfoReply{
			Uuid:  user.Uuid,
			Name:  user.Name,
			Score: user.Score,
		}, nil
	}
	return nil, pb.ErrorUserNotFound("user not found")
}

func (s *AuthService) parseAndValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := s.tc.ParseToken(tokenStr)
	if err != nil {
		return nil, pb.ErrorTokenInvalid("parse token error: %s", err)
	}
	if !s.tc.IsTokenValid(token) {
		return nil, pb.ErrorTokenInvalid("token invalid")
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
