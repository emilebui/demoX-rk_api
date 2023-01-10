# demoX-rk_api

- Just a demo microservice API project using rk_boot
- To learn more about the rk_echo, visit here: https://github.com/rookie-ninja/rk-boot

## Installation Instructions

### Requirements
- git
- golang 1.19

### Initial setup
- Install swaggo/swag:
    - Run: `go install github.com/swaggo/swag/cmd/swag@latest`
- Install dependencies:
    - Run: `go mod download`

### Run application
- Dev mode: `go run cmd/dev/main.go`
- Prod mode: `go run cmd/prod/main.go`

- **Note**:
    - You may need to run the database and config them first to run the app
    - After running the app, you can check out the api docs URLs from the stdout

### Docker

- Build: `docker build -t rk_echo:latest .`
- To build and run the entire stack (rk_echo, mongo, postgres, redis):
    - Run: `docker-compose up -d`

### Testing

- Run: `go test ./... -cover`

### Problems when running/installation
- If you encounter some problems with sqlite during running or installation process
    - Remove previous installed gcc package (if you use minGW)
    - Install this library then restart your machine: https://jmeubank.github.io/tdm-gcc/
    - If the issue isn't resolved, run the app in docker
        - Run `docker-compose up -d`

### Template Structure

- `cmd`: This contains all the run commands for the app
    - `dev/main.go`: run the app in dev mode (auto init swagger before running)
    - `prod/main.go`: run the app in prod mode (no swagger init)
- `docs`: Contains the OpenAPI documentation
- `pkg`: Contains all the pkg, helper modules to build Echo REST API
    - `swaginit`: Contains helper function for generating swagger file automatically
    - `conn`: Contains helper functions to connect to message queue (kafka, redis)
    - `unitest`: Contains helper functions for unit testing
- `internal`: Contain endpoint and logic for each API. The structure of this module is as follows:
    - `handler`: Handle the all the endpoints of the application (take request, get input and pass it to logic layer)
    - `repository`: Persistence handlers, all the persistence accees (database, message queue) are handled here
    - `service`: Logic Layer, all the logic of the service are handled here
- `boot.yaml`: The boot config contains all the configuration to run rk_boot application. (Configuration as code)