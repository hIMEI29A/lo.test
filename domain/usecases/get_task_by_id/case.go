package get_task_by_id

import (
	"errors"
	"lo.test/domain"
	"lo.test/domain/model"
)

type Request struct {
	Id int
}

type Response struct {
	Task *model.Task
}

func Run(repositories domain.Repositories, r *Request) (*Response, error) {
	if err := validate(repositories, r); err != nil {
		return nil, err
	}

	t, err := repositories.InMemory().GetById(r.Id)
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

	if r.Id == 0 {
		return errors.New("request.Id is zero")
	}

	return nil
}
