FROM alpine:latest
RUN mkdir -p /usr/local/go/lib/time/
ADD ./zoneinfo.zip /usr/local/go/lib/time/
RUN mkdir -p /data
ADD ./.env /data
ADD ./fantasy-user-service /data
WORKDIR /data
CMD ["./fantasy-user-service"]