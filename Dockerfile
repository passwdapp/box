FROM alpine:3.13.1 AS base
EXPOSE 3000

FROM golang:1.15.8-alpine3.13 AS builder
RUN apk update
RUN apk add build-base
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o passwdbox -ldflags "-s" cmd/passwdbox/main.go

FROM base as FINAL
RUN mkdir -pv /app/data/uploads
WORKDIR /app
COPY --from=builder /build/passwdbox .
CMD [ "/app/passwdbox", "-use-dotenv=false" ]
