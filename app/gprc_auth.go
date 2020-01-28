package app

import (
	"context"

	"github.com/muizidn/todo-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type serviceAuth struct {
	pb.UnimplementedAuthServiceServer
	rUser UserRepo
}

func (s *serviceAuth) Register(ctx context.Context, in *pb.RegisterReq) (*pb.User, error) {
	hash, err := passwordGenerateHash(in.GetPassword())
	if err != nil {
		return nil, err
	}
	user, err := s.rUser.Create(in.GetUsername(), in.GetEmail(), hash)
	if err != nil {
		return nil, err
	}
	return &pb.User{Uuid: user.uuid, Username: user.username, Email: user.email}, nil
}
func (s *serviceAuth) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	user, err := s.rUser.GetUsername(in.GetUsername())
	if err != nil {
		return nil, err
	}
	if !passwordValidate(in.GetPassword(), user.hashedPassword) {
		return nil, log.TError(grpc.Errorf(codes.Unauthenticated, "username or password not match"))
	}

	jwtToken, err := jwtGenerateToken(user.uuid)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Token: jwtToken,
		User:  &pb.User{Uuid: user.uuid, Username: user.username, Email: user.email},
	}, nil
}
