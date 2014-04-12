TARGETS := github.com/meatballhat/logrus-formatters/l2met

all: clean test save

test: build fmtpolice coverage.html

coverage.html: coverage.out
	go tool cover -html=$^ -o $@

coverage.out:
	go test -covermode=count -coverprofile=$@ -x -v $(TARGETS)
	go tool cover -func=$@

build: deps
	go install -x -v $(TARGETS)
	go build -x -v $(TARGETS)

deps:
	godep restore

clean:
	go clean -x $(TARGETS)
	rm -vf coverage.html coverage.out

save:
	godep save -copy=false $(TARGETS)

fmtpolice:
	set -e; for f in $(shell git ls-files '*.go'); do gofmt $$f | diff -u $$f - ; done

.PHONY: all build clean deps serve test fmtpolice
