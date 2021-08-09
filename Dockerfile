FROM golang:1.16 as builder
RUN mkdir -p /go/src/test-adv-tech
ADD . /go/src/test-adv-tech
WORKDIR /go/src/test-adv-tech
RUN go env -w GO111MODULE=on
RUN go mod download
RUN go mod tidy
RUN go mod vendor
RUN go build -a -o test-adv-tech

FROM alpine:3.14
COPY --from=builder /go/src/test-adv-tech/test-adv-tech .

ENTRYPOINT [ "./test-adv-tech" ]
