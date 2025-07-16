[![Postman](https://skillicons.dev/icons?i=postman)](https://documenter.getpostman.com/view/43798053/2sB34ijzCo) **View the interactive Postman API documentation here:** [Task Manager API on Postman](https://documenter.getpostman.com/view/43798053/2sB34ijzCo)

# Task Manager API Documentation

This document describes the available endpoints for the Task Manager API. Use these endpoints to manage tasks (create, read, update, delete).

> **Note:** The `status` field is of the following values: `In_Progress`, `Completed`, or `Pending`.

## Endpoints

### 1. Get All Tasks

- **URL:** `/tasks`
- **Method:** GET
- **Description:** Returns a list of all tasks.
- **Response Example:**

```json
[
  {
    "id": 1,
    "title": "Task 1",
    "description": "First task",
    "due_date": "2025-07-16T12:00:00Z",
    "status": "In_Progress"
  },
  ...
]
```

### 2. Get Task by ID

- **URL:** `/tasks/{id}`
- **Method:** GET
- **Description:** Returns a single task by its ID.
- **Response Example:**

```json
{
  "id": 1,
  "title": "Task 1",
  "description": "First task",
  "due_date": "2025-07-16T12:00:00Z",
  "status": "In_Progress"
}
```

### 3. Add a New Task

- **URL:** `/tasks`
- **Method:** POST
- **Description:** Adds a new task. Provide task details in the request body.
- **Request Body Example:**

```json
{
  "title": "New Task",
  "description": "A new task",
  "due_date": "2025-07-20T12:00:00Z",
  "status": "Pending"
}
```

- **Response:** The created task object.

### 4. Update a Task

- **URL:** `/tasks/{id}`
- **Method:** PUT
- **Description:** Updates an existing task by ID. Provide updated fields in the request body.
- **Request Body Example:**

```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-07-21T12:00:00Z",
  "status": "Completed"
}
```

- **Response:** The updated task object.

### 5. Delete a Task

- **URL:** `/tasks/{id}`
- **Method:** DELETE
- **Description:** Deletes a task by its ID.
- **Response Example:**

```json
{
  "message": "Task deleted successfully"
}
```

---
