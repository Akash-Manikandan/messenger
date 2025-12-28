package user

import (
	"context"

	"github.com/Akash-Manikandan/app-backend/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	UnimplementedUserServiceServer
	service Service
}

func NewGRPCServer(s Service) *GRPCServer {
	return &GRPCServer{service: s}
}

func (g *GRPCServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := g.service.CreateUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return toProtoUser(user), nil
}

func (g *GRPCServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	user, err := g.service.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return toProtoUser(user), nil
}

func (g *GRPCServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	user, err := g.service.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// Update fields if provided
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Bio != nil {
		user.Bio = *req.Bio
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}
	if req.IsVerified != nil {
		user.IsVerified = *req.IsVerified
	}

	if err := g.service.UpdateUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	return toProtoUser(user), nil
}

func (g *GRPCServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	if err := g.service.DeleteUser(req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	return &DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}

func (g *GRPCServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	limit := int(req.Limit)
	offset := int(req.Offset)

	if limit <= 0 {
		limit = 10
	}

	users, err := g.service.ListUsers(limit, offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	protoUsers := make([]*UserResponse, len(users))
	for i, user := range users {
		protoUsers[i] = toProtoUser(&user)
	}

	return &ListUsersResponse{
		Users:  protoUsers,
		Limit:  req.Limit,
		Offset: req.Offset,
		Total:  int32(len(users)),
	}, nil
}

func (g *GRPCServer) VerifyUser(ctx context.Context, req *VerifyUserRequest) (*VerifyUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	if err := g.service.VerifyUser(req.UserId); err != nil {
		if err.Error() == ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, ErrUserNotFound)
		}
		if err.Error() == ErrUserAlreadyVerified {
			return nil, status.Errorf(codes.AlreadyExists, ErrUserAlreadyVerified)
		}
		return nil, status.Errorf(codes.Internal, "failed to verify user: %v", err)
	}

	// Get updated user
	user, err := g.service.GetUserByID(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user: %v", err)
	}

	return &VerifyUserResponse{
		Success: true,
		Message: "User verified successfully",
		User:    toProtoUser(user),
	}, nil
}

// toProtoUser converts a models.User to a UserResponse proto message
func toProtoUser(user *models.User) *UserResponse {
	return &UserResponse{
		Id:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Avatar:     user.Avatar,
		Bio:        user.Bio,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:  user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
