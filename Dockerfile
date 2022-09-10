FROM golang:1.18.5-bullseye

WORKDIR /go-host
COPY . /go-host

RUN apt-get update
RUN go mod download 

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]

EXPOSE 3000