# Task Management REST API Documentation

## Base URL
http://localhost:8080

yaml
Copy code

---

## Endpoints

### 1. Get all tasks
**GET** `/tasks`

**Response Example**
```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, Bread, Eggs",
    "due_date": "2025-11-10",
    "status": "pending"
  }
]
2. Get task by ID
GET /tasks/:id

Response Example

json
Copy code
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Milk, Bread, Eggs",
  "due_date": "2025-11-10",
  "status": "pending"
}
Error (404)

json
Copy code
{"error": "Task not found"}
3. Create a new task
POST /tasks

Request Example

json
Copy code
{
  "title": "Finish report",
  "description": "Submit monthly report to manager",
  "due_date": "2025-11-15",
  "status": "in-progress"
}
Response (201 Created)

json
Copy code
{
  "id": 2,
  "title": "Finish report",
  "description": "Submit monthly report to manager",
  "due_date": "2025-11-15",
  "status": "in-progress"
}
4. Update a task
PUT /tasks/:id

Request Example

json
Copy code
{
  "status": "completed"
}
Response Example

json
Copy code
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Milk, Bread, Eggs",
  "due_date": "2025-11-10",
  "status": "completed"
}
5. Delete a task
DELETE /tasks/:id

Response Example

json
Copy code
{"message": "Task deleted successfully"}
Testing
Use Postman or curl:

bash
Copy code
curl -X GET http://localhost:8080/tasks
Response Codes
Code	Meaning
200	OK
201	Created
400	Bad Request
404	Not Found

yaml
Copy code

---

## ✅ Folder structure summary

task_manager/
├── main.go
├── controllers/
│ └── task_controller.go
├── models/
│ └── task.go
├── data/
│ └── task_service.go
├── router/
│ └── router.go
├── docs/
│ └── api_documentation.md
└── go.mod

yaml
Copy code

---

### 🧪 To run and test

```bash
# from task_manager/
go mod init task_manager
go get github.com/gin-gonic/gin
go run main.go
Now open Postman or your browser to http://localhost:8080/tasks