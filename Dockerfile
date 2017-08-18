FROM golang

# ENV GOPATH /Users/JS/Developer/go
# ENV PATH $GOPATH:$PATH
WORKDIR /go/src/github.com/nlimpid/rss

ADD . /go/src/github.com/nlimpid/rss

#RUN go get -u github.com/golang/dep/cmd/dep
#RUN /Users/JS/Developer/go/bin/dep ensure

RUN go build -o myapp


ENTRYPOINT ["./myapp"]

EXPOSE 6334
