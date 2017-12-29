FROM golang:1.12 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \
  && CGO_ENABLED=0 GOOS=linux go build -v -o includedeps

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/app/includedeps /includedeps

CMD ["/includedeps"]