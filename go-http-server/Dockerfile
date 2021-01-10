FROM golang:rc-alpine3.12 AS builder

WORKDIR $GOPATH

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/challenge05 src/challenge05/main.go

FROM scratch

COPY --from=builder /go/bin/challenge05 /go/bin/challenge05

ENTRYPOINT ["/go/bin/challenge05"]
