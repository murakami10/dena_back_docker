FROM golang:1.16.8-buster

WORKDIR /app/src

COPY ./src/go.mod ./src/go.sum ./
RUN go get -u github.com/cosmtrek/air
RUN go mod download

# とりあえず起動したらサーバ動いているように
# buildしてもいいかも
CMD ["go", "run", "main.go"]