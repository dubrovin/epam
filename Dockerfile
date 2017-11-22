FROM golang:latest
RUN mkdir /epam
COPY . /epam
#ADD . /app/
WORKDIR /epam
RUN go get
RUN go build
CMD ["/epam/main"]