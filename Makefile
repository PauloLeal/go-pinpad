COVER_PROFILE_FILE=/tmp/goemv-go-cover.tmp

target: test

t:
	go test -v -cover -count 1 ./...

c:
	go test -coverprofile=${COVER_PROFILE_FILE} ./...
	go tool cover -html=${COVER_PROFILE_FILE}

badge:
	gopherbadger -md="README.md" -png=false > /dev/null
	rm coverage.out

test: t badge

cover-html: c badge
