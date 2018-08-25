# Found here: http://jbeckwith.com/2015/05/08/docker-revel-appengine/


FROM golang
MAINTAINER Alex Standke "xanderstrike@gmail.com"

ADD . /go/src/github.com/XanderStrike/manual-upload

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

VOLUME ["/watch"]

EXPOSE 8080
ENTRYPOINT revel run github.com/XanderStrike/manual-upload dev 8080

