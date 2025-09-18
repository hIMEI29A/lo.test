package repository

// Предполагаем, что у нас могут быть другие реализации, кроме in_memory.
// Допустим, что это будет БД MySQL
// В Этом случае интерфейс Repos должен быть расширен - в нем добавится метод MySQL() domain.TaskRepository

import (
	"lo.test/domain"
	"lo.test/repository/in_memory"
)

type Repos interface {
	InMemory() domain.TaskRepository
}

type repo struct {
	inMemory domain.TaskRepository
}

func (repo *repo) InMemory() domain.TaskRepository {
	return repo.inMemory
}

func New() domain.Repositories {
	return &repo{
		inMemory: in_memory.New(),
	}
}
