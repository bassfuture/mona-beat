#!/bin/bash

# 设置环境变量
export GO_ENV=${GO_ENV:-development}
export PORT=${PORT:-8080}
export DB_HOST=${DB_HOST:-localhost}
export DB_PORT=${DB_PORT:-5432}
export DB_USER=${DB_USER:-postgres}
export DB_PASSWORD=${DB_PASSWORD:-password}
export DB_NAME=${DB_NAME:-nft_capture_game}
export DB_SSLMODE=${DB_SSLMODE:-disable}

echo "Starting NFT Capture Game Backend..."
echo "Environment: $GO_ENV"
echo "Port: $PORT"
echo "Database: $DB_HOST:$DB_PORT/$DB_NAME"

# 构建应用
echo "Building application..."
go build -o bin/api ./cmd/api

# 运行应用
echo "Starting server..."
./bin/api