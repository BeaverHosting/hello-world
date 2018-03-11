FROM golang:1.10-alpine
EXPOSE 8080
WORKDIR /go/src/app
COPY main.go .
CMD go run main.go -customer="${MY_CUSTOMER_NAME}" -pod="${MY_POD_NAME}"