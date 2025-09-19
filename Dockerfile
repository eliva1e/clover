FROM golang:alpine AS builder
WORKDIR /usr/src/app

COPY . .

RUN go mod download
RUN go build -o clover

FROM alpine
WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/clover .

EXPOSE 80

CMD ["./clover"]
