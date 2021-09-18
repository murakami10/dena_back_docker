FROM golang:1.16.8-buster

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go get -u github.com/cosmtrek/air
RUN go mod download

# とりあえず起動したらサーバ動いているように
# buildしてもいいかも
CMD ["go", "run", "main.go"]