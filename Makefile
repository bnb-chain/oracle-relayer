BUILD_TAGS = netgo
PACKAGES=$(shell go list ./...)

build:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/relayer.exe main.go
else
	go build $(BUILD_FLAGS) -o build/relayer main.go
endif

build_docker: build
	docker build --tag oracle_relayer .

test:
	make set_with_deadlock
	make test_unit
	make cleanup_after_test_with_deadlock

test_unit:
	@echo "--> go test "
	@go test -race $(PACKAGES)

# uses https://github.com/sasha-s/go-deadlock/ to detect potential deadlocks
set_with_deadlock:
	go get -u golang.org/x/tools/cmd/goimports
	go get github.com/sasha-s/go-deadlock

	find . -name "*.go" | xargs -n 1 sed -i.mutex_bak 's/sync.RWMutex/deadlock.RWMutex/'
	find . -name "*.go" | xargs -n 1 sed -i.mutex_bak 's/sync.Mutex/deadlock.Mutex/'
	go mod download
	find . -name "*.go" | xargs -n 1 goimports -w

# cleanes up after you ran test_with_deadlock
cleanup_after_test_with_deadlock:
	find . -name "*.go" | xargs -n 1 sed -i.mutex_bak 's/deadlock.RWMutex/sync.RWMutex/'
	find . -name "*.go" | xargs -n 1 sed -i.mutex_bak 's/deadlock.Mutex/sync.Mutex/'
	find . -name "*.go" | xargs -n 1 goimports -w
	find . -name "*.go.mutex_bak" | xargs rm

update_mock_executor:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	$(shell mockgen -source=executor/executor.go -package mock > executor/mock/mock_executor.go)

.PHONY: build build_docker test test_unit set_with_deadlock cleanup_after_test_with_deadlock update_mock_executor
