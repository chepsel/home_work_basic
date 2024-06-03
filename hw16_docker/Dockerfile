FROM golang:latest AS builder
COPY ./hw15_go_sql /usr/local/go/src/store
WORKDIR /usr/local/go/src/store
RUN go clean --modcache && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
-mod=readonly -o store /usr/local/go/src/store/main.go
FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/go/src/store/store /app/
CMD ["/app/store"]