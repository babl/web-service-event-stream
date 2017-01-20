FROM alpine:3.4
EXPOSE 7000
COPY /static /static
ADD web-service-event-stream_linux_amd64 /bin/web-service-event-stream
WORKDIR /
CMD ["web-service-event-stream"]
