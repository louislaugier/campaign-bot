FROM golang:1.18-alpine

WORKDIR /go/src/campaign-bot

COPY . .

WORKDIR /go/src/campaign-bot/cmd

RUN go mod download

RUN go build -o app

EXPOSE 80

CMD ["./app"]