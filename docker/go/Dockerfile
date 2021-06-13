FROM golang:1.16.5-alpine3.13

COPY . /go/src/

WORKDIR /go/src/

RUN apk update \
  && apk add --no-cache git \
  && go get github.com/gin-gonic/gin \
  && go get -u github.com/codegangsta/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get -tags 'mysql' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
#   # && go mod init \
#   # && go mod tidy

EXPOSE 3001
CMD ["gin", "-i", "run"]
