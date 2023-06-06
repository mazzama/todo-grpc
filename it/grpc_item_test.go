package it

import (
	"context"
	"github.com/mazzama/todo-grpc/pkg/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func (s *e2eTestSuite) Test_Integration_CreateItem() {
	ctx := context.Background()

	client, closer := setupServer(ctx, s.port)
	defer closer()

	type expectation struct {
		out *pb.CreateItemResponse
		err error
	}

	tests := []struct {
		name     string
		in       *pb.CreateItemRequest
		expected expectation
		on       func()
	}{
		{
			name: "Given valid request when create item via grpc should return no error",
			in: &pb.CreateItemRequest{
				Name:        "Advanced Mathematics",
				Description: "Integral",
				Notes:       "For friday",
			},
			expected: expectation{
				err: nil,
			},
		},
		{
			name: "Given invalid request when create item via grpc should return invalid argument",
			in: &pb.CreateItemRequest{
				Name:        "",
				Description: "Integral",
				Notes:       "For friday",
			},
			expected: expectation{
				err: status.Error(codes.InvalidArgument, "field should not empty"),
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			res, err := client.CreateItem(ctx, tt.in)
			if err != nil {
				if err.Error() != tt.expected.err.Error() {
					t.Errorf("got %v, want %v", err, tt.expected.err)
				}
				return
			}

			assert.NotNil(t, res)
			assert.NotEqual(t, 0, res.Id)
		})
	}
}

func (s *e2eTestSuite) Test_Integration_UpdateItem() {
	ctx := context.Background()

	client, closer := setupServer(ctx, s.port)
	defer closer()

	type expectation struct {
		out *pb.Item
		err error
	}

	tests := []struct {
		name     string
		in       *pb.UpdateItemRequest
		expected expectation
		on       func()
	}{
		{
			name: "Given invalid request when update item via grpc should return error code 5",
			in: &pb.UpdateItemRequest{
				Name:        "Advanced Mathematics",
				Description: "Integral",
				Notes:       "For friday",
				Status:      pb.Status_TODO,
			},
			expected: expectation{
				err: status.Error(codes.NotFound, "record not found"),
			},
		},
		{
			name: "Given valid request when find item by should updated item and no error",
			in: &pb.UpdateItemRequest{
				Id:          1,
				Name:        "Advanced Mathematics",
				Description: "Integral",
				Notes:       "For friday",
				Status:      pb.Status_IN_PROGRESS,
			},
			expected: expectation{
				out: &pb.Item{Status: pb.Status_IN_PROGRESS},
				err: nil,
			},
			on: func() {
				_, err := client.CreateItem(ctx, &pb.CreateItemRequest{
					Name:        "Mathematics",
					Description: "Desc",
					Notes:       "For later",
				})
				if err != nil {
					s.T().Errorf("error insert %v", err.Error())
				}
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.on != nil {
				tt.on()
			}

			res, err := client.UpdateItem(ctx, tt.in)
			if err != nil {
				if err.Error() != tt.expected.err.Error() {
					t.Errorf("got %v, want %v", err, tt.expected.err)
				}
				return
			}

			assert.NotNil(t, res)
			assert.Equal(t, res.Status, res.Status)
		})
	}
}

func (s *e2eTestSuite) Test_Integration_FindItemById() {
	ctx := context.Background()

	client, closer := setupServer(ctx, s.port)
	defer closer()

	type expectation struct {
		out *pb.Item
		err error
	}

	tests := []struct {
		name     string
		in       *pb.FindItemRequest
		expected expectation
		on       func()
	}{
		{
			name: "Given invalid id request when find item by id should return not found error",
			in:   &pb.FindItemRequest{Id: 1},
			expected: expectation{
				err: status.Error(codes.NotFound, "record not found"),
			},
		},
		{
			name: "Given valid request when find item by id should return item and no error",
			in:   &pb.FindItemRequest{Id: 1},
			expected: expectation{
				out: &pb.Item{Id: 1},
				err: nil,
			},
			on: func() {
				_, err := client.CreateItem(ctx, &pb.CreateItemRequest{
					Name:        "Mathematics",
					Description: "Desc",
					Notes:       "For later",
				})
				if err != nil {
					s.T().Errorf("error insert %v", err.Error())
				}
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.on != nil {
				tt.on()
			}
			res, err := client.FindItemByID(ctx, tt.in)
			if err != nil {
				if err.Error() != tt.expected.err.Error() {
					t.Errorf("got %v, want %v", err, tt.expected.err)
				}
				return
			}

			assert.NotNil(t, res)
			assert.Equal(t, tt.expected.out.Id, res.Id)
		})
	}
}

func (s *e2eTestSuite) Test_Integration_ViewItemList() {
	ctx := context.Background()

	client, closer := setupServer(ctx, s.port)
	defer closer()

	type expectation struct {
		out *pb.ViewItemListResponse
		err error
	}

	var st = pb.Status_TODO

	tests := []struct {
		name     string
		in       *pb.ViewItemListRequest
		expected expectation
		on       func()
	}{
		{
			name: "Given valid request when view list should return list of item and no error",
			in: &pb.ViewItemListRequest{
				Status: &st,
			},
			expected: expectation{
				out: &pb.ViewItemListResponse{
					Total: 2,
				},
				err: nil,
			},
			on: func() {
				var request = []pb.CreateItemRequest{
					{
						Name:        "Mathematics",
						Description: "Desc",
						Notes:       "For later",
					},
					{
						Name:        "Mathematics",
						Description: "Desc",
						Notes:       "For later",
					},
				}

				for _, v := range request {
					_, err := client.CreateItem(ctx, &v)
					if err != nil {
						s.T().Errorf("error insert %v", err.Error())
					}
				}
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.on != nil {
				tt.on()
			}
			res, err := client.ViewItemList(ctx, tt.in)
			if err != nil {
				if err.Error() != tt.expected.err.Error() {
					t.Errorf("got %v, want %v", err, tt.expected.err)
				}
				return
			}

			assert.NotNil(t, res)
			assert.Equal(t, tt.expected.out.Total, res.Total)
		})
	}
}

func (s *e2eTestSuite) Test_Integration_DeleteItem() {
	ctx := context.Background()

	client, closer := setupServer(ctx, s.port)
	defer closer()

	type expectation struct {
		out *pb.Item
		err error
	}

	tests := []struct {
		name     string
		in       *pb.DeleteItemRequest
		expected expectation
		on       func()
	}{
		{
			name: "Given invalid id request when delete item should return not found error",
			in:   &pb.DeleteItemRequest{Id: 1},
			expected: expectation{
				err: status.Error(codes.NotFound, "record not found"),
			},
		},
		{
			name: "Given valid request when delete should return item and no error",
			in:   &pb.DeleteItemRequest{Id: 1},
			expected: expectation{
				out: &pb.Item{Id: 1},
				err: nil,
			},
			on: func() {
				_, err := client.CreateItem(ctx, &pb.CreateItemRequest{
					Name:        "Mathematics",
					Description: "Desc",
					Notes:       "For later",
				})
				if err != nil {
					s.T().Errorf("error insert %v", err.Error())
				}
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.on != nil {
				tt.on()
			}
			res, err := client.DeleteItem(ctx, tt.in)
			if err != nil {
				if err.Error() != tt.expected.err.Error() {
					t.Errorf("got %v, want %v", err, tt.expected.err)
				}
				return
			}

			assert.NotNil(t, res)
		})
	}
}
