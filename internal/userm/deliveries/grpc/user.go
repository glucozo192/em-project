package deliveries

import (
	"context"
	"fmt"

	"github.com/glu/shopvui/idl/pb"
	"github.com/glu/shopvui/internal/userm/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userDelivery struct {
	pb.UnimplementedUserServiceServer
	userService services.UserService
}

func NewUserDelivery(userService services.UserService) pb.UserServiceServer {
	return &userDelivery{
		userService: userService,
	}
}

func (d *userDelivery) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "the request body data cant be nil!")
	}

	if _, err := d.userService.Login(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("unable to create user: %v", err).Error())
	}

	return &pb.LoginResponse{}, nil
}

func (d *userDelivery) GetMe(ctx context.Context, req *pb.GetMeRequest) (*pb.GetMeResponse, error) {

	return &pb.GetMeResponse{}, nil
}

func (d *userDelivery) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	return &pb.RegisterResponse{}, nil
}
