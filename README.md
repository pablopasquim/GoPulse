# GoPulse

GoPulse is a RESTful API built with Go using the Gin framework. This project was developed as part of the Alura API Rest course and demonstrates best practices for building APIs, organizing code, and integrating with a PostgreSQL database.

## Features

- RESTful endpoints for managing resources
- PostgreSQL integration with Docker Compose
- Modular code organization (controllers, models, routes, data)
- Example environment setup for local development
- Ready-to-use database management with pgAdmin

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- Go 1.25 or newer (for local development outside Docker)

### Environment Variables

Create a `.env` file in the project root with the following variables:

```
POSTGRES_USER=yourusername
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=gopulse
PGADMIN_DEFAULT_EMAIL=admin@example.com
PGADMIN_DEFAULT_PASSWORD=yourpassword
```

### Running with Docker Compose

To start the API server and PostgreSQL database using Docker Compose:

```bash
docker-compose up
```

- The API will run on [http://localhost:8080](http://localhost:8080)
- PostgreSQL will be available on port `5432`
- pgAdmin (database management) will be accessible at [http://localhost:54321](http://localhost:54321)

### Local Development

To run the API locally (requires Go installed):

```bash
go run main.go
```

## API Endpoints

> You can document your main endpoints here. For example:

- `GET /api/items/` - List item
- `GET /api/items/:id` - List item for ID
- `POST /api/items/` - Create a new item 
- `PUT /api/items/:id` - Update a item for ID
- `DELETE /api/items/:id` - Delete a item for ID

## Project Structure

- `main.go` — Entry point of the application
- `controllers/` — Request handlers and business logic
- `routes/` — Endpoint definitions
- `data/` — Database connection and setup
- `models/` — Data models and schema definitions

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE)

---

Developed as part of the Alura API Rest course.
