# Golang REST APIs using Gin

This project demonstrates a REST API implementation using the [Gin](https://github.com/gin-gonic/gin) web framework in Golang. The API interacts with a PostgreSQL database, using [GORM](https://gorm.io/) as the ORM (Object Relational Mapper).

## Features

- RESTful API design
- PostgreSQL integration with GORM
- Basic CRUD operations
- Organized structure with handlers for different routes

## Tech Stack

- **Language:** Golang
- **Framework:** Gin
- **Database:** PostgreSQL
- **ORM:** GORM

## Installation

### Prerequisites

- Go 1.23 installed on your machine. [Download Go](https://golang.org/dl/)
- PostgreSQL installed and running.

### Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/logeshwarann-dev/golang-rest-apis.git
    cd golang-rest-apis
    ```

2. Install Go dependencies:

    ```bash
    go mod tidy
    ```

3. Set up PostgreSQL database:

    ```bash
    createdb <your_db_name>
    ```

4. Configure your environment variables (PostgreSQL settings):

    Create a `.env` file in the root directory and add the following:

    ```bash
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=your_db_name
    DB_HOST=localhost
    DB_PORT=5432
    ```

5. Run the application:

    ```bash
    go run main.go
    ```

The API will be running at `http://localhost:8080`.

## Project Structure

```bash
.
├── db
│   └── db.go               # Database connection and migrations  
│   └── models.go           # GORM models for entities
├── handlers
│   └── movies.go           # API handlers for various resources       
├── go.mod                  # Module dependencies
├── go.sum                  # Version hashes
└── main.go                 # Application entry point
```

## API Endpoints

### User Resource

1. **Create User**
    - `POST /movies`
    - Request Body:
      ```json
      {
        "title": "Deadpool and Wolverine",
        "description": "Deadpool recruits a variant of Wolverine to save his universe from extinction.",
        "year": 2024
      }
      ```
    - Response: `201 Created`

2. **Get All Users**
    - `GET /users`
    - Response:
      ```json
      [
        {
          "id": 1,
          "title": "Deadpool and Wolverine",
          "description": "Deadpool recruits a variant of Wolverine to save his universe from extinction.",
          "year": 2024
        }
      ]
      ```

3. **Get User by ID**
    - `GET /movies/1`
    - Response:
      ```json
        {
          "id": 1,
          "title": "Deadpool and Wolverine",
          "description": "Deadpool recruits a variant of Wolverine to save his universe from extinction.",
          "year": 2024
        }
      ```
    - Response: `200 OK`

4. **Update User**
    - `PUT /movies/1`
    - Request Body:
      ```json
      {
        "title": "Deadpool Universe"
      }
      ```
    - Response: `200 OK`

5. **Delete User**
    - `DELETE /movies/1`
    - Response: `204 No Content`

## Database Migrations

Database migrations are automatically handled using GORM. To modify the database structure, adjust your models and restart the application. GORM will manage schema changes accordingly.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

