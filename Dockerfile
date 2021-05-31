FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .

RUN go build -o main cmd/hub/main.go 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/configs configs/
COPY --from=builder /app/cert cert/


EXPOSE 8080
CMD [ "/app/main" ]
