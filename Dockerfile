FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o ./descartes

FROM alpine

WORKDIR /app

RUN apk add bash curl
COPY --from=builder /app/descartes /app/
EXPOSE 8080

CMD [ "/app/descartes" ]