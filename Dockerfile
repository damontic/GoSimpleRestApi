FROM golang:1.8

WORKDIR /go/src/GoSimpleRestApi
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["GoSimpleRestApi"]