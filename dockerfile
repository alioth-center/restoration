# 使用golang:1.21作为构建镜像
FROM golang:1.21 as builder

# 设置工作目录
WORKDIR /app

# 将当前目录文件拷贝到容器中
COPY . .

# 安装依赖
RUN go mod tidy

# 禁用CGO并构建应用程序
RUN CGO_ENABLED=0 go build -o restoration_server

# 使用alpine作为运行镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将构建的二进制文件拷贝到运行容器中
COPY --from=builder /app/restoration_server /app/restoration_server

# 设置环境变量
ENV LOG_FILE="/app/data/restoration_collection.log"

# 暴露28081端口
EXPOSE 28081

# 运行二进制文件
ENTRYPOINT ["/app/restoration_server"]
