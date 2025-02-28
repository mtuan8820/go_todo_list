# ToDo - App for tracking and managing tasks effectively.
This is a simple To-do list application implemented by Go, HTML, CSS, JS. The app helps user in tracking and managing tasks with features include: Add, Delete, Toggle & GetAll. Data is store 'in-memory' using slice.
## Table of Contents
## Features
* Add: add new task with required title and optional description.
* Toggle: change the complete state of a task with given ID
* Delete: delete a task with given ID
* GetAll: retrieve all tasks.
## Getting Started
### Prequisites
Go: Ensure you have Go installed
### Installation
1. Clone the repository
```
git clone https://github.com/mtuan8820/go_todo_list.git
cd go_todo_list
```
2. Build the application
```
go build
```
3. Run the application
```
cd cmd
go run main.go
```
The server will start on `http://localhost:8080`
### API endpoints
1. Get All Tasks
API: `GET /tasks`

2. Add a new task
API: `POST /tasks `
  
3. Delete a task
API: `DELETE /tasks/{id}`

4. Toggle complete state of a task
API: `PUT /tasks/{id}`

## Project structure
