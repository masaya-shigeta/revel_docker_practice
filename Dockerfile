FROM golang:1.11.0

ENV GOPATH $GOPATH:/go/src

RUN apt-get update && \
    apt-get upgrade -y

# install revel and revel-cli
RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

# change current dir
RUN mkdir /go/src/myapp

EXPOSE 9000
