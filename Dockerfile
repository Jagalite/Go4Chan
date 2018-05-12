FROM golang:1.8

WORKDIR /music
COPY /music/. .

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
CMD ["go", "run", "http1.go"]
