package create_task

import (
	"errors"
	"lo.test/domain"
	"lo.test/domain/model"
)

type Request struct {
	Task *model.Task
}

type Response struct {
	Task *model.Task
}

func Run(repositories domain.Repositories, r *Request) (*Response, error) {
	if err := validate(repositories, r); err != nil {
		return nil, err
	}

	t, err := repositories.InMemory().Create(r.Task)
	if err != nil {
		return nil, err
	}

	return &Response{Task: t}, nil
}

func validate(repositories domain.Repositories, r *Request) error {
	if repositories == nil {
		return errors.New("repositories is nil")
	}

	if repositories.InMemory() == nil {
		return errors.New("in-memory repository is nil")
	}

	if r == nil {
		return errors.New("request is nil")
	}

	if r.Task == nil {
		return errors.New("request.Task is nil")
	}

	if r.Task.Status == "" {
		return errors.New("Task.Status is empty")
	}

	if err := model.ValidateStatus(string(r.Task.Status)); err != nil {
		return err
	}

	if r.Task.Description == "" {
		return errors.New("Task.Description is empty")
	}

	return nil
}
