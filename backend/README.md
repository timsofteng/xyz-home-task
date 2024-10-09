# XYZ Home Task Backend

This is the backend application for the XYZ home task, designed to run a server based on the OpenAPI specification. This README outlines the main commands and instructions to build, run, test, and manage the project using the provided `Makefile`.

## Prerequisites

Before you can build and run the project, make sure the following dependencies are installed:

- [Go](https://golang.org/) (version 1.16+)
- [Docker](https://www.docker.com/)
- [air](https://github.com/cosmtrek/air) (for live reloading in development)
- [golangci-lint](https://golangci-lint.run/) (for linting)

You also need to create a `.env` file based on the environment variables in the `Makefile` or export them directly in your shell.

## Environment Variables

The following environment variables are used by the application:

- `APP_NAME`: Name of the application (default: `xyz-home-task-backend`)
- `HOST`: Host address for the server (default: `0.0.0.0`)
- `HTTP_SERVER_PORT`: Port for the HTTP server (default: `8080`)
- `LOG_LEVEL`: Log level for the application (default: `info`)
- `SPEC_PATH`: Path to the OpenAPI specification file (default: `../spec/openapi.yaml`)

## Makefile Commands

### Build the Application

To build the Go application:

```bash
make build
```
This will create the binary in the bin/ directory with the name specified in the APP_NAME environment variable.

## Run the Application
You can run the application in development mode using:

```bash
make dev
```
This will set the required environment variables and run the application with live reloading using air.

Alternatively, you can run the application normally using:

```bash
make run
```

or run production build:

```bash
make build && make app
```

## Testing
To run all the tests:

```bash
make test
```

## Linting
To run linting on the codebase:

```bash
make test
```
## Code Generation
If you need to regenerate the HTTP server handler based on the OpenAPI specification:

```bash
make codegen-http-server-handler
```
This will use oapi-codegen to generate the code based on the specification file at SPEC_PATH.

## Docker Commands
Build the Docker Container
To build the Docker image for the application:

```bash
make container-build
```

## Run the Docker Container
To run the application in a Docker container:

```bash
make container-run
```
This will start the container and expose the port defined by HTTP_SERVER_PORT to the host.
