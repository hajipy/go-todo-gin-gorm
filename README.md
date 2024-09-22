# go-todo-gin-gorm
This is a personal learning project focused on mastering Go, Gin, and GORM. The frontend uses React, but any frontend framework could be used.

## Technologies Used

### Backend
- **Programming Language**: Go (Golang)
- **Web Framework**: Gin
- **ORM (Object-Relational Mapping)**: GORM
- **Database**: MySQL (or any compatible SQL database)

### Frontend
- **UI Library**: React
- **State Management**: React Hooks
- **API Integration**: Fetch API for interaction with backend
- **Build Tool**: Vite
- **Code Formatting & Linting**: Biome

### Common Tools
- **Containerization**: Docker
- **Service Orchestration**: Docker Compose

## Prerequisites
- **Docker**: Ensure Docker is installed and running (latest version recommended)
- **Docker Compose**: Make sure Docker Compose is available (part of Docker Desktop or standalone)

## How to Run
1. Run Docker Compose to start the backend and frontend services:
   ```bash
   docker compose up
   ```
2. Access the services:
   - **Backend**: `http://localhost:8080`
   - **Frontend**: `http://localhost:5173`
   - **Database (MySQL)**: Accessible on port `33306` (e.g., `mysql://localhost:33306`)

## Future Plans
- Improve UI/UX
    - Add basic styling for better design
- Refactor Frontend
    - Split code into well-structured components for clarity
