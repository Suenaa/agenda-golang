FROM golang:1.8

COPY . /go/src/github.com/Suenaa/agenda-golang
WORKDIR /go/src/github.com/Suenaa/agenda-golang/cli

RUN go-wrapper download
RUN go build -o /go/bin/agendalocal .

WORKDIR /go/src/github.com/Suenaa/agenda-golang/service

RUN go-wrapper download
RUN go build -o /go/bin/agendaserver .

CMD ["agendaserver"]

EXPOSE 8080