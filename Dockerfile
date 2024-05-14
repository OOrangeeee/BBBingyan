# 使用官方的Golang镜像作为构建环境
FROM golang:1.22-alpine as builder

# 设置工作目录
WORKDIR /app

# 添加Go Modules文件
COPY go.mod .
COPY go.sum .

# 下载所有依赖项
RUN go mod download

# 安装时区数据
RUN apk add --no-cache tzdata

# 将源代码添加到工作目录
COPY . .

COPY configs/ /app/configs/

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/BBBingyan/main.go

# 使用scratch作为最小运行环境
FROM scratch

# 从builder镜像中复制构建的二进制文件
COPY --from=builder /app/main /main

# 从builder镜像中复制时区数据
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/configs/ /configs/

# 设置时区环境变量
ENV TZ=Asia/Shanghai

# 暴露端口
EXPOSE 714

# 运行应用程序
CMD ["./main"]