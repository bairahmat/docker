FROM golang
MAINTAINER jihar jihar.akakom14@gmail.com

RUN go get github.com/gorilla/securecookie
RUN go get golang.org/x/crypto/bcrypt
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/gorilla/websocket

ADD . /go/src/github.com/kanciogo/kancio-chat
WORKDIR /go/src/github.com/kanciogo/kancio-chat
RUN go build
CMD ./kancio-chat
