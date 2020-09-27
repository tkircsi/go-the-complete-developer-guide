FROM golang:1.15.2-alpine3.12 as builder

WORKDIR /app
COPY . .

RUN go build -o main.e

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /app/main.e ./
ENTRYPOINT [ "/app/main.e" ]