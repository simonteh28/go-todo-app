FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY . .

RUN go get ./cmd
RUN go build -o bin ./cmd

ENTRYPOINT [ "/app/bin" ]


