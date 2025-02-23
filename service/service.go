package service

import (
	"errors"
	"mtuan8820/go-todo-list/v2/pkg/models"
)

type TaskService interface {
	Add(models.Task) models.Task
	GetAll() []models.Task
	Delete(int) error
	// Edit(string, int) error
	Toggle(int) error
}

type taskService struct {
	tasks  []models.Task
	nextID int
}

func New() TaskService {
	return &taskService{nextID: 1}
}

func (service *taskService) Add(task models.Task) models.Task {
	task.ID = service.nextID
	service.nextID++
	service.tasks = append(service.tasks, task)
	return task
}

func (service *taskService) GetAll() []models.Task {
	return service.tasks
}

func (service *taskService) Delete(id int) error {
	for index, task := range service.tasks {
		if task.ID == id {
			service.tasks = append(service.tasks[:index], service.tasks[index+1:]...)
			return nil
		}
	}
	return errors.New("delete error, task with given ID not found")

}

func (service *taskService) Edit(title string, id int) error {
	for index, task := range service.tasks {
		if task.ID == id {
			service.tasks[index].Title = title
			return nil
		}
	}
	return errors.New("edit error, task with given ID not found")
}

func (service *taskService) Toggle(id int) error {
	for index, task := range service.tasks {
		if task.ID == id {
			service.tasks[index].Completed = !service.tasks[index].Completed
			return nil
		}
	}
	return errors.New("toggle error, task with given ID not found")
}
