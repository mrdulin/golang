FROM golang:1.12 as builder

WORKDIR /go/src/cloud-run/05-connecting-to-cloud-sql
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \ 
  && CGO_ENABLED=0 GOOS=linux go build -v -o app

# stage 2
FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/cloud-run/05-connecting-to-cloud-sql/app .

EXPOSE 8080
CMD ["./app"]