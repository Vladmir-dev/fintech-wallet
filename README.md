# Fintech Wallet

A Go-based fintech wallet API built with Gin, GORM, and PostgreSQL.

## Prerequisites

- [Go 1.23+](https://go.dev/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup & Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Vladmir-dev/fintech-wallet.git
cd fintech-wallet
```

### 2. Run with Docker Compose

The easiest way to run the application is using Docker Compose, which sets up both the application and PostgreSQL database:

```bash
docker-compose up --build
```

This will:
- Build the Go application
- Start a PostgreSQL database container
- Connect the app to the database
- Expose the API on port `8080`

### 3. Environment Variables

The application requires the following environment variables:

| Variable   | Description      | Default (Docker Compose) |
|------------|------------------|--------------------------|
| `DB_HOST`  | Database host    | `db`                     |
| `DB_PORT`  | Database port    | `5432`                   |
| `DB_USER`  | Database user    | `wallet_user`            |
| `DB_PASSWORD` | Database password | `wallet_pass`         |
| `DB_NAME`  | Database name    | `wallet_db`              |

When running with Docker Compose, these are automatically configured in the `docker-compose.yml` file.

### 4. Run Locally (Without Docker)

If you prefer to run locally without Docker:

1. Ensure PostgreSQL is installed and running
2. Create a `.env` file in the project root:
   ```bash
   cp .env.example .env
   ```
3. Update the `.env` file with your database credentials
4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

Once running, the API is available at `http://localhost:8080`.

## Development

### Build Docker Image

```bash
docker build -t fintech-wallet .
```

### Stop Services

```bash
docker-compose down
```

To also remove the database volume:

```bash
docker-compose down -v
```
