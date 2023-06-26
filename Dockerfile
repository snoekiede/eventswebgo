FROM golang:1.20-alpine as builder
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o eventsweb main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app
COPY --from=builder /app/eventsweb .
EXPOSE 8080
CMD ["./eventsweb"]
