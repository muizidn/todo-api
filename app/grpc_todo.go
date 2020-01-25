package app

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/muizidn/todo-api/pb"
)

type serviceTodo struct {
	pb.UnimplementedTodoServiceServer
	repo TodoRepo
}

func (s *serviceTodo) List(ctx context.Context, in *empty.Empty) (*pb.TodoList, error) {
	todos, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	list := &pb.TodoList{}
	for _, v := range todos {
		list.Todos = append(list.Todos, &pb.Todo{
			Id:          &pb.TodoIdentifier{Uuid: v.uuid},
			Title:       v.title,
			Description: v.description,
			CreatedAt: &timestamp.Timestamp{
				Seconds: v.createdAt,
			},
		})
	}
	return list, nil
}
func (s *serviceTodo) Get(ctx context.Context, in *pb.TodoIdentifier) (*pb.Todo, error) {
	todo, err := s.repo.Get(in.GetUuid())
	if err != nil {
		return nil, err
	}
	return &pb.Todo{
		Id:          &pb.TodoIdentifier{Uuid: todo.uuid},
		Title:       todo.title,
		Description: todo.description,
		CreatedAt: &timestamp.Timestamp{
			Seconds: todo.createdAt,
		},
	}, nil
}
func (s *serviceTodo) Create(ctx context.Context, in *pb.TodoCreate) (*pb.Todo, error) {
	todo, err := s.repo.Create(in.GetTitle(), in.GetDescription())
	if err != nil {
		return nil, err
	}
	return &pb.Todo{
		Id:          &pb.TodoIdentifier{Uuid: todo.uuid},
		Title:       todo.title,
		Description: todo.description,
		CreatedAt: &timestamp.Timestamp{
			Seconds: todo.createdAt,
		},
	}, nil
}
func (s *serviceTodo) Update(ctx context.Context, in *pb.Todo) (*pb.Todo, error) {
	todo, err := s.repo.Update(Todo{
		uuid:        in.GetId().GetUuid(),
		title:       in.GetTitle(),
		description: in.GetDescription(),
		createdAt:   in.GetCreatedAt().GetSeconds(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Todo{
		Id:          &pb.TodoIdentifier{Uuid: todo.uuid},
		Title:       todo.title,
		Description: todo.description,
		CreatedAt: &timestamp.Timestamp{
			Seconds: todo.createdAt,
		},
	}, nil
}
func (s *serviceTodo) Delete(ctx context.Context, in *pb.TodoIdentifier) (*pb.Status, error) {
	if err := s.repo.Delete(in.GetUuid()); err != nil {
		return nil, err
	}
	return &pb.Status{Code: 0, Message: "Success"}, nil
}
