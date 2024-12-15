# RBAC System Proof of Concept (POC)

This project is a Proof of Concept (POC) for a Role-Based Access Control (RBAC) system implemented using Golang, GORM, and PostgreSQL. The system allows users to register, log in, and manage roles and permissions.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Database Schema](#database-schema)
- [Setup Instructions](#setup-instructions)
- [Running the Application](#running-the-application)
- [Testing with Postman](#testing-with-postman)
- [Endpoints](#endpoints)
- [License](#license)

## Features

- User registration and authentication
- Role management
- Permission management
- Role-based access control for protected routes

## Technologies Used

- **Golang**: Programming language for backend development.
- **GORM**: ORM library for Golang.
- **PostgreSQL**: Database management system.
- **Gin**: Web framework for Golang.

## Database Schema

The database consists of the following tables:

1. **Users**: Stores user information with roles.
6. **Products**: Stores product information listed by users.
7. **WebBuilds**: Stores web build information listed by users.


## Setup Instructions

1. **Install Go**:
   - Download and install Go from the [official Go website](https://golang.org/dl/).

2. **Clone the Project**:

3. **Set Up PostgreSQL**:
- Ensure PostgreSQL is installed and running.
- Create a new database:


4. **Configure Database Connection**:
- Update the database connection settings in `config/config.go` with your PostgreSQL credentials. you can add the connection configuration in .env file 

5. **Install Dependencies**:

```
 go mod tidy
 ```


6. **Start the Server**:

```
go run main.go
 ```
the server will be running on localhost:8080
