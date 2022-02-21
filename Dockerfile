FROM golang:1.17-alpine AS builder

WORKDIR /app

ENV GOPRIVATE github.com/aeekayy/descartes
COPY * ./
RUN go mod download

RUN go build -o ./descartes

FROM alpine

WORKDIR /app

RUN apk add bash curl
COPY --from=builder /app/descartes /app/
EXPOSE 8080

CMD [ "/app/descartes" ]