build-user:
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o user-app/.builds/user-app user-app/cmd/main.go

build-consumer:
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o consumer-app/.builds/consumer-app consumer-app/cmd/main.go 

nuke:
	rm -rf consumer-app/.builds && rm -rf user-app/.builds
