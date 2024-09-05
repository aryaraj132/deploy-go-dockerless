FROM --platform=linux/amd64 golang:1.22.5-alpine3.19

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY main.go ./

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]