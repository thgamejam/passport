package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPassportUseCase)

// PassportUseCase is a Passport use case.
type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

// NewPassportUseCase new a Passport use case.
func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{repo: repo, log: log.NewHelper(logger)}
}
