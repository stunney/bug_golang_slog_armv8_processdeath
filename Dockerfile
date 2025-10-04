# Stage 1: Builder
FROM arm64v8/golang:1.25.1 AS builder

WORKDIR /

COPY ./go.mod ./go.sum ./main.go ./
RUN go mod download

ENV CC="aarch64-linux-gnu-gcc"
ENV GOOS="linux"
ENV GOARCH="arm64"
ENV CGO_ENABLED="0"
ENV GOARM=8

RUN go build -o ./main.exe ./main.go
RUN chmod +x ./main.exe

# Stage 2: Final Image
FROM arm64v8/alpine:3.22.1
WORKDIR /

COPY --from=builder /main.exe /main.exe
RUN chmod +x /main.exe

ENV LOG_DIR="/var/log"

CMD ["/main.exe"]