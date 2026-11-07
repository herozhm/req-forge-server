# ===== 后端 req-forge-server 多阶段构建 =====
# 阶段 1:构建(编译 Go 二进制)
FROM golang:1.22-alpine AS build
WORKDIR /app

# 先复制 go.mod/go.sum,利用 Docker 层缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并编译(纯 Go 的 glebarez/sqlite,无 CGO 依赖)
COPY . .
RUN CGO_ENABLED=0 go build -o req-forge-server .

# 阶段 2:运行(最小 alpine 镜像)
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/req-forge-server .
EXPOSE 8080
CMD ["./req-forge-server"]
