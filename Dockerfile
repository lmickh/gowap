FROM golang:1.5.1-alpine

RUN apk -U upgrade \
 && apk -U add bash git curl \
 && mkdir -p /go/src/github.com/lmickh/gowap \
 && go get github.com/garyburd/redigo/redis \
 && go get github.com/benschw/dns-clb-go/clb

ADD . /go/src/github.com/lmickh/gowap

RUN go install github.com/lmickh/gowap

ENV SERVICE_NAME gowap-dev
ENV SERVICE_TAGS dev,rest
ENV REDIS_URL redis:6379

ENTRYPOINT /go/bin/gowap

EXPOSE 8080
