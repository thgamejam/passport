package biz

import (
	"context"
	"errors"
	accountV1 "passport/proto/api/account/v1"
	userV1 "passport/proto/api/user/v1"
)

type Account struct {
	Email    string
	Password string
	Hash     string
}

type User struct {
	Id       uint32
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

// PassportRepo is a Passport repo.
type PassportRepo interface {
	// PrepareCreateAccount 预创建账户到缓存
	PrepareCreateAccount(ctx context.Context, account Account) error
	// CreatAccount 创建账户
	CreatAccount(ctx context.Context, sid string, key string) (id uint32, err error)
	// SignLoginToken 签署登录token
	SignLoginToken(ctx context.Context, accountID uint32) (token string, err error)
	// GetPublicKey 获取公钥和哈希值
	GetPublicKey(ctx context.Context) (key string, hash string, err error)
	// LoginVerify 登录校验
	LoginVerify(ctx context.Context, username string, ciphertext string, hash string) (id uint32, err error)
	// ChangeUserPassword 修改用户密码
	ChangeUserPassword(ctx context.Context, accountID uint32, password string, hash string) (ok bool, err error)
	// VerifyAccountTokenId 验证用户Token是否合法并返回合法账户ID
	VerifyAccountTokenId(ctx context.Context) (accountId uint32, err error)
	// ChangePassword 修改密码
	ChangePassword(ctx context.Context, id uint32, ciphertext string, hash string) (err error)
	// AccountLogout 注销会话号ID
	AccountLogout(ctx context.Context, id uint32, sid string) (err error)
	// GetUserByAccountID 通过账户ID获取用户
	GetUserByAccountID(ctx context.Context, id uint32) (user *userV1.UserInfo, err error)
	// GetAccountByID 通过账号ID获取账户信息
	GetAccountByID(ctx context.Context, id uint32) (accountInfo *accountV1.GetAccountReply, err error)
	// CreateUserByAccountID 通过账户ID创建用户
	CreateUserByAccountID(ctx context.Context, id uint32) (user *userV1.UserInfo, err error)
}

// Logout 登出
func (uc *PassportUseCase) Logout(ctx context.Context, id uint32, sid string) (err error) {
	err = uc.repo.AccountLogout(ctx, id, sid)
	return
}

// ChangePassword 修改密码
func (uc *PassportUseCase) ChangePassword(ctx context.Context, id uint32, ciphertext string, hash string) (token string, err error) {
	err = uc.repo.ChangePassword(ctx, id, ciphertext, hash)
	if err != nil {
		return "", err
	}
	token, err = uc.repo.SignLoginToken(ctx, id)
	if err != nil {
		return "", err
	}
	return
}

// RenewalToken 验证并续签Token
func (uc *PassportUseCase) RenewalToken(ctx context.Context) (string, error) {
	accountId, err := uc.repo.VerifyAccountTokenId(ctx)
	if err != nil {
		return "", err
	}
	token, err := uc.repo.SignLoginToken(ctx, accountId)
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetKey 获取公钥和哈希
func (uc *PassportUseCase) GetKey(ctx context.Context) (key string, hash string, err error) {
	return uc.repo.GetPublicKey(ctx)
}

// CreatAccountAndUser 验证sid的md5值并创建用户签署登录token
func (uc *PassportUseCase) CreatAccountAndUser(ctx context.Context, sid string, key string) (userInfo *User, err error) {
	// 创建账户
	id, err := uc.repo.CreatAccount(ctx, sid, key)
	if err != nil {
		return nil, err
	}

	// 通过账户ID创建用户
	user, err := uc.repo.CreateUserByAccountID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 签署登录Token
	token, err := uc.repo.SignLoginToken(ctx, id)
	return &User{
		Id:       user.Id,
		Token:    token,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.AvatarUrl,
	}, nil
}

// Login 登录并签署验证token
func (uc *PassportUseCase) Login(ctx context.Context, username string, ciphertext string, hash string) (token string, userInfo *userV1.UserInfo, err error) {
	// 验证密码
	accountID, err := uc.repo.LoginVerify(ctx, username, ciphertext, hash)
	if err != nil {
		return "", nil, err
	}

	// 根据账户ID查找用户是否存在,若不存在则创建用户
	userInfo, err = uc.repo.GetUserByAccountID(ctx, accountID)
	if err != nil {
		if userV1.IsUserNotFoundByAccount(err) {
			userInfo, err = uc.repo.CreateUserByAccountID(ctx, accountID)
			if err == nil {
				return "", nil, err
			}
		}
		return "", nil, errors.New("InternalServiceError")
	}

	loginToken, err := uc.repo.SignLoginToken(ctx, accountID)
	if err != nil {
		return "", nil, err
	}

	return loginToken, userInfo, nil
}

// PrepareCreateAccount 预创建账户
func (uc *PassportUseCase) PrepareCreateAccount(ctx context.Context, account Account, token string) error {
	// TODO 验证人机token
	return uc.repo.PrepareCreateAccount(ctx, account)
}
