# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copy go module files
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Run tidy to resolve any missing module declarations, then build
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o api ./backend/cmd/api

# Run stage
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api", "-port", "8080"]
