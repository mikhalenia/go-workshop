FROM alpine:3.7

ENV MC_VERSION 2018-04-28T00-08-20Z

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache --virtual .build-deps curl
RUN curl "https://dl.minio.io/client/mc/release/linux-amd64/archive/mc.RELEASE.$MC_VERSION" > /usr/bin/mc
RUN chmod +x /usr/bin/mc && apk del .build-deps

COPY config.json /root/.mc/

COPY entrypoint.sh /opt/bin/entrypoint.sh
RUN chmod +x /opt/bin/entrypoint.sh

ENTRYPOINT [ "sh", "-c" ]
CMD ["/opt/bin/entrypoint.sh"]