# Compile stage
FROM golang:1.17 AS build-env

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

EXPOSE 8080

# Command to run
ENTRYPOINT ["/main"]