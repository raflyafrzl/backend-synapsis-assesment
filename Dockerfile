FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

CMD ["/app/binary"]
EXPOSE 9000