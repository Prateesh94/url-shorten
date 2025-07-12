# url-shorten

A simple URL shortening service written in Go.

## Overview

`url-shorten` is a RESTful API service that lets users create shortened URLs, retrieve the original URLs, update or delete them, and track basic access statistics. It uses PostgreSQL for persistent storage and the Gorilla Mux router for HTTP routing.

## Features

- **Shorten URLs:** Generate a unique short code for a given long URL.
- **Retrieve original URL:** Get the original URL from a short code.
- **Update URLs:** Update the long URL associated with a short code.
- **Delete URLs:** Remove a short code and its URL mapping.
- **Stats:** View statistics such as creation/update time and access count for a short URL.

## API Endpoints

| Method | Endpoint                  | Description                       |
|--------|---------------------------|-----------------------------------|
| POST   | `/shorten`                | Shorten a long URL                |
| GET    | `/shorten/{url}`          | Retrieve the original URL         |
| PUT    | `/shorten/{url}`          | Update the long URL               |
| DELETE | `/shorten/{url}`          | Delete a shortened URL            |
| GET    | `/shorten/{url}/stats`    | Get stats for a shortened URL     |

## Example Usage

### Shorten a URL

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten
```

### Retrieve original URL

```bash
curl -X GET http://localhost:8080/shorten/{shortCode}
```

### Update long URL

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"url":"https://new-url.com"}' http://localhost:8080/shorten/{shortCode}
```

### Delete shortened URL

```bash
curl -X DELETE http://localhost:8080/shorten/{shortCode}
```

### Get stats

```bash
curl -X GET http://localhost:8080/shorten/{shortCode}/stats
```

## Contributing

Feel free to fork the repository and submit pull requests for improvements or bug fixes.

## License

This project does not yet specify a license.
