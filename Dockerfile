FROM golang:1.10-alpine as builder
ADD . /go/src/github.com/adolphlwq/tinyurl
RUN cd /go/src/github.com/adolphlwq/tinyurl && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tinyurl .

FROM scratch
COPY --from=builder /go/src/github.com/adolphlwq/tinyurl/tinyurl /
COPY --from=builder /go/src/github.com/adolphlwq/tinyurl/defult.properties /
# ENTRYPOINT ["/main"]
CMD ["/tinyurl"]