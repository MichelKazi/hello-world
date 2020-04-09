FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

RUN go get -u github.com/cosmtrek/air

EXPOSE 8080

CMD air
