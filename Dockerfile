FROM golang:1.12.1-alpine
MAINTAINER Vanapagan <olddevil@vanapagan.com>
EXPOSE 8080
RUN apk add --update git; \
    mkdir -p ${GOPATH}/test-rest-api-go; \
    go get -u github.com/gorilla/mux
WORKDIR ${GOPATH}/test-rest-api-go/
COPY test-rest-api.go ${GOPATH}/test-rest-api-go/
RUN go build -o test-rest-api-go .
#
FROM alpine:latest
LABEL vendor=Vanapagan\ Software \
      com.example.is-production="No" \
      com.example.version="0.0.1-beta" \
      com.example.release-date="2018-10-31"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/test-rest-api-go/test-rest-api-go .
CMD [ "/app/test-rest-api-go" ]
