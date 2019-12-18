FROM golang:1.10-alpine as builder
ADD . /go/src/github.com/tinyurl/tinyurl
RUN cd /go/src/github.com/tinyurl/tinyurl && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tinyurl .

FROM scratch
COPY --from=builder /go/src/github.com/tinyurl/tinyurl/tinyurl /
COPY --from=builder /go/src/github.com/tinyurl/tinyurl/default.properties /
# ENTRYPOINT ["/main"]
CMD ["/tinyurl"]