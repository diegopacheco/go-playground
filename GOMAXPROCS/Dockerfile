FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -g 1001 appgroup && \
    adduser -D -u 1001 -G appgroup appuser
WORKDIR /root/
COPY --from=builder /app/main .
RUN chown appuser:appgroup main
USER appuser
CMD ["./main"]
