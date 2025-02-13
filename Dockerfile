FROM golang:1.23.4-alpine AS development

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go mod tidy

COPY . .

RUN go install github.com/cespare/reflex@latest

EXPOSE 8080

CMD reflex -g '*.go' go run main.go --start-service