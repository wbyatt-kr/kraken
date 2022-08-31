FROM golang:1.19.0-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o kraken .

CMD ["/app/kraken"]