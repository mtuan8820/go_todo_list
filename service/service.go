package service

import "mtuan8820/go-todo-list/v2/pkg/models"

type TaskService interface {
	Save(models.Task) models.Task
	FindAll() []models.Task
}

type taskService struct {
	tasks []models.Task
}

func New() TaskService {
	return &taskService{}
}

func (service *taskService) Save(task models.Task) models.Task {
	service.tasks = append(service.tasks, task)
	return task
}

func (service *taskService) FindAll() []models.Task {
	return service.tasks
}
