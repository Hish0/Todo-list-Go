## To-Do List Application

This repository contains the code for a simple To-Do List application built using React and Go.

### Backend

- **`backend/routes/`**: Contains Go handlers for API endpoints.
  - **`auth.go`**: Handlers for user registration and login, **including JWT token generation and verification**.
  - **`tasks.go`**: Handlers for creating, retrieving, updating, and deleting tasks, **with JWT token authentication required for protected routes**.
- **`backend/database/`**: Handles database connections and migrations.
  - **`db.go`**: Connects to the PostgreSQL database.
- **`backend/models/`**: Defines the data structures used in the application.
  - **`models.go`**: Defines the `User` and `Task` models.
- **`main.go`**: The main entry point of the Go server.

### Frontend

- **`frontend/src/`**: Contains the React code for the application.
  - **`App.js`**: The main component that renders the entire application.
  - **`LoginForm.js`**: Component for user login, **handles storing and retrieving JWT tokens from localStorage**.
  - **`RegisterForm.js`**: Component for user registration.
  - **`TaskItem.js`**: Component for displaying individual tasks.
  - **`TaskList.js`**: Component for displaying a list of tasks.
  - **`index.js`**: The entry point of the React application.
  - **`index.css`**: Stylesheet for the application.
  - **`services/`**: Contains utility functions for interacting with the backend API.
    - **`AuthService.js`**: Handles user authentication and token management, **including storing, retrieving, and sending JWT tokens with requests**.
- **`frontend/public/`**: Contains static files, including the HTML template for the application.
  - **`index.html`**: The base HTML file.

## Running the Application

1. **Setup Environment Variables:**
   - Create a `.env` file in the root directory and add the following environment variables:
     ```
     DATABASE_URL=postgres://username:password@host:port/database
     JWT_SECRET=your_secret_key
     ```
     - Replace `username`, `password`, `host`, `port`, and `database` with your PostgreSQL database credentials.
     - Replace `your_secret_key` with a strong and unique secret key for JWT signing.
2. **Start the Backend Server:**
   - Run the command `go run main.go` in the `backend` directory.
3. **Start the Frontend Server:**
   - Run the command `npm start` in the `frontend` directory.

The application should be accessible at `http://localhost:3000`.

## Features

- **User Authentication:** Users can register and log in to the application using JWT authentication.
- **Task Management:** Users can create, view, update, and delete tasks.
- **Persistence:** Tasks are stored in a PostgreSQL database.

## Technologies Used

- **Backend:**
  - **Go:** Programming language for the backend server.
  - **Gin:** Web framework for Go.
  - **GORM:** ORM library for Go.
  - **bcrypt:** Library for password hashing.
  - **PostgreSQL:** Database management system.
  - **JWT:** JSON Web Token for authentication.
- **Frontend:**
  - **React:** JavaScript library for building user interfaces.
  - **Axios:** HTTP client for making API requests.
  - **JWT:** JSON Web Token for authentication.
  - **localStorage:** Browser storage for storing user tokens.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request to improve the application.
