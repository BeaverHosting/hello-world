# Beaver Hosting Product - Hello World

Simple golang program serving content as a HTTP Server on port 8080.

## Getting Started

### Prerequisites

You need golang 1.9 or later to compile the source.

## Deployment

### Local

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

### Docker

To use the app with docker user the Dockerfile, build the image with :

```
docker image build . -t bh-p-hello-world
```

Start a container with :

```
docker run --name hello-world-kevin -d -p 8080:8080 -e customername=Kevin bh-p-hello-world
```

Start from docker compose :

```
docker-compose up -d
```

Start with Kubernetes :

```
kubectl create -f env.yml
kubectl create -f deployment.yml
minikube service hello-world
```

### Result

If the server is up and running you should see this :

```
{
    app_name: "hello-world",
    customer_name: "Kevin",
    date_start: "2018-01-11T23:21:58.410500459Z",
    status: 200
}
```

