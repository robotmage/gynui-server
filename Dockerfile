FROM golang:1.16 AS builder
WORKDIR /build
COPY * /build
RUN env GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -ldflags="-s -w" -o "superserver" .
RUN echo $(ls -la /build)
RUN echo $(pwd)

FROM alpine:latest
WORKDIR /cmd/
COPY --from=builder /build/superserver ./
RUN echo $(ls -la /cmd)
EXPOSE 1986
CMD ["/cmd/superserver"] 

