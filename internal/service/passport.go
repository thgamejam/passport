package service

import (
	"context"
	
	"passport/internal/biz"
	v1 "passport/proto/api/passport/v1"
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

// SayHello implements helloworld.PassportServer.
func (s *PassportService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreatePassport(ctx, &biz.Passport{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
