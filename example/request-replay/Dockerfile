FROM golang:1.16

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/github.com/kzmake/dapr-kit

COPY go.mod .
COPY go.sum .

RUN go mod download

EXPOSE 4000

CMD ["air", "-v"]
