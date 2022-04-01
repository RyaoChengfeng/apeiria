FROM golang:1.15.11-alpine3.13 AS builder
COPY ./src /src
WORKDIR /src
ENV GOPROXY=https://goproxy.cn GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app  .

FROM alpine:3.13
COPY ./env/config /env/config
COPY --from=builder /build/app /bin/app
ENTRYPOINT ["sh","-c","/bin/app"]
