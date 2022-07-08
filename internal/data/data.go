package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"passport/internal/biz"
	"passport/internal/conf"
	accountV1 "passport/proto/api/account/v1"
	userV1 "passport/proto/api/user/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewAccountServiceClient,
	NewUserServiceClient,
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

func NewUserServiceClient(r registry.Discovery) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///thjam.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userV1.NewUserClient(conn)
}

func NewAccountServiceClient(r registry.Discovery) accountV1.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///thjam.account.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := accountV1.NewAccountClient(conn)
	return c
}
