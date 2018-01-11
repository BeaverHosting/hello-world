FROM golang:1.9.2
ENV customername unknown
WORKDIR /go/src/app
COPY main.go .
CMD go run main.go -customer=${customername}
