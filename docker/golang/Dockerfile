FROM --platform=linux/x86_64 golang:1.18

RUN go install github.com/cosmtrek/air@latest

RUN mkdir /go/src/api
WORKDIR /go/src/api
COPY ./api .
RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]