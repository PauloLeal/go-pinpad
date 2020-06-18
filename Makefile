COVER_PROFILE_FILE=/tmp/goemv-go-cover.tmp

target: test

test:
	go test -v -cover -count 1 -failfast ./...
	gopherbadger -md="README.md" -png=false > /dev/null
	rm coverage.out

cover-html:
	go test -coverprofile=${COVER_PROFILE_FILE} ./...
	go tool cover -html=${COVER_PROFILE_FILE}
