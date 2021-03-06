FROM golang:1.11-alpine AS builder

ENV CGO_ENABLED 0

WORKDIR /go/apiserver

RUN apk add --no-cache git

COPY . .

RUN go install -v gitlab.com/404busters/inventory-management/apiserver/cmd/apiserver

FROM gcr.io/distroless/base

ENV PORT 8080

COPY --from=builder /go/bin/apiserver /
COPY assets /assets

CMD ["/apiserver"]
