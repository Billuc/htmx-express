FROM golang:1.21-alpine

WORKDIR /app

RUN apk add --no-cache gcc g++ libc-dev

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD [ "air", "-c", ".air-docker.toml" ]