# syntax=docker/dockerfile:1

# setting go version
ARG GO_VERSION=1.23

# using go and alpine for image
FROM golang:${GO_VERSION}-alpine

# set default production mode
ENV GO_ENV production

# set directory 
WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

# build application
RUN go build -o app ./cmd/main.go

# using user non root for run the apps
USER nobody

# expose port
EXPOSE 8080

# run apps
CMD ["./app"]
