# Backend Dockerfile
FROM golang:1.24 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o aurelia ./cmd/aurelia

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/aurelia .
EXPOSE 8080
CMD ["./aurelia"]
