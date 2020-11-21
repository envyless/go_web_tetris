FROM golang:1.15-alpine

#use git
RUN apk add --update git

RUN mkdir /hello
COPY ./app /hello
CMD ["go", "get", "."]
CMD ["go", "run", "/hello/main.go"]