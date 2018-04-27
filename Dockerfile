FROM alpine:latest

RUN apk add --no-cache bash curl wget
RUN addgroup -g 1000 -S www-user && \
    adduser -u 1000 -S www-user -G www-user

ENV MY_VERSION=0.0.1

ADD demo-istio /
USER www-user
CMD ["./demo-istio"]