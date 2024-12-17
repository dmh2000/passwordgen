# Password Generator Web Service

This is a simple web service that generates secure passwords through a HTTP interface. Written in Go, it provides both a web interface and a JSON API for password generation.

## Features

- Generates passwords with configurable length (minimum 24 characters)
- Optional symbol inclusion
- Ability to generate multiple passwords at once
- Returns both raw and formatted password versions
- Web interface for easy access
- RESTful API endpoint for programmatic access

## API Endpoints

### GET /
Serves the main web interface

### GET /generate
Generates passwords based on query parameters:
- `length`: Password length (≥24, defaults to 24)
- `count`: Number of passwords to generate (≥1, defaults to 1)
- `symbols`: Include symbols (true/false)

Returns JSON in the format:
```json
{
    "passwords": [
        {
            "raw": "unformatted-password",
            "formatted": "formatted-password"
        }
    ]
}
```

## Running the Service

The service runs on `http://localhost:8080` by default.

## Technical Details

- Uses Go's `embed` package for static file serving
- Implements template rendering for the web interface
- Includes error handling for invalid inputs
- Serves static assets from embedded filesystem

## Dependencies

- Built-in Go packages
- Custom password generation package (`sqirvy.xyz/passwords/internal/pwd`)

This service is designed to be a secure and reliable password generation solution, suitable for both human users through the web interface and automated systems through the API.
