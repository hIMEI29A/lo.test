package domain

import (
	"lo.test/domain/model"
)

// TaskRepository - общий интерфейс хранилища
type TaskRepository interface {
	GetAll(status string) ([]*model.Task, error)
	GetById(id int) (*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
}

// Repositories включает в себя все реализации интерфейса TaskRepository как in-memory, так и другие, основанные на различных БД
type Repositories interface {
	InMemory() TaskRepository
}
