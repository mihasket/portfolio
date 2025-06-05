FROM golang:1.24-alpine AS BUILDER
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/web/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/ui ./ui

EXPOSE 8080
CMD ["./main", "-addr=:8080"]
