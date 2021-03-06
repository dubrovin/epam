FROM golang:latest

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY . /go/src/app

RUN go-wrapper download

RUN go-wrapper install

CMD ["go-wrapper", "run", "-service=reserver", "-addr=0.0.0.0:8082", "-db_addr=db:8080"]

EXPOSE 8081