## Build

# Base image, golang 1.18
FROM golang:latest as Build

# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

RUN  go build -o /challenge-server ./application/web/ 

## Deploy
FROM alpine:latest 

WORKDIR /

COPY --from=build /challenge-server /challenge-server 

EXPOSE 8080

ENTRYPOINT ["./challenge-server"]







