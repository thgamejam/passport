package service

import (
	"github.com/google/wire"
	"passport/internal/biz"
	v1 "passport/proto/api/passport/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewPassportService,
)

// PassportService is a passport service.
type PassportService struct {
	v1.UnimplementedPassportServer

	uc *biz.PassportUseCase
}

// NewPassportService new a passport service.
func NewPassportService(uc *biz.PassportUseCase) *PassportService {
	return &PassportService{uc: uc}
}
