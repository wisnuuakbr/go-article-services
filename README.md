# Submission Sagala

A tech test Sagala using :

<p align="center"><a href="https://go.dev" target="_blank"><img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" width="400"></a></p>

## Requirements

Sagala is currently extended with the following requirements.  
Instructions on how to use them in your own application are linked below.

| Requirement | Version |
| ----------- | ------- |
| Go          | 1.21.5  |
| Postgres    | 14.10.^ |
| Redis       | 8.11.5  |

## Installation

Make sure the requirements above already install on your system.  
Clone the project to your directory and install the dependencies.

```bash
$ git clone https://github.com/wisnuuakbr/sagala
$ cd sagala
$ go mod tidy
```

## Configuration

Copy the .env.example file and rename it to .env  
Change the config for your local server

```bash
APP_NAME = "sagala_test"
APP_ENV = "development"
APP_PORT = 3000

POSTGRES_HOST_MASTER = "localhost"
POSTGRES_PORT_MASTER = 5432
POSTGRES_USER_MASTER = "postgres"
POSTGRES_PASSWORD_MASTER = "postgres"
POSTGRES_DB_MASTER = "sagala_v1_db"

REDIS_HOST = "localhost"
REDIS_PORT = 6379
REDIS_PASSWORD = ""
REDIS_DB = 0
```

## Running Server

```bash
$ cd .\cmd\
$ go run .\main.go
```

## Output

![Output Create Response](docs/create_response.png)

![Output GetAll Response](docs/get_all_response.png)

![Output Search Response](docs/search_response.png)
