pkgname="where-to"
fpath="/usr/share/zsh/site-functions"

dependencies:
	go mod tidy

install:
	go build -o find-where-to-go ./find-where-to-go-cli/main.go
	cp ./find-where-to-go /usr/bin/

# run to get zsh completions
zsh-completions:
	sudo mkdir -p "${fpath}"
	./"${pkgname}" --completion zsh | sudo tee "${fpath}/_${pkgname}" > /dev/null


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
