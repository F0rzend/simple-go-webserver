ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GO_OSARCH="linux/amd64"
RUN go build -o ./app .

FROM gcr.io/distroless/base:latest

COPY --from=builder /build/app /app

CMD ["/app"]

