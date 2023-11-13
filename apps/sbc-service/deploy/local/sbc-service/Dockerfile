# first stage
FROM golang:1.21.3-alpine as builder
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /main apps/sbc-service/cmd/main.go
# second stage
FROM alpine:latest
COPY --from=builder /main .
CMD ["./main"]
EXPOSE 8000