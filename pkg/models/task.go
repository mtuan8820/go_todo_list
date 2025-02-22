package models

import "errors"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoList []Task

func (t *TodoList) Add(title string, id int) {
	newTask := Task{ID: id, Title: title, Completed: false}
	*t = append(*t, newTask)
}

func (t *TodoList) Remove(id int) error {
	for index, task := range *t {
		if task.ID == id {
			*t = append((*t)[:index], (*t)[index+1:]...)
			return nil
		}
	}
	return errors.New("cannot find task that has given ID to delete")
}

func (t *TodoList) Edit(id int, title string) error {
	for index, task := range *t {
		if task.ID == id {
			(*t)[index].Title = title
			return nil
		}
	}
	return errors.New("cannot find task that has given ID to edit")
}
