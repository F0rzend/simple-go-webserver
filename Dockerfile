FROM golang:1.19-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GO_OSARCH="linux/amd64"
RUN go build -o ./binary app/main.go

FROM gcr.io/distroless/base:latest

COPY --from=builder /build/binary /app

CMD ["/app"]
