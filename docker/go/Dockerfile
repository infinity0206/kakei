FROM golang:1.10.3-alpine3.8

COPY src/api /go/src/api/

WORKDIR /go/src/api/

RUN apk update \
  && apk add --no-cache git \
  && go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get -tags 'mysql' -u github.com/golang-migrate/migrate/v4/cmd/migrate/ \
  && go mod init \
  && go mod tidy

EXPOSE 3001