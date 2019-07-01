FROM golang:1.12 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \
  && CGO_ENABLED=0 GOOS=linux go build -v -o sls-fns-go

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/app/sls-fns-go /sls-fns-go

CMD ["/sls-fns-go"]