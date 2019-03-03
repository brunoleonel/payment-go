FROM golang:1.11.5-alpine3.9

RUN set -x \
    apk update && apk upgrade \
    && apk add --no-cache bash git openssh \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/kataras/iris \
    && go get github.com/jinzhu/gorm \
    && go get github.com/jinzhu/gorm/dialects/mysql

WORKDIR $GOPATH/src/github.com/brunoleonel/payment

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

# Run the executable
CMD ["payment"]