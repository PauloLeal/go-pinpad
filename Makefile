TARGET_FILE=${shell head -n1 go.mod | sed -r 's/.*\/(.*)/\1/g' }
COVER_PROFILE_FILE="/tmp/$(TARGET_FILE)-go-cover.tmp"

target: test

clean:
	rm -rf $(TARGET_FILE) $(BUILD_DIR)

lint:
	@go fmt ./...

############## test tasks

test: lint
	go test -v -cover -count 1 ./...
	$(MAKE) badge

cover-html: test
	go test -coverprofile=$(COVER_PROFILE_FILE) ./...
	go tool cover -html=$(COVER_PROFILE_FILE)

badge:
	@go get github.com/jpoles1/gopherbadger/...
	gopherbadger -md="README.md" -png=false 1>&2 2> /dev/null
	@if [ -f coverage.out ]; then rm coverage.out ; fi;