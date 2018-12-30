FROM golang:1.10
LABEL MAINTAINER Ali Shanaakh <github@shanaakh.pro>

ENV PORT=8080

WORKDIR $GOPATH/src/github.com/shal/ua-cities

COPY . ./

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o app

EXPOSE $PORT

CMD ["./app"]