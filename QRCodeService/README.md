# Go API Project

This project is a Go-based API with PostgreSQL as the database. Below are the steps to generate the OpenAPI specification, spin up the Docker environment, and manage database migrations.

---

## Table of Contents

- [Generating OpenAPI Specification](#generating-openapi-specification)
- [Spinning Up Docker Compose](#spinning-up-docker-compose)
- [Creating and Running Migrations](#creating-and-running-migrations)

---

## Generating OpenAPI Specification

To generate the OpenAPI specification in JSON format:

1. Run the following command:
   ```bash
   swag init --parseDependency --parseInternal
   ```
2. This will create a file named `openapi.json` containing the OpenAPI specification for your API.

---

## Spinning Up Docker Compose

This project uses `docker-compose` to set up the PostgreSQL database and pgAdmin GUI.

1. Ensure you have Docker installed on your system.
2. Run the following command to spin up the containers:
   ```bash
   docker-compose up -d
   ```
3. Access the services:
   - **PostgreSQL**: Available at `localhost:5432`.
   - **pgAdmin**: Open your browser and visit [http://localhost:8080](http://localhost:8080).
     - Default credentials:
       - Email: `admin@admin.com`
       - Password: `admin`

4. To stop the containers:
   ```bash
   docker-compose down
   ```

---

## Creating and Running Migrations

This project uses **Goose** for database migrations.

### Creating a New Migration

1. Generate a new migration file:
   ```bash
   goose create <migration_name> sql
   ```
   Example:
   ```bash
   goose create add_services_table sql
   ```

2. This will create a single SQL file in the root directory. The file will look like:
   - \`20231115000123_add_services_table.sql\`

3. Edit the SQL file to include both the `up` and `down` migrations in a single file. Example:
   ```sql
   -- +migrate Up
   CREATE TABLE services (
       email VARCHAR NOT NULL,
       ipAddress VARCHAR NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
       PRIMARY KEY (email, ipAddress)
   );

   -- +migrate Down
   DROP TABLE IF EXISTS services;
   ```

---

### Running Migrations
Keep in mind this assumes the migrations are in the `migrations` directory.

1. Apply all pending migrations:
   ```bash
   goose -dir ./migrations postgres "$DATABASE_URL" up
   ```

2. Rollback the last migration:
   ```bash
   goose -dir ./migrations postgres "$DATABASE_URL" down
   ```

3. Check migration status:
   ```bash
   goose -dir ./migrations postgres "$DATABASE_URL" status
   ```

---

## Environment Variables

Ensure your `.env` file has the following configuration for database connection:
```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable
```

---

## Project Structure

```
my-go-app/
├── migrations/              # Migration files
├── openapi.json             # OpenAPI specification (generated)
├── main.go                  # Main application
├── migrations.go            # Migration handler
├── Dockerfile               # Docker configuration for the app
├── docker-compose.yml       # Docker Compose setup
├── .env                     # Environment variables
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksum
```

---
