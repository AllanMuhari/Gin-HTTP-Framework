FROM golang:latest

LABEL maintainer = "Quique <hello@pragmaticreviews.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . . 

ENV PORT 5000

RUN go build 

RUN find . -name "*.go" -type f -delete


EXPOSE $port

CMD ["./golang-gin-poc"]

