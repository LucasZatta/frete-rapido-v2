# Frete Grátis Challenge

## Project Structure

```
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
├── init.sql
├── internal
│   ├── database
│   │   └── database.go
│   ├── quote
│   │   ├── handlers
│   │   │   └── handlers.go
│   │   ├── models
│   │   │   ├── quoteModel.go
│   │   │   └── requestModel.go
│   │   ├── repository
│   │   │   ├── mocks
│   │   │   │   └── QuoteRepository.go
│   │   │   └── repository.go
│   │   └── service
│   │       ├── service.go
│   │       └── service_test.go
│   ├── server
│   │   └── server.go
│   └── util
│       └── util.go
├── Dockerfile
├── Makefile
└── docker-compose.yml
```

### Structure

This project was implemented using the [Standard Go Project Layout](https://github.com/golang-standards/project-layout), dependency injection for improved testability and maintainability and a repository pattern for database operations(Postgres).  

## Rest API

Built using http.Server as a wrapper with Gin engine as a http.Handler in order to make the server easily customizable.

### Post    /quote

Receives data necessary to build the request body for a 3rd party API and builds a fictitious shipment quote with the API response. Data is then saved in the database for later use.

### GET   /metrics?last_quotes={?}

Receives an optional parameter that limits the database search. It retrieves entries from the database and returns a new object containing several metrics.

## Running the project

In order to run the project, after cloning the repo, the user needs to fill a .env file as follows:  

```{env}
PORT=8080
APP_ENV=local

DB_HOST=postgres_docker
DB_PORT=5432
DB_DATABASE=fretegratis
DB_USERNAME=admin
DB_PASSWORD=password1234
DB_SCHEMA=public

API_PATH=
API_TOKEN=
PLATFORM_CODE=
CNPJ=
DISPATCHER_ZIPCODE= 

```

Then just use the make docker-run command in the terminal and test away 👍.
If permission is denied, simply run  

`sudo docker-compose up` or `sudo docker compose up`

Also, if you'd like to keep using the same terminal window for curl requests, use:  
`sudo docker-compose up -d`