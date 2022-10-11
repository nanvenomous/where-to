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
	mkdir -p "${fpath}"
	where-to --completion zsh | sudo tee "${fpath}/_where-to" > /dev/null
	find-where-to-go --completion zsh | sudo tee "${fpath}/_find-where-to-go" > /dev/null
	cp ./.completions/zsh/_to "${fpath}"

bash-completions:
	mkdir -p "${fpath}"
	where-to --completion bash | sudo tee "${fpath}/_where-to" > /dev/null
	find-where-to-go --completion bash | sudo tee "${fpath}/_find-where-to-go" > /dev/null
	cp ./.completions/zsh/_to "${fpath}"

release:
	mkdir -p ./bin/macos
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/find-where-to-go ./find-where-to-go-cli/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/where-to ./where-to-cli/main.go
	mkdir -p ./bin/macosArm
	GOOS=darwin GOARCH=arm64 go build -o bin/macosArm/find-where-to-go ./find-where-to-go-cli/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/macosArm/where-to ./where-to-cli/main.go
	mkdir -p ./bin/linux
	GOOS=linux GOARCH=amd64 go build -o bin/linux/find-where-to-go ./find-where-to-go-cli/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/where-to ./where-to-cli/main.go
	tar -cJf ./bin/macos.tar.xz ./bin/macos
	rm -rf ./bin/macos
	tar -cJf ./bin/macosArm.tar.xz ./bin/macosArm
	rm -rf ./bin/macosArm
	tar -cJf ./bin/linux.tar.xz ./bin/linux
	rm -rf ./bin/linux

clean:
	rm -rf ./bin
	rm -rf ./find-where-to-go
	rm -rf ./where-to
	rm -rf "${exedir}/find-where-to-go"
	rm -rf "${exedir}/where-to"
	rm -rf "${fpath}/_to"
	rm -rf "${fpath}/_where-to"
	rm -rf "${fpath}/_find-where-to-go"

test:
	go test -count=1 ./...
