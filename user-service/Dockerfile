# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
RUN apk add --no-cache ca-certificates
WORKDIR /src

# 利用缓存：先拷 go.mod/sum
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# 再拷源码并构建当前服务
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /out/app /app
USER nonroot:nonroot
EXPOSE 48482
ENTRYPOINT ["/app"]
