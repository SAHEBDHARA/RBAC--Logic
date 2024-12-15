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

1. **Users**: Stores user information.
2. **Roles**: Defines roles that can be assigned to users.
3. **Permissions**: Lists permissions that can be granted to roles.
4. **UserRoles**: A join table to associate users with roles.
5. **RolePermissions**: A join table to associate roles with permissions.

### Example Schema

