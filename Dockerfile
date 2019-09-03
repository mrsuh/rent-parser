FROM golang:1.11

COPY . /go/src/rent-parser
RUN sh /go/src/rent-parser/bin/compile.sh

EXPOSE 9080

CMD ["/go/src/rent-parser/bin/server", "/go/src/rent-parser/config/config.yml"]
