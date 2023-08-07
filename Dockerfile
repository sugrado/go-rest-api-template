FROM golang:1.20-alpine as builder
WORKDIR /app
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o tama cmd/tama/main.go

FROM alpine:latest
COPY --from=builder /app/tama /
COPY --from=builder /go/bin/dlv /
EXPOSE 3008 40000
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec","/tama"]
