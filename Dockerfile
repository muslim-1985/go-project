# Compile stage
FROM golang:1.18.2 AS build-env

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /var/www

COPY . ./
RUN go mod download

RUN go build go_project/cmd/app

EXPOSE 8080

# Command to run
CMD ["/var/www/go_project/app"]
