#!/bin/bash

# 获取当前目录
PROJECT_DIR=$(dirname "$0")
# 生成目录
GENERATE_DIR="$PROJECT_DIR/cmd/generate"

# 打开目录 目录不存在时 退出
cd "$GENERATE_DIR" || exit

echo "Start Generating"
# 生成, 执行gen文件
go run .
