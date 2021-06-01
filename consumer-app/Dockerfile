FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN ls
ADD .builds/consumer-app /consumer-app
EXPOSE 3000
ENTRYPOINT ["./consumer-app"]
