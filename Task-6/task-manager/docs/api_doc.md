# Task Manager API Documentation

This document provides a comprehensive guide to the Task Manager API, including authentication, authorization, endpoint details, setup instructions, and development workflow enhancements.

---

## Table of Contents

1.  [API Endpoints](#api-endpoints)
    - [User Management](#user-management)
    - [Task Management (Protected)](#task-management-protected)
2.  [Authentication & Authorization](#authentication--authorization)
3.  [Project Setup Guide](#project-setup-guide)
4.  [Development Workflow](#development-workflow)
    - [Setting Up Hot-Reloading with Air](#setting-up-hot-reloading-with-air)

---

## API Endpoints

### User Management

#### Register a New User

- **URL**: `/register`
- **Method**: `POST`
- **Description**: Create a new user account.
- **Request Body Example**:
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- **Response Example**:
  ```json
  {
    "message": "User registered successfully"
  }
  ```

#### Login

- **URL**: `/login`
- **Method**: `POST`
- **Description**: Authenticate user and receive a JWT token.
- **Request Body Example**:
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- **Response Example**:
  ```json
  {
    "token": "<JWT_TOKEN>"
  }
  ```

### Task Management (Protected)

> **Note:** All `/tasks` endpoints require a valid JWT token in the `Authorization` header: `Bearer <JWT_TOKEN>`

#### Get All Tasks

- **URL**: `/tasks`
- **Method**: `GET`
- **Header**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**: List of tasks.

#### Get Task by ID

- **URL**: `/tasks/{id}`
- **Method**: `GET`
- **Header**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**: Task object.

#### Add a New Task

- **URL**: `/tasks`
- **Method**: `POST`
- **Header**: `Authorization: Bearer <JWT_TOKEN>`
- **Request Body**:
  ```json
  {
    "title": "New Task",
    "description": "A new task",
    "due_date": "2025-07-20T12:00:00Z",
    "status": "Pending"
  }
  ```
- **Response**: Created task ID.

#### Update a Task

- **URL**: `/tasks/{id}`
- **Method**: `PUT`
- **Header**: `Authorization: Bearer <JWT_TOKEN>`
- **Request Body**:
  ```json
  {
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2025-07-21T12:00:00Z",
    "status": "Completed"
  }
  ```
- **Response**: Success message.

#### Delete a Task

- **URL**: `/tasks/{id}`
- **Method**: `DELETE`
- **Header**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**: Success message.

---

## Authentication & Authorization

### How It Works

- Register a user with `/register`.
- Login with `/login` to receive a JWT token.
- Include the token in the `Authorization` header for all `/tasks` requests.
- Only authenticated users can access task endpoints.

### JWT Token Example

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### User Roles

- (Optional) You can extend the user model and middleware to support roles (e.g., admin, user) and restrict endpoints as needed.

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

3.  **Add JWT Secret Key**: Also, add a secret key for JWT signing:
    ```
    JWT_SECRET_KEY=your_super_secret_key
    ```

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

---

## Testing

- Use Postman or similar tools to test registration, login, and all protected endpoints.
- Ensure you include the JWT token in the `Authorization` header for all `/tasks` requests.
