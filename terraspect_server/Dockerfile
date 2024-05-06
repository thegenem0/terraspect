FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN apk update && apk add --no-cache git \
    && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM golang:1.22-alpine

LABEL maintainer="Gergo Nemeth <nemethgergo02@gmail.com>"

WORKDIR /app

# Run the image as a non-root user
RUN adduser -D nonroot && \
    chown -R nonroot:nonroot /app

USER nonroot

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]
