# first stage
FROM golang:1.21.3-alpine as builder
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /main apps/sip-server/cmd/main.go
# second stage
FROM alpine:latest
COPY --from=builder /main .
CMD ["./main"]
EXPOSE 5000
EXPOSE 6000