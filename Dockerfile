FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . /app

RUN go install github.com/vektra/mockery/v2@v2.23 && mockery --all
RUN go test -v ./...&& CGO_ENABLED=0 go build -o /app/binfile ./cmd/blossom/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/binfile .
RUN apk add --no-cache --update ffmpeg

CMD ["./binfile"]