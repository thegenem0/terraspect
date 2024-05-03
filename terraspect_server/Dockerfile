FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN apk update && apk add --no-cache git \
    && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM golang:1.22-alpine

WORKDIR /var/task

COPY --from=builder /app/main .

CMD [ "./main" ]
