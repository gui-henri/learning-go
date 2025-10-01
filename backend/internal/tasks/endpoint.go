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
		task, err := svc.InsertTask(ctx, req.Descricao, req.Prazo, req.PacienteID)
		if err != nil {
			return nil, err
		}
		return InsertTaskResponse{Task: task}, nil
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

func makeGetAllEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		tasks, err := svc.GetAllTasks(ctx)
		if err != nil {
			return nil, err
		}
		return GetAllTasksResponse{Tasks: tasks}, nil
	}
}

func makeDeleteTaskEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(DeleteTaskRequest)
		err = svc.DeleteTasks(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
