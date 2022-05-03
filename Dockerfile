ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /go/src/

COPY . .

ENV CGO_ENABLED=0
ENV GO_OSARCH="linux/amd64"

RUN go build -o /go/bin/binary /go/src/cmd/api/main.go

FROM gcr.io/distroless/base

COPY --from=builder /go/bin/binary /go/bin/binary
CMD ["/go/bin/binary"]
