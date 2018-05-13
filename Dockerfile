FROM golang:1.8

WORKDIR /go/src/Go4Chan
COPY . .


RUN go get -d -v ./...
RUN go install -v ./...


RUN go install
CMD ["go", "run", "http1.go"]
