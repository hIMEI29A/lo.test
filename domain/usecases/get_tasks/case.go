package get_tasks

import (
	"errors"
	"lo.test/domain"
	"lo.test/domain/model"
)

type Request struct {
	Status string
}

type Response struct {
	Tasks []*model.Task
}

func Run(repositories domain.Repositories, r *Request) (*Response, error) {
	if err := validate(repositories, r); err != nil {
		return nil, err
	}

	tasks, err := repositories.InMemory().GetAll(r.Status)
	if err != nil {
		return nil, err
	}

	return &Response{Tasks: tasks}, nil
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

	if r.Status != "" {
		if err := model.ValidateStatus(r.Status); err != nil {
			return err
		}
	}

	return nil
}
