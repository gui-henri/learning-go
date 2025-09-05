package tasks

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetTaskEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(GetTaskRequest)
		t, err := svc.GetTask(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return GetTaskResponse{Task: t}, nil
	}
}

func makeInsertTaskEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(InsertTaskRequest)
		id, err := svc.InsertTask(ctx, req.Descricao, req.Prazo)
		if err != nil {
			return nil, err
		}
		return InsertTaskResponse{Id: id}, nil
	}
}

func makeFinishTaskEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(FinishTaskRequest)
		t, err := svc.FinishTask(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return FinishTaskResponse{Task: t}, nil
	}
}

func makeGetAllIncompleteEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		tasks, err := svc.GetAllIncompleteTasks(ctx)
		if err != nil {
			return nil, err
		}
		return GetAllIncompleteTasksResponse{Tasks: tasks}, nil
	}
}
