FROM golang:1.12.5-alpine

#RUN apk update && apk upgrade && apk add --no-cache bash git

RUN apk --no-cache add gcc g++ make ca-certificates

#RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/inven-wms/
COPY . /go/src/inven-wms/

RUN cd /go/src/inven-wms/ && CGO_ENABLED=0 go build

WORKDIR /go/src/inven-wms/
CMD /go/src/inven-wms/inven-wms
EXPOSE 8484