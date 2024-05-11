FROM golang:1.21.9-alpine3.18 AS builder

COPY . /github.com/kirillmc/trainings_server/source/
WORKDIR /github.com/kirillmc/trainings_server/source/

RUN go mod download
RUN go build -o ./bin/trainings_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kirillmc/trainings_server/source/bin/trainings_server .
COPY .env .
CMD ["./trainings_server"]



#COPY .env /github.com/kirillmc/trainings-auth/source/bin/trainings_auth_server