fpath="/usr/share/zsh/site-functions"
exedir="/usr/local/bin/"

build:
	go mod tidy
	go build -o find-where-to-go ./find-where-to-go-cli/main.go
	go build -o where-to ./where-to-cli/main.go

install:
	mv ./find-where-to-go "${exedir}"
	mv ./where-to "${exedir}"

zsh-completions:
	sudo mkdir -p "${fpath}"
	where-to --completion zsh | sudo tee "${fpath}/_where-to" > /dev/null
	find-where-to-go --completion zsh | sudo tee "${fpath}/_find-where-to-go" > /dev/null
	sudo cp ./.completions/zsh/_to /usr/share/zsh/site-functions

bash-completions:
	sudo mkdir -p "${fpath}"
	where-to --completion bash | sudo tee "${fpath}/_where-to" > /dev/null
	find-where-to-go --completion bash | sudo tee "${fpath}/_find-where-to-go" > /dev/null
	sudo cp ./.completions/zsh/_to /usr/share/zsh/site-functions

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
