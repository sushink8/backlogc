.PHONY: build run
RUN=./backlog -F key.txt
build:
	goimports -w *.go
	go build
gox:
	gox --osarch="linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64"

run:
	$(RUN)
test:
	go test -v

list_project:
	$(RUN) -A list_project

list_issueTypes:
	$(RUN) -A list_issueTypes -C "40"

list_categories:
	$(RUN) -A list_categories -C "41"

list_users:
	$(RUN) -A list_users

list_versions:
	$(RUN) -A list_versions -C "80"

