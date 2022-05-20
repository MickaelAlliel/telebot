FROM golang:1.18.2-alpine3.15 as builder

WORKDIR /app

RUN apk --no-cache add make bash git gcc libtool musl-dev ca-certificates dumb-init 

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY test.sh .
COPY build.sh .
COPY cmd/ cmd/
COPY internal/ internal/

RUN ./test.sh
RUN ./build.sh

# RELEASE
FROM alpine:3.15 as release 

WORKDIR /app
RUN apk add ca-certificates

COPY --from=builder /app/parser /app/parser

EXPOSE 3000
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["/app/parser"]