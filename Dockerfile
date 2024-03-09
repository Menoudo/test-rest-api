FROM golang:1.22.1-alpine3.19
EXPOSE 8080
RUN apk add --update git; \
    apk add curl; \
    mkdir -p ${GOPATH}/test-rest-api-go
WORKDIR ${GOPATH}/test-rest-api-go/
COPY test-rest-api.go ${GOPATH}/test-rest-api-go/
COPY go.* ${GOPATH}/test-rest-api-go/
RUN cd ${GOPATH}/test-rest-api-go/ && \
    go build -o test-rest-api-go .
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.56.2
# TODO Run linter

#
FROM golang:1.22.1-alpine3.19
LABEL vendor=Menoudo\ Software \
      com.example.is-production="Yes" \
      com.example.version="1.00.00" \
      com.example.release-date="2024-03-09"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/test-rest-api-go/test-rest-api-go .
CMD [ "/app/test-rest-api-go" ]
