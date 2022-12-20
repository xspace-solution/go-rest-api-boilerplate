# Go REST API Project for Dan Lam 

Project for Go API with a PostgreSQL database.

## Requirements
At a minimum, Golang, Postgres, and Git need to be installed on your local machine. Optionally, you may run the app using Docker.
* [Install Go](https://golang.org/doc/install)
* [Install Postgres](https://www.postgresql.org/download/)
* [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* [Install Docker](https://docs.docker.com/v17.09/engine/installation/)

### Updating Project Module Names
**Important! Replace ALL occurrences of the boilerplate name with you project name.** 

**This is necessary to ensure existing boilerplate modules are found and associated with your project!**

The easiest way to update all instances of the project name in the boilerplate is to use an IDE, such as VS Code (recommended). 

Simply use the find and replace tool and replace `go-api-boilerplate` with your `<project name>`.

## Local Setup
Ensure that you have started your local Postgres service and identified or created a database that you wish to use.

You will need to create a `.env` file containing the following environment variables:
```
DB_CLIENT=postgres
DB_USER=<Your DB User>
DB_PASSWORD=<Your DB Password>
DB_NAME=<Your Local Postgres DB>
DB_HOST=<Your DB Host>
DB_PORT=<Your DB Port (5432 for postgres)>
```

### Using a Database Other Than Postgres
If you wish to use a relational database other than postgres, you will need to update the `DB_CLIENT` environment variable and ensure that the connection string in the [config/database.go](./config/database.go) file is formatted correctly for the service. 

The necessary driver will need to be added to the config as well. 

Lastly, if you are using Docker, you will need to update the [docker-compose.yml](docker-compose.yml) db service with the appropriate image and supporting attributes. 

## Running the App
### With Local Distribution
This repository contains a shell script that is useful for running the app with your local Golang distribution. 

    $ ./run.sh

This script will install dependencies identified in your `go.mod` file and run your app
with `main.go` as the default Go entry point.
### With Docker
To run the application with Docker, first build the API image
   
    $ docker build -t katochojiro/<project name> .

Additionally, you must make the following changes to the project's [docker-compose.yml](./docker-compose.yml) file:
* Add/change the appropriate environment variables if you are not using the default docker postgres credentials
* Update the api service image with the name of your build image if it differs from your project name
  * `katochojiro/go-api-boilerplate:latest` --> `katochojiro/<build image>:latest`
  * See instructions for [cloning the repository](#cloning-the-repository) for details on globally replacing the boilerplate name with your project name.

Now, you may bring up the project stack with the command

    $ docker-compose up

And bring down the stack with the command

    $ docker-compose down

## Running Tests

Per the golang documentation, tests in Golang can be run with the command `go test` which automates the execution of any function of the form:
```
func TestXxx(*testing.T)
```

Convention in this project is to add "_test" to the filename before the file extension to organize tests. See the `models/` folder for an example test file with the start of a mock database interface.

Additional sample commands are as follows: 
```
go test -run ''      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run sub-tests matching "A=".
go test -run /A=1    # For all top-level tests, run sub-tests matching "A=1".
```

## Built With
- [Golang](https://golang.org) - Fast, statically typed, compiled language that feels like a dynamically typed, interpreted language
- [Docker](https://docker.com) - Container platform
- [PostgreSQL](https://www.postgresql.org) - Leading open source relational database
- [Git](https://git-scm.com/) - Open source version control system
- [gorilla/mux](https://www.gorillatoolkit.org/pkg/mux) - HTTP request multiplexer
- [godotenv](https://godoc.org/github.com/joho/godotenv) - Go port of the ruby dotenv library
