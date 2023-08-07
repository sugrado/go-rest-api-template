FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o tama cmd/tama/main.go

FROM alpine:latest
COPY --from=builder /app/tama /
EXPOSE 3008
CMD ["/tama"]
