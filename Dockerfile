# ---- Stage 1: build the React app ----
FROM node:22-alpine AS frontend
RUN corepack enable
WORKDIR /frontend

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY frontend/ .
RUN pnpm build               # outputs to /frontend/dist

# ---- Stage 2: build the Go binary ----
FROM golang:1.24-alpine AS backend
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/api/

# ---- Stage 3: runtime ----
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=backend /app/main .
COPY --from=frontend /frontend/dist ./frontend/dist

ENV ADDR=8080 \
    ENV=production

EXPOSE 8080
CMD ["./main"]
