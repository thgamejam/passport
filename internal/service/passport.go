package service

import (
	"context"
	"errors"
	"github.com/thgamejam/pkg/jwt"
	"passport/internal/biz"
	pb "passport/proto/api/passport/v1"
)

//Logout 登出请求
func (s *PassportService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	loginToken, ok := jwt.FromLoginTokenContext(ctx)
	if !ok {
		return nil, errors.New("TokenNotFound")
	}
	err := s.uc.Logout(ctx, loginToken.UserID, loginToken.UUID)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutReply{}, nil
}

// CreateAccount 预创建账户
func (s *PassportService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	err := s.uc.PrepareCreateAccount(ctx, biz.Account{
		Email:    req.User.Email,
		Password: req.User.Password.Value,
		Hash:     req.User.Password.Hash,
	}, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountReply{}, nil
}

// VerifyEmail 验证邮箱并返回登录token
func (s *PassportService) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.LoginReply, error) {
	loginInfo, err := s.uc.CreatAccountAndUser(ctx, req.Sid, req.Key)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		User: &pb.LoginReply_User{
			Id:       loginInfo.Id,
			Token:    loginInfo.Token,
			Username: loginInfo.Username,
			Bio:      loginInfo.Bio,
			Image:    loginInfo.Image,
		},
	}, nil
}

// GetPublicKey 获取公钥和哈希值
func (s *PassportService) GetPublicKey(ctx context.Context, req *pb.GetPublicKeyRequest) (*pb.GetPublicKeyReply, error) {
	k, h, err := s.uc.GetKey(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetPublicKeyReply{
		Key: &pb.GetPublicKeyReply_PublicKey{
			Hash:  h,
			Value: k,
		},
	}, nil
}

// Login 登录
func (s *PassportService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, user, err := s.uc.Login(ctx, req.User.Email, req.User.Password.Value, req.User.Password.Hash)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		User: &pb.LoginReply_User{
			Id:       user.Id,
			Token:    token,
			Username: user.Username,
			Bio:      user.Bio,
			Image:    user.AvatarUrl,
		},
	}, nil
}

// ChangePassword 修改密码
//func (s *PassportService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.LoginReply, error) {
//	loginToken, ok := jwt.FromLoginTokenContext(ctx)
//	if !ok {
//		return nil, errors.New("TokenNotFound")
//	}
//
//	_, err := s.uc.ChangePassword(ctx, loginToken.UserID, req.NewPassword.Value, req.NewPassword.Hash)
//	if err != nil {
//		return nil, err
//	}
//
//	token, u, err := s.uc.Login(ctx, loginToken.UserID, req.NewPassword.Value, req.NewPassword.Hash)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.LoginReply{
//		User: &pb.LoginReply_User{
//			Id:       u.Id,
//			Token:    token,
//			Username: u.Username,
//			Bio:      u.Bio,
//			Image:    u.AvatarUrl,
//		},
//	}, nil
//}
