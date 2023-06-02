# HTTP Server for Todo List CRUD Operations in Go

This repository includes a simple HTTP server implemented in Go. The server supports CRUD (Create, Read, Update, Delete) operations for a Todo List. 

## Features

- Create a new Todo
- Read all Todos
- Read a specific Todo by ID
- Delete a specific Todo by ID

## Endpoints

- GET `/api/todos`: Returns all the Todos.
- POST `/api/todos`: Creates a new Todo. 
- GET `/api/todos/{id}`: Returns a specific Todo.
- DELETE `/api/todos/{id}`: Deletes a specific Todo.

## Getting Started

1. Clone the repository to your local machine.
2. Navigate to the project root directory.
3. Run `go build` to compile the server.
4. Run `./main` (or `main.exe` on Windows) to start the server.

The server will start and listen on port `8080`. You can interact with it using any HTTP client like `curl` or Postman.

## Dependencies

This project uses the `gorilla/mux` package for request routing and URL parameter parsing. You can install it by running `go get -u github.com/gorilla/mux`.

## Note

This is a demonstration project. The Todo store is an in-memory map, and thus, the data will not persist when the server is restarted.

## License

This project is open source and available under the [MIT License](LICENSE).
