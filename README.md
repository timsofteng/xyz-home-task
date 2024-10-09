XYZ Home Task Application
====================

Description
-----------

XYZ Home Task is a multi-service application comprising:

-   **backend**: The server-side component that handles requests and interacts with the API's.
-   **frontend**: The client-side component that provides the user interface for interaction.
-   **doc-server**: A service that serves API documentation.

Requirements
------------

-   Docker
-   Docker Compose
-   Make

Getting Started
---------------

### 1\. Clone the Repository


`git clone git@github.com:timsofteng/xyz-home-task.git`

### 2. Start the Application

Run the following command to build and start all services:

`make up`

### 3\. Access the Services

-   **Backend Service**: Accessible at `http://localhost:8080/api/v1`
-   **Frontend Service**: Accessible at `http://localhost:3000`
-   **Documentation Server**: Accessible at `http://localhost:3030`

### 4\. Stopping the Services

To stop and remove all containers, networks, and volumes, run:

`make down`

Docker Compose Configuration
----------------------------

The `docker-compose.yml` file defines the configuration for the services. It includes:

-   **Backend**: Builds from the `./server` directory, exposes port `8080`, and uses environment variables to configure the service.
-   **Frontend**: Builds from the `./client` directory, exposes port `3000`, and depends on the `backend` service.
-   **Doc-Server**: Builds from the `./doc-server` directory, exposes port `3030`.

Troubleshooting
---------------

If you encounter issues, ensure that:

-   Docker and Docker Compose are installed and running.
-   The `.env` file is correctly configured.
-   Ports are not being used by other services on your local machine.

For detailed information on Docker and Docker Compose commands, refer to their official documentation.
