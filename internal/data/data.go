package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"passport/internal/biz"

	"passport/internal/conf"
	accountV1 "passport/proto/api/account/v1"
	userV1 "passport/proto/api/user/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewPassportRepo,
)

// Data .
type Data struct {
	accountClient accountV1.AccountClient
	userClient    userV1.UserClient
	conf          *conf.Passport
}

// NewData .
func NewData(
	accountClient accountV1.AccountClient,
	userClient userV1.UserClient,
	c *conf.Passport,
	logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		accountClient: accountClient,
		userClient:    userClient,
		conf:          c,
	}, cleanup, nil
}

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
