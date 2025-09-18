package presenters

import "lo.test/domain/model"

type TaskResponse struct {
	Id          int    `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func PresentTask(task *model.Task) TaskResponse {
	return TaskResponse{
		Id:          task.Id,
		Status:      string(task.Status),
		Description: task.Description,
	}
}

func PresentTasks(tasks []*model.Task) []TaskResponse {
	result := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		result[i] = PresentTask(task)
	}
	return result
}
