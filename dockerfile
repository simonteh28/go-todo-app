FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go get .
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]


