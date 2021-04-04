# placeAttractionsApi

placeAttractionsApi

documentation: <https://documenter.getpostman.com/view/3827865/TVmJheUN>

## Requirement

- install go

## Testing and run

```zsh
// install deps
$ go mod tidy

// use go start server
$ go run main.go

// use air start server
$ air

open localhost:3000

// run test case
$ cd src/test
$ go test -v

// format code
$ go fmt
```

## Docker

```zsh
// build images and start container in one line
docker-compose up -d --build

// go inside container
docker exec -it <containerId> /bin/bash

// check container logs
docker logs <containerId>

// remove and stop container
docker-compose down
```

open localhost:3000
