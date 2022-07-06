package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"passport/internal/biz"
)

type passportRepo struct {
	data *Data
	log  *log.Helper
}

// NewPassportRepo .
func NewPassportRepo(data *Data, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *passportRepo) Save(ctx context.Context, b *biz.Passport) (*biz.Passport, error) {
	return b, nil
}

func (r *passportRepo) Update(ctx context.Context, b *biz.Passport) (*biz.Passport, error) {
	return b, nil
}

func (r *passportRepo) FindByID(context.Context, int64) (*biz.Passport, error) {
	return nil, nil
}

func (r *passportRepo) ListByHello(context.Context, string) ([]*biz.Passport, error) {
	return nil, nil
}

func (r *passportRepo) ListAll(context.Context) ([]*biz.Passport, error) {
	return nil, nil
}
