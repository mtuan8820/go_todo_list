# Go-Todo 
This is a simple To-do list application implemented by Go, HTML, CSS, JS. The app helps user in tracking and managing tasks with features include: Add, Delete, Toggle & GetAll. Data is stored 'in-memory' using slice.
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
Response example:
```json
[
    {
        "id": 1,
        "title": "task title 3",
        "description": "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.",
        "completed": false
    }
]
```  
2. Add a new task    
API: `POST /tasks`  
Response example:
```json
{
    "message": "Task created successfully"
}
```
3. Delete a task  
API: `DELETE /tasks/{id}`  
Response example:
```json
{
    "message": "Task deleted successfully"
}
```

4. Toggle complete state of a task  
API: `PUT /tasks/{id}`  
Response example:
```json
{
    "message": "Task toggled successfully"
}
```
5. View todo app UI  
API: `GET /view/tasks`
![ToDo app UI](https://github.com/user-attachments/assets/5086a830-04c5-48fd-8a0c-76a032c65a84)


## Project structure
System design: Dependency Inversion principle
```bash
cmd
   |-- main.go
controller
   |-- controller.go
go.mod
go.sum
pkg
   |-- models
   |   |-- task.go
service
   |-- service.go
web
   |-- static
   |   |-- css
   |   |   |-- index.css
   |   |-- js
   |   |   |-- index.js
   |-- templates
   |   |-- index.html
```
