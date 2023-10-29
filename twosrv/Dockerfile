FROM golang:alpine


WORKDIR /APP

ADD . .
RUN go mod tidy

RUN go build main.go

#COPY . /APP

#RUN go build main.go

CMD /APP/main
