# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o usersApp ./

RUN chmod +x /app/usersApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/usersApp /app

CMD [ "/app/usersApp" ]