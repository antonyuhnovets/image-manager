FROM golang:1.17-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/imgm ./cmd/daemon

FROM alpine

COPY --from=builder /build/bin /bin

CMD ["imgm"]
