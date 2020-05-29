FROM golang:1.14.3-alpine as builder

RUN apk add --update make

WORKDIR /app
ADD . .
RUN go get -d -v ./...

RUN make build


FROM alpine:3.11.3

WORKDIR /app

COPY --from=builder /app/bin/lainoa /bin/

ENTRYPOINT ["lainoa"]
