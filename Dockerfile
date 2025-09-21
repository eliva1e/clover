FROM golang:alpine AS builder
WORKDIR /usr/src/app

COPY . .

RUN go mod download
RUN go build -v -o clover ./cmd/clover/main.go

FROM alpine
WORKDIR /clover

COPY --from=builder /usr/src/app/clover .

EXPOSE 80 443
CMD ["./clover"]
