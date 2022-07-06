package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPassportUseCase)

// PassportRepo is a Passport repo.
type PassportRepo interface {
	Save(context.Context, *Passport) (*Passport, error)
	Update(context.Context, *Passport) (*Passport, error)
	FindByID(context.Context, int64) (*Passport, error)
	ListByHello(context.Context, string) ([]*Passport, error)
	ListAll(context.Context) ([]*Passport, error)
}
