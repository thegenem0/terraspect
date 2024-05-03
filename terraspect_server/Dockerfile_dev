FROM golang:1.22-alpine AS dev

WORKDIR /app

COPY go.mod go.sum ./

RUN apk update && apk add --no-cache git

RUN go mod download

COPY . .

EXPOSE 8000

CMD ["go", "run", "."]