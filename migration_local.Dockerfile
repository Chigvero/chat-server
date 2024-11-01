FROM alpine:3.20

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

WORKDIR /root/

RUN mkdir -p bin migrations


ADD https://github.com/pressly/goose/releases/download/v3.22.1/goose_linux_x86_64 ./bin/goose
RUN chmod +x ./bin/goose

COPY migration_local.sh .
COPY local.env .
ADD migrations/*.sql migrations/
RUN chmod +x migration_local.sh

ENTRYPOINT ["bash","migration_local.sh"]