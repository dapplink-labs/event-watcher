GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

TM_ABI_ARTIFACT := /Users/guoshijiang/theweb3/event-watcher/abis/TreasureManager.sol/TreasureManager.json


event-watcher:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/event-watcher

clean:
	rm event-watcher

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings:
	$(eval temp := $(shell mktemp))

	cat $(TM_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(TM_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/treasure_manager.go \
		--type TreasureManager \
		--bin $(temp)

		rm $(temp)

.PHONY: \
	event-watcher \
	bindings \
	clean \
	test \
	lint