FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o go-rest-api-template cmd/go-rest-api-template/main.go

FROM alpine:latest
COPY --from=builder /app/go-rest-api-template /
EXPOSE 3008
CMD ["/go-rest-api-template"]
