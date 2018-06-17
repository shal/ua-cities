FROM golang:1.10
MAINTAINER Ali Shanaakh <ashanaakh@gmail.com>

ENV PORT=8080

WORKDIR $GOPATH/src/github.com/ashanaakh/ua-cities

COPY . ./

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o app

EXPOSE $PORT

CMD ["./app"]