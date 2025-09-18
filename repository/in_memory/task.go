package in_memory

import (
	"errors"
	"lo.test/domain/model"
	"sync"
)

type Task struct {
	Id          int
	Status      string
	Description string
}

func NewFromModel(task *model.Task) *Task {
	return &Task{
		Status:      string(task.Status),
		Description: task.Description,
	}
}

func (t *Task) ToModel() *model.Task {
	return &model.Task{
		Id:          t.Id,
		Status:      model.Status(t.Status),
		Description: t.Description,
	}
}

type InMemoryTaskRepository struct {
	tasks  map[int]*Task
	mu     sync.RWMutex
	nextId int
}

func New() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]*Task),
		nextId: 1,
	}
}

func (r *InMemoryTaskRepository) GetAll(status string) ([]*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*model.Task

	for _, task := range r.tasks {
		if task.Status == status || status == "" {
			result = append(result, task.ToModel())
		}
	}
	return result, nil
}

func (r *InMemoryTaskRepository) GetById(id int) (*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task.ToModel(), nil
}

func (r *InMemoryTaskRepository) Create(task *model.Task) (*model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t := NewFromModel(task)
	t.Id = r.nextId

	r.tasks[r.nextId] = t
	r.nextId++

	return t.ToModel(), nil
}
