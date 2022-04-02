FROM golang:1.17-alpine AS builder

WORKDIR /app

ENV GOPRIVATE github.com/aeekayy/descartes/server
COPY . .
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o ./descartes
RUN ls -latr

FROM alpine

WORKDIR /app

RUN apk add bash curl
COPY --from=builder /app/descartes /app/
# Expose is not supported by 8080
# EXPOSE 8080/tcp

# ENV PORT 8080

RUN adduser -D app
USER app

#HEALTHCHECK --interval=5m --timeout=3s \
#  CMD curl -f http://localhost:8080/ping || exit 1

CMD /app/descartes server start -p $PORT