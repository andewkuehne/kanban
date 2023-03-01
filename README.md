# Kanban Board Web Application
This is a simple web application for managing tasks using a kanban board. The application is built using Golang for the backend and React for the frontend.

## Installation
To run the application locally, you need to have Golang and Node.js installed on your machine. You also need to have Docker and Docker Compose installed if you want to run the application in a Docker container.

1. Clone the repository to your local machine:
`git clone https://github.com/andrewkuehne/kanban.git`
2. Navigate to the project directory:
`cd kanban`
3. Install the dependencies for the backend:
`cd kanban-backend && go mod download && cd ..`
4. Install the dependencies for the frontend:
`cd kanban-frontend`
`npm install`
5. Start the application:
\# Start the backend
`cd ..\kanban-backend && go run main.go`

\# Start the frontend
`cd ..\kanban-frontend`
`npm start`
Open your web browser and navigate to http://localhost:3000 to access the application.

## Usage
The application has a simple user interface for managing tasks using a kanban board. You can create new tasks, update existing tasks, and delete tasks.

##Docker
To run the application in a Docker container, follow these steps:

Navigate to the project directory:
`cd kanban-board`
Build the Docker image:
`docker-compose build`
Start the Docker container:
`docker-compose up`
Open your web browser and navigate to http://localhost:3000 to access the application.

License
This project is licensed under the MIT License. See the LICENSE file for details.