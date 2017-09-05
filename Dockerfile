FROM golang

WORKDIR /go/src/github.com/nlimpid/rss

ADD . /go/src/github.com/nlimpid/rss

RUN go build -o rss

ENTRYPOINT ["./rss"]

EXPOSE 6334
