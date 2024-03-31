FROM golang:1.18

WORKDIR /app

COPY . .

#RUN go mod init fizzbuzz
RUN go build -o fizzbuzz .

CMD ["./fizzbuzz"]
