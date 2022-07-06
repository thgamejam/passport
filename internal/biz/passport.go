package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Passport is a Passport model.
type Passport struct {
	Hello string
}

// PassportUseCase is a Passport use case.
type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

// NewPassportUseCase new a Passport use case.
func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreatePassport creates a Passport, and returns the new Passport.
func (uc *PassportUseCase) CreatePassport(ctx context.Context, g *Passport) (*Passport, error) {
	uc.log.WithContext(ctx).Infof("CreatePassport: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
