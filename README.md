# go-todo-gin-gorm
This is a personal learning project using Go, Gin, and GORM to practice web API development and database operations.

## Technologies Used

### Backend
- **Programming Language**: Go (Golang)
- **Web Framework**: Gin
- **ORM (Object-Relational Mapping)**: GORM
- **Database**: MySQL (or any compatible SQL database)

### Frontend
- **WIP**: The frontend will be built using React in the future.

### Common Tools
- **Containerization**: Docker
- **Service Orchestration**: Docker Compose

## How to Run
1. Run Docker Compose to start the backend and frontend services:
   ```bash
   docker compose up
   ```
2. Access the services:
   - **Backend**: `http://localhost:8080`
   - **Frontend**: **WIP**
   - **Database (MySQL)**: Accessible on port `33306` (e.g., `mysql://localhost:33306`)

## Future Plans
- Build a simple frontend using React to interact with the backend API.
    - Display a list of todos.
    - Add new todos through a form.
    - Mark todos as completed or delete them.
