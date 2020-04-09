FROM golang:1.14

WORKDIR /usr/src/app

copy go.mod go.sum ./

run go mod download

copy . .

run CGO_ENABLE=0 GOOS=linux go build

CMD go get github.com/cosmtrek/air && air

EXPOSE 8080
