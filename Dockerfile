ARG GO=golang:1.21-alpine

# Step 1: Dependencies caching
FROM ${GO} as deps
WORKDIR /modules
COPY go.mod go.sum .
RUN go mod download

# Step 2: Builder
FROM ${GO} as build
COPY --from=deps /go/pkg /go/pkg
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o binfile ./cmd/blossom


# Step 3: Final
FROM alpine
COPY --from=build /app/binfile /binfile
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN apk add --no-cache --update ffmpeg
CMD ["/binfile"]
