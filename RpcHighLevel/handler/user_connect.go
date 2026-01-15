package handler

import (
	"context"
	"errors"
	userv1 "highlevel/proto/user/v1"
	"highlevel/proto/user/v1/userv1connect"
	"highlevel/service"

	"connectrpc.com/connect"
)

// UserServiceHandler implements the Connect RPC UserService
type UserServiceHandler struct {
	userv1connect.UnimplementedUserServiceHandler
}

// NewUserServiceHandler creates a new UserServiceHandler
func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{}
}

// GetUser retrieves user information by user ID
func (h *UserServiceHandler) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	userID := req.Msg.UserId
	if userID == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("user_id is required"),
		)
	}

	// Use service to get user info
	userInfo, err := service.GetUserService(userID)
	if err != nil {
		connectCode := connect.CodeNotFound
		if err.Error() != "user not found in database" {
			connectCode = connect.CodeInternal
		}
		return nil, connect.NewError(connectCode, err)
	}

	// Convert to Connect response
	response := &userv1.GetUserResponse{
		UserId:  userInfo.UserID,
		EmailId: userInfo.EmailID,
		Name:    userInfo.Name,
		Role:    userInfo.Role,
	}

	return connect.NewResponse(response), nil
}
