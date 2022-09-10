FROM golang:1.18.5-bullseye

WORKDIR /dena
COPY . /dena

RUN apt-get update
RUN go install

EXPOSE 3000

# RUN go get -u github.com/cosmtrek/air && \
#     go build -o /go/bin/air github.com/cosmtrek/air

# CMD ["air", "-c", ".air.toml"]