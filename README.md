# Shortify - URL Shortener

A simple and fast URL shortener built with Go, MySQL, and a modern web interface.

## Features

- Shorten long URLs to short codes
- Redirect short codes to original URLs
- Dark theme UI with black and white design
- MySQL database for persistence
- Docker containerization for easy deployment

## Quick Start with Docker

1. **Clone and navigate to the project:**

   ```bash
   git clone <your-repo>
   cd shortify
   ```

2. **Start the application:**

   ```bash
   docker-compose up -d
   ```

3. **Access the application:**
   - Web UI: http://localhost:8080
   - API: http://localhost:8080/api/shorten

## Manual Setup

1. **Install dependencies:**

   ```bash
   go mod download
   ```

2. **Set up MySQL database:**

   - Create database: `shortify`
   - Run the SQL from `init.sql` to create tables

3. **Configure environment variables:**

   ```bash
   cp env.example .env
   # Edit .env with your database credentials
   ```

4. **Run the application:**
   ```bash
   go run cmd/shortner/main.go
   ```

## API Usage

### Shorten URL

```bash
curl -X POST http://localhost:8080/api/shorten \
  -d "url=https://example.com/very/long/url"
```

### Redirect

Visit: `http://localhost:8080/s/{short_code}`

## Deployment Options

### Docker (Recommended)

```bash
docker-compose up -d
```

### Cloud Deployment

- **Heroku**: Use the included Dockerfile
- **AWS/GCP/Azure**: Deploy using Docker containers
- **VPS**: Use docker-compose for easy setup

## Environment Variables

| Variable    | Default     | Description         |
| ----------- | ----------- | ------------------- |
| DB_HOST     | localhost   | MySQL host          |
| DB_PORT     | 3306        | MySQL port          |
| DB_USER     | shortify    | MySQL username      |
| DB_PASSWORD | shortify123 | MySQL password      |
| DB_NAME     | shortify    | MySQL database name |
| PORT        | 8080        | Server port         |

## Production Considerations

1. **Security:**

   - Change default database passwords
   - Use environment variables for sensitive data
   - Enable HTTPS in production

2. **Performance:**

   - Configure MySQL connection pooling
   - Add caching layer (Redis) for high traffic
   - Use CDN for static assets

3. **Monitoring:**
   - Add health check endpoints
   - Set up logging and monitoring
   - Configure backup strategies

## Development

```bash
# Run tests
go test ./...

# Build binary
go build -o shortify ./cmd/shortner

# Run with hot reload (install air first)
air
```
