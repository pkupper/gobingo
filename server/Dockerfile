FROM golang:alpine

RUN go build .
COPY server /
CMD ["/server"]
EXPOSE 80
