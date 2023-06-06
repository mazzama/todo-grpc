package handler

import (
	custom_error "github.com/mazzama/todo-grpc/pkg/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func WrapError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return status.Error(codes.NotFound, err.Error())
	case custom_error.ErrInvalidStatus, custom_error.ErrEmptyField:
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
