FROM golang:1.17-alpine AS builder

WORKDIR /app

ENV GOPRIVATE github.com/aeekayy/descartes
COPY . .
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o ./descartes
RUN ls -latr

FROM alpine

WORKDIR /app

RUN apk add bash curl
COPY --from=builder /app/descartes /app/
EXPOSE 8080/tcp

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:8080/ping || exit 1

CMD [ "/app/descartes", "server", "start" ]