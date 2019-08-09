.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/scrape scrape/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --force

deploy.prod: clean build
	sls deploy --verbose --force --stage production

run: clean build
	sls invoke local -f scrape
