FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN ls
ADD .builds/user-app /user-app
EXPOSE 9000
ENTRYPOINT ["./user-app"]
