# Sử dụng Golang image để build ứng dụng
FROM golang:1.23.2 AS builder

# Thiết lập thư mục làm việc và sao chép các file cần thiết
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build ứng dụng Go
RUN go build -o main .

# Image chạy ứng dụng (debian image nhẹ hơn)
FROM debian:bookworm-slim

# Cài đặt thư viện cần thiết để chạy ứng dụng Go
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Sao chép các file đã build từ builder image
WORKDIR /root
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

# Copy chứng chỉ SSL vào container, Nếu bạn sử dụng https
COPY /etc/letsencrypt/live/goparking.duckdns.org/fullchain.pem /etc/ssl/certs/fullchain.pem
COPY /etc/letsencrypt/live/goparking.duckdns.org/privkey.pem /etc/ssl/private/privkey.pem

# Expose cổng 8080
EXPOSE 8080

CMD ["./main"]