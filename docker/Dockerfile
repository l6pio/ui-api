FROM golang:1.16 as builder
ENV GOPROXY=https://goproxy.cn
ADD . /src
WORKDIR /src
RUN make all

FROM alpine:3.13 as release
COPY --from=builder /src/releases/api /api
COPY docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh" ]
