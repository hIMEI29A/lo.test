package api

import (
	"encoding/json"
	"lo.test/domain"
	"lo.test/domain/model"
	"lo.test/domain/usecases/create_task"
	"lo.test/domain/usecases/get_task_by_id"
	"lo.test/domain/usecases/get_tasks"
	"lo.test/presenters"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	Repos   domain.Repositories
	LogChan chan<- string
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	resp, err := get_tasks.Run(h.Repos, &get_tasks.Request{Status: status})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.LogChan <- "ERROR:" + err.Error()

		return
	}

	response := presenters.PresentTasks(resp.Tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	h.LogChan <- "GET /tasks processed"
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errStr := "Invalid task ID"
		http.Error(w, errStr, http.StatusBadRequest)
		h.LogChan <- "ERROR:" + errStr

		return
	}

	resp, err := get_task_by_id.Run(h.Repos, &get_task_by_id.Request{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		h.LogChan <- "ERROR:" + err.Error()

		return
	}

	response := presenters.PresentTask(resp.Task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	h.LogChan <- "GET /tasks/" + idStr + " processed"
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	h.LogChan <- r.URL.String()
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		errStr := "Invalid request body"
		http.Error(w, errStr, http.StatusBadRequest)
		h.LogChan <- "ERROR:" + errStr

		return
	}

	resp, err := create_task.Run(h.Repos, &create_task.Request{Task: &task})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.LogChan <- "ERROR:" + err.Error()

		return
	}

	response := presenters.PresentTask(resp.Task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	h.LogChan <- "POST /tasks processed - created task ID: " + strconv.Itoa(task.Id)
}
