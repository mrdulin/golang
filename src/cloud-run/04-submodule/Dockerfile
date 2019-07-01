FROM golang:1.12 as builder

WORKDIR /go/src/cloud-run/04-submodule
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o submodule

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/cloud-run/04-submodule /submodule

CMD ["/submodule"]