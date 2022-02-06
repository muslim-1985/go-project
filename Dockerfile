# Compile stage
FROM golang:1.17 AS build-env

# Set necessary environmet variables needed for our image
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64

WORKDIR /go/src/myapp

COPY . ./
RUN go mod download

RUN go build -o /go/src/myapp/go_pr

EXPOSE 8080

# Command to run
CMD ["/go/src/myapp/go_pr"]
