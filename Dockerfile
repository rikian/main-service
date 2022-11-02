FROM golang

WORKDIR /workdir/server-main

COPY . .

RUN go mod tidy
RUN go build -o ./server-main .

EXPOSE 8080

CMD ["./server-main"]