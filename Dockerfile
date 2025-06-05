# Dockerfile
#FROM golang:1.24-alpine AS builder
#WORKDIR /app
#COPY . .
#RUN go build -o main ./cmd/web/

#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=builder /app/main .
#EXPOSE 4000
#CMD ["./main"]

FROM golang:1.24

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/web/

EXPOSE 8080

CMD ["./main", "-addr=:8080"]
