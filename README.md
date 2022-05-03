# SimpleGoWebserver

![GitHub release (latest by date)](https://img.shields.io/github/v/release/F0rzend/SimpleGoWebserver?display_name=tag)
![GitHub top language](https://img.shields.io/github/languages/top/F0rzend/SimpleGoWebserver)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/F0rzend/SimpleGoWebserver)
![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2FF0rzend%2FSimpleGoWebserver%2Fmaster%2Fdocs%2Fopenapi.yaml)

There is a simple webserver written in Go.

We use DDD (Domain Driven Design) to separate the domain logic from the application logic.

## Features

 * The project uses the structure proposed in the book
[Go with domain](https://threedots.tech/go-with-the-domain/)
 * [moq](https://github.com/matryer/moq) for adapters mocking
 * [Swagger](https://swagger.io)
 * [OpenAPI](https://openapi.io)
 * [JSON Web Token](https://jwt.io)
 * [chi router](https://go-chi.io/)
 * [Docker](https://www.docker.com/)
 * [docker-compose](https://docs.docker.com/compose/overview/)
 * [GitHub Actions](https://github.com/F0rzend/SimpleGoWebserver/actions)
 

## Run in docker *(Recommended)*

It's simple to run it with docker.
You don't even need to clone the project to your repository:
```shell
docker run \
  -e ADDRESS=0.0.0.0:8080 \
  ghcr.io/f0rzend/simplegowebserver:master
```

There is also some environment variables you can set, to configure project.
Read more about it in the [docker documentation](https://docs.docker.com/engine/reference/run/#env-environment-variables).

List of available environment variables above.

| Variable  | Description          | Default |
|-----------|----------------------|---------|
| `ADDRESS` | Address to listen on | `:8080` |

Be carefully. If you want to run it on localhost in docker container, 
you need to set `ADDRESS` to `0.0.0.0:8080`.

## Build from source

If you want to build the project from source, you need to have
[Go](https://golang.org/doc/install) and [git](https://git-scm.com/downloads) installed.

Firstly you need to clone the project repository on your local machine:
```shell
git clone https://github.com/F0rzend/SimpleGoWebserver.git
```

Entry point is `cmd/api/main.go` file.
Then you need to go to the project directory and build it:
```shell
cd SimpleGoWebserver && go build -o simplegowebserver cmd/api/main.go
```

Now you can run the project. You must be in the same folder as the binary
and run the following command:
```shell
./simplegowebserver
```

After this you will see the following message:
```shell
0:00AM INF cmd/api/main.go:27 > starting server on :8080
```

## Run project via docker compose

You can run the project via docker compose.
See the [docker-compose.yml](https://github.com/F0rzend/SimpleGoWebserver/blob/master/docker-compose.yml) file.