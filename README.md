# Todo List API

This is a simple todo list API built with Go and Echo.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository.
2. Copy .env.example to your own .env and replace environment variables' values with your own, if needed. 
3. Run `make up` in project's directory (docker compose required). 
4. Go to localhost:8080 to access API.

### Known bugs
- Database in docker seems to be loading after the go container, and depends_on does not prevent it. Use `docker compose up -d database` before `make up` to start it first.