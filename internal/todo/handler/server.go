package handler

import (
	"context"
	"github.com/mazzama/todo-grpc/internal/todo/model"
	"github.com/mazzama/todo-grpc/internal/todo/service"
	"github.com/mazzama/todo-grpc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppTodoServer struct {
	pb.UnimplementedTodoServiceServer
	service service.ItemService
}

func NewTodoServerGrpc(grpcServer *grpc.Server, itemService service.ItemService) {
	todoServer := &AppTodoServer{service: itemService}

	pb.RegisterTodoServiceServer(grpcServer, todoServer)
}

func (s *AppTodoServer) CreateItem(ctx context.Context, in *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	response, err := s.service.Create(ctx, model.CreateItemRequest{
		Name:        in.Name,
		Description: in.Description,
		Notes:       in.Notes,
	})
	if err != nil {
		return nil, WrapError(err)
	}
	return &pb.CreateItemResponse{
		Id: response.ID,
	}, nil
}

func (s *AppTodoServer) FindItemByID(ctx context.Context, in *pb.FindItemRequest) (*pb.Item, error) {
	item, err := s.service.FindItemById(ctx, in.Id)
	if err != nil {
		return nil, WrapError(err)
	}

	return &pb.Item{
		Id:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Notes:       item.Notes,
		Status:      pb.Status_TODO,
		CreatedAt:   timestamppb.New(item.CreatedAt),
		UpdatedAt:   timestamppb.New(item.UpdatedAt),
	}, nil
}

func (s *AppTodoServer) UpdateItem(ctx context.Context, in *pb.UpdateItemRequest) (*pb.Item, error) {
	item, err := s.service.Update(ctx, model.UpdateItemRequest{
		ID:          in.Id,
		Name:        in.Name,
		Description: in.Description,
		Notes:       in.Notes,
		Status:      in.Status.String(),
	})
	if err != nil {
		return nil, WrapError(err)
	}

	return &pb.Item{
		Id:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Notes:       item.Notes,
		Status:      pb.Status(pb.Status_value[item.Status]),
		CreatedAt:   timestamppb.New(item.CreatedAt),
		UpdatedAt:   timestamppb.New(item.UpdatedAt),
	}, nil
}

func (s *AppTodoServer) ViewItemList(ctx context.Context, req *pb.ViewItemListRequest) (*pb.ViewItemListResponse, error) {
	var reqStatus = ""
	if req.Status != nil {
		reqStatus = req.Status.String()
	}

	items, err := s.service.ViewItemList(ctx, model.ViewItemListRequest{
		Status: reqStatus,
	})
	if err != nil {
		return nil, WrapError(err)
	}

	var itemListResponse []*pb.Item
	for _, v := range items.Items {
		item := pb.Item{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Notes:       v.Notes,
			Status:      pb.Status(pb.Status_value[v.Status]),
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}

		itemListResponse = append(itemListResponse, &item)
	}

	return &pb.ViewItemListResponse{
		Total: uint32(items.Total),
		Items: itemListResponse,
	}, nil
}

func (s *AppTodoServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	err := s.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, WrapError(err)
	}

	return &pb.DeleteItemResponse{}, nil
}
