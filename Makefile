.DEFAULT_GOAL := start

start: build execute

build:
	go build

execute:
	./star-wars-go

commit: git-add git-commit git-push

git-add:
	git add --all

git-commit:
	git commit -m "${BRANCH} ${MESSAGE}"

git-push:
	git push origin ${BRANCH}