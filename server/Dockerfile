FROM golang

ADD httpserv.go /go
ADD main.go /go

RUN export GOPATH=/tmp
RUN go get github.com/jawher/mow.cli
RUN go get -u github.com/gorilla/mux
RUN go get github.com/GoRethink/gorethink



RUN go build -o /bin/serv
EXPOSE 8000 8000
CMD ["serv", "run"]