# syntax=docker/dockerfile:1  ## optional,dockerfile:1 always points to the last of the version

# base image we would like to use
FROM golang:1.16-alpine

# destination for all subsequent commands
WORKDIR /app

# copy the go.mod and go.sum file into our project directory /app which,
# owing to our use of WORKDIR, is the current directory 
# COPY go.mod ./
# COPY go.sum ./
# Go modules (current application) will be installed into a directory inside the image.
# Copy our source code into the image
COPY . ./

RUN go mod download

## compile an application, result: static application binary named docker-gs-ping
RUN go build -o library cmd/library/main.go

EXPOSE 8000

# command to execute when our image is used to start a container
CMD [ "/app/library" ]