# URL Shortener

A simple and efficient URL shortening service built with Go, PostgreSQL, and HTMX.

## Features

- Shorten long URLs to compact, shareable links
- Redirect short URLs to original destinations
- In-memory caching for improved performance
- Simple web interface with HTMX for dynamic updates
- PostgreSQL database for persistent storage
- Docker support for easy deployment

## Tech Stack

- **Backend**: Go
- **Database**: PostgreSQL
- **Frontend**: HTML + HTMX
- **Caching**: In-memory cache
- **Configuration**: Environment variables
- **Deployment**: Docker

## Prerequisites

- Go 1.21+
- PostgreSQL
- Docker (recommended)

## Installation

### Using Docker (Recommended)

1. Clone the repository:
   ```bash
   git clone https://github.com/furkankarayel/URL_Shortener.git
   cd URL_Shortener
   ```

2. Build and run the container:
   ```bash
   docker build -t url-shortener .
   docker run -p 8080:8080 -p 5432:5432 url-shortener
   ```

### Manual Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/furkankarayel/URL_Shortener.git
   cd URL_Shortener
   ```

2. Set up PostgreSQL:
   ```bash
   make postgres
   make createdb
   make migrateup
   ```

3. Run the application:
   ```bash
   make run
   ```

## Configuration

Create a `.env` file with your database credentials. Take a look at `.env.example` file.
## API Endpoints

- `POST /url/shorten` - Shorten a URL
- `GET /url/{shortCode}` - Redirect to original URL

## Makefile Commands

| Command       | Description                          |
|---------------|--------------------------------------|
| `make postgres` | Start PostgreSQL container          |
| `make createdb` | Create database                     |
| `make migrateup` | Run database migrations             |
| `make run`     | Start the application     

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/YourFeatureName`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeatureName`)
5. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please open an issue.