FROM golang:1.21.3-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o myapp .

FROM scratch AS runner
COPY --from=builder /app/myapp /myapp
EXPOSE 8080
ENTRYPOINT ["/myapp"]