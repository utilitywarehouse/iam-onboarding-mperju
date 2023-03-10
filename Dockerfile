FROM golang:alpine
COPY main.go .
COPY Makefile .
RUN GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:latest
COPY --from=0 /go/main ./

CMD ["./main"]

