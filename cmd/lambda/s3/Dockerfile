FROM alpine:3.7

ENV MINIO_UPDATE off
ENV MINIO_ACCESS_KEY workshop
ENV MINIO_SECRET_KEY workshop
ENV MINIO_PORT 8080
ENV MINIO_VERSION 2018-05-11T00-29-24Z

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache --virtual .build-deps curl
RUN echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf
RUN curl "https://dl.minio.io/server/minio/release/linux-amd64/archive/minio.RELEASE.$MINIO_VERSION" > /usr/bin/minio
RUN chmod +x /usr/bin/minio && apk del .build-deps

COPY config.json /root/.minio/

COPY entrypoint.sh /opt/bin/entrypoint.sh
RUN chmod +x /opt/bin/entrypoint.sh

EXPOSE $MINIO_PORT

ENTRYPOINT [ "sh", "-c" ]
CMD [ "/opt/bin/entrypoint.sh" ]
