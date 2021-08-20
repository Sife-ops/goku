FROM golang:alpine

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app .

EXPOSE 80
CMD ["./app"]
