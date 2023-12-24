#FROM golang:1.21.5-bullseye - debian/ubuntu has build dependencies(gcc)
FROM golang:1.21.5-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

#Install the build dependencies
RUN apk add --no-cache gcc libc-dev

#CGO_ENABLED=1 for go-sqlite3 to work
RUN CGO_ENABLED=1 go build -o main .

# Volume for SSL/TLS Certificates(optional for development)
VOLUME /certs

EXPOSE 3000

CMD ["/app/main"]
