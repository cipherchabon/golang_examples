# Simple Go HTTP Server

This repository contains a simple HTTP server implemented in Go. It serves as a basic example of how to handle HTTP requests and serve static files using Go's built-in `net/http` package.

## Features

1. **Default greeting**: Accessing the root URL (`http://localhost:8080/`) returns a default greeting, "Hello, World!".

2. **Static file server**: Files located in the `static` directory can be accessed under the `/web/` endpoint. For instance, accessing `http://localhost:8080/web/index.html` will return the contents of `static/index.html`.

3. **Personalized greeting**: The server can return a personalized greeting by visiting the `/greeter/` endpoint followed by your name. For instance, accessing `http://localhost:8080/greeter/YourName` will return "Hello, YourName!".

## Usage

To run the server, make sure you have [Go installed](https://golang.org/doc/install) on your machine, then execute:

```bash
go run main.go
```

## Future Enhancements
This server serves as a simple demonstration of Go's capabilities. Future enhancements could include more robust error handling, support for additional HTTP methods, and the integration of a database for more dynamic content.

## License
This project is released under the MIT License. See LICENSE file for details.
