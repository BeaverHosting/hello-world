# Beaver Hosting Product - Hello World

Simple golang program serving content as a HTTP Server on port 8080.

## Getting Started

### Prerequisites

You need golang 1.9 or later to compile the source.

## Deployment

To start the go program :

```
go build main.go
```

Or :

```
go run main.go
```

The program needs one parameter :

```
go run main.go -customer=Test
```

If no parameter is given the default value is "Unknown".

To check if the server is up and running, go to :

```
http://localhost:8080/alive
```