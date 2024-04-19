FROM quay.io/projectquay/golang:1.20 as builder

ARG TARGET_OS=linux

RUN echo $TARGET_OS
WORKDIR /go/src/app
COPY . .
RUN make goget
RUN make $TARGET_OS

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/gbot .
COPY --from=builder /go/src/app/.env .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["./gbot", "start"]

#docker run --env-file ./.env  gbot start