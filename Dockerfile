FROM golang:1.11

WORKDIR /go/src/Groot
COPY . .

RUN go get -d -v ./... Groot
RUN go install -v ./... Groot

CMD ["Groot"]
