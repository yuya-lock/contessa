FROM --platform=linux/x86_64 golang:1.18

RUN go install github.com/cosmtrek/air@latest

WORKDIR /go/src/app/lib/
COPY lib/go.mod .
COPY lib/go.sum .

WORKDIR /go/src/app/api/
COPY api/go.mod .
COPY api/go.sum .
RUN go mod download

WORKDIR /go/src/
COPY . ./app

WORKDIR /go/src/app/api

CMD ["air"]
