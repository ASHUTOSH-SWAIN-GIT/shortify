<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/c91d5843-d9b7-4a32-aacc-e1caced18d15" />

# Shortify - URL Shortener

A modern, fast, and lightweight URL shortener built with Go and a sleek dark-themed web interface. Transform long URLs into short, memorable links with ease.

## Features

- **URL Shortening**: Convert long URLs into short, shareable codes
- **Smart Redirects**: Seamlessly redirect short codes to original URLs
- **Modern UI**: Beautiful dark theme with black and white design
- **Database Persistence**: Store URLs with MySQL/PostgreSQL support
- **RESTful API**: Clean API endpoints for integration
- **Responsive Design**: Works perfectly on desktop and mobile
- **Fast Performance**: Built with Go for high-speed operations

## Project Structure

```
shortify/
├── cmd/shortner/          # Main application entry point
├── internals/             # Core application logic
│   ├── handler.go         # HTTP request handlers
│   ├── shortner.go        # URL shortening logic
│   └── db/                # Database layer
│       ├── mysql.go       # MySQL database operations
│       └── models/        # Data models
├── web/                   # Frontend assets
│   ├── index.html         # Main web interface
│   └── js/                # JavaScript files
│       └── main.js        # Frontend logic
└── README.md              # This file
```

## Quick Start

### Prerequisites

- Go 1.24+ installed
- MySQL or PostgreSQL database
- Git

### Local Development

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/shortify.git
   cd shortify
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Set up your database:**

   - Create a database named `shortify`
   - Run the SQL schema (see Database Setup section)

4. **Configure environment variables:**

   ```bash
   # Create .env file with your database credentials
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=shortify
   PORT=8080
   ```

5. **Run the application:**

   ```bash
   go run cmd/shortner/main.go
   ```

6. **Access the application:**
   - Web UI: http://localhost:8080
   - API: http://localhost:8080/api/shorten

## Database Setup

### MySQL

```sql
CREATE DATABASE shortify;
USE shortify;

CREATE TABLE urls (
    short_code VARCHAR(10) PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### PostgreSQL

```sql
CREATE DATABASE shortify;

\c shortify;

CREATE TABLE urls (
    short_code VARCHAR(10) PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## API Usage

### Shorten URL

```bash
curl -X POST http://localhost:8080/api/shorten \
  -d "url=https://example.com/very/long/url"
```

**Response:**

```json
{
  "short_code": "abc123",
  "short_url": "http://localhost:8080/s/abc123",
  "long_url": "https://example.com/very/long/url"
}
```

### Redirect

Visit: `http://localhost:8080/s/{short_code}`

The application will automatically redirect you to the original URL.

### API Endpoints

| Method | Endpoint       | Description              |
| ------ | -------------- | ------------------------ |
| `POST` | `/api/shorten` | Shorten a long URL       |
| `GET`  | `/s/{code}`    | Redirect to original URL |
| `GET`  | `/`            | Web interface            |

## Environment Variables

| Variable      | Default     | Description       |
| ------------- | ----------- | ----------------- |
| `DB_HOST`     | localhost   | Database host     |
| `DB_PORT`     | 3306        | Database port     |
| `DB_USER`     | shortify    | Database username |
| `DB_PASSWORD` | shortify123 | Database password |
| `DB_NAME`     | shortify    | Database name     |
| `PORT`        | 8080        | Server port       |

## Development

### Prerequisites

- Go 1.24+
- MySQL or PostgreSQL
- Git

### Running Tests

```bash
go test ./...
```

### Building Binary

```bash
go build -o shortify ./cmd/shortner
```

### Hot Reload (Optional)

```bash
# Install air for hot reloading
go install github.com/cosmtrek/air@latest
air
```
