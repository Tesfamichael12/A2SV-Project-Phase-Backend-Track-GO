# Task Manager API Documentation

This document provides a comprehensive guide to the Task Manager API, including endpoint details, setup instructions, and development workflow enhancements.

---

## Table of Contents

1.  [API Endpoints](#api-endpoints)
    - [Get All Tasks](#get-all-tasks)
    - [Get Task by ID](#get-task-by-id)
    - [Add a New Task](#add-a-new-task)
    - [Update a Task](#update-a-task)
    - [Delete a Task](#delete-a-task)
2.  [Project Setup Guide](#project-setup-guide)
    - [Step 1: MongoDB Setup](#step-1-mongodb-setup)
    - [Step 2: Environment Configuration](#step-2-environment-configuration)
3.  [Development Workflow](#development-workflow)
    - [Setting Up Hot-Reloading with Air](#setting-up-hot-reloading-with-air)

---

## API Endpoints

This API allows you to manage a simple list of tasks.

### Get All Tasks

Retrieves a list of all tasks in the database.

-   **URL**: `/tasks`
-   **Method**: `GET`
-   **Success Response**:
    -   **Code**: `200 OK`
    -   **Content**: `[{ "id": "60d5ec49e77c2b5b3c8e4f2a", "title": "Finish project", "description": "Complete the final report.", "due_date": "2025-07-23T10:00:00Z", "status": "In_Progress" }]`
-   **Error Response**:
    -   **Code**: `500 Internal Server Error`
    -   **Content**: `{"error": "Failed to retrieve tasks"}`

### Get Task by ID

Retrieves a single task by its unique MongoDB ObjectID.

-   **URL**: `/tasks/:id`
-   **Method**: `GET`
-   **URL Params**:
    -   `id=[string]` (Required) - The MongoDB ObjectID of the task.
-   **Success Response**:
    -   **Code**: `200 OK`
    -   **Content**: `{ "id": "60d5ec49e77c2b5b3c8e4f2a", "title": "Finish project", ... }`
-   **Error Response**:
    -   **Code**: `400 Bad Request` - `{"error": "Invalid ID format"}`
    -   **Code**: `404 Not Found` - `{"error": "Task not found"}`

### Add a New Task

Creates a new task.

-   **URL**: `/tasks`
-   **Method**: `POST`
-   **Body (JSON)**:
    ```json
    {
        "title": "New API Task",
        "description": "Testing the POST endpoint.",
        "due_date": "2025-08-01T15:00:00Z",
        "status": "Pending"
    }
    ```
-   **Success Response**:
    -   **Code**: `201 Created`
    -   **Content**: `{"id": "60d5ecb0e77c2b5b3c8e4f2b"}`
-   **Error Response**:
    -   **Code**: `400 Bad Request` - `{"error": "Validation error details..."}`

### Update a Task

Updates an existing task by its ID.

-   **URL**: `/tasks/:id`
-   **Method**: `PUT`
-   **URL Params**:
    -   `id=[string]` (Required) - The MongoDB ObjectID of the task.
-   **Body (JSON)**:
    ```json
    {
        "title": "Updated Task Title",
        "status": "Completed"
    }
    ```
-   **Success Response**:
    -   **Code**: `200 OK`
    -   **Content**: `{"message": "Task updated successfully"}`
-   **Error Response**:
    -   **Code**: `404 Not Found` - `{"error": "Task not found or not modified"}`

### Delete a Task

Deletes a task by its ID.

-   **URL**: `/tasks/:id`
-   **Method**: `DELETE`
-   **URL Params**:
    -   `id=[string]` (Required) - The MongoDB ObjectID of the task.
-   **Success Response**:
    -   **Code**: `200 OK`
    -   **Content**: `{"message": "Task deleted successfully"}`
-   **Error Response**:
    -   **Code**: `404 Not Found` - `{"error": "Task not found"}`

---

## Project Setup Guide

Follow these steps to get the project running locally with a MongoDB database.

### Step 1: MongoDB Setup

This project uses MongoDB for persistent data storage. You can use a free cloud instance from [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) or a local installation.

1.  **Get a Connection String**: Once your database is set up, you'll get a connection string (URI). It will look something like this:
    `mongodb+srv://<username>:<password>@<cluster-address>/<database-name>?retryWrites=true&w=majority`

2.  **Configure IP Access**: If you're using MongoDB Atlas, make sure to add your current IP address to the IP Access List. You can find this under **Security > Network Access** in your Atlas project dashboard. This is a common reason for connection timeouts!

### Step 2: Environment Configuration

We use a `.env` file to securely manage the database connection string without hardcoding it.

1.  **Create the File**: In the root of the `task-manager` directory, create a file named `.env`.

2.  **Add the URL**: Add your MongoDB connection string to the `.env` file like this:
    ```
    DATABASE_URL=mongodb+srv://your_username:your_password@your_cluster.mongodb.net/taskdb
    ```
    The `data/task_service.go` file is configured to read this variable when the application starts, establishing the connection to your database.

---

## Development Workflow

To make development faster and more efficient, we use a tool for hot-reloading.

### Setting Up Hot-Reloading with Air

**Air** is a tool that automatically recompiles and restarts your Go application whenever you save a file.

1.  **Install Air**: Open your terminal and run the following command. Note that the project has moved, so we use the `air-verse` path:
    ```sh
    go install github.com/air-verse/air@latest
    ```

2.  **Create Configuration File**: In the `task-manager` directory, create a file named `.air.toml`. This file tells Air how to run your project. Since you are on Windows, it's important to specify the `.exe` extension.
    ```toml
    # .air.toml
    root = "."
    tmp_dir = "tmp"

    [build]
    cmd = "go build -o ./tmp/main.exe ."
    bin = "tmp/main.exe"
    include_ext = ["go", "tpl", "tmpl", "html"]
    exclude_dir = ["assets", "tmp", "vendor"]
    log = "air-build.log"

    [log]
    time = true

    [misc]
    clean_on_exit = true
    ```

3.  **Run with Air**: Instead of `go run .`, navigate to the `task-manager` directory and simply run:
    ```sh
    air
    ```
    Air will now watch your files. When you save a change, it will rebuild and restart the server instantly!
