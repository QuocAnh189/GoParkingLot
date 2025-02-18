# Sử dụng Golang image để build ứng dụng
FROM golang:1.23.2 AS builder

# Thiết lập thư mục làm việc và sao chép các file cần thiết
WORKDIR /app
RUN curl -fLo /usr/local/bin/air https://github.com/air-verse/air/releases/download/v1.61.1/air_1.61.1_linux_amd64 && chmod +x /usr/local/bin/air

# Kiểm tra xem air có được cài đặt đúng không
RUN which air

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
COPY --from=builder /app/app.env .
COPY --from=builder /app/configs ./configs

# Expose cổng 8080
EXPOSE 8080

# Chạy ứng dụng bằng Air để tự động reload khi có thay đổi
#CMD ["air", "-c", "air.toml"]
CMD ["./main"]