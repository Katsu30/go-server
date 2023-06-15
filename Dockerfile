# FROM golang:1.18-alpine

# RUN apk update && apk add git
# RUN go install github.com/cosmtrek/air@v1.29.0

# CMD ["air", "-c", ".air.toml"]

FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]