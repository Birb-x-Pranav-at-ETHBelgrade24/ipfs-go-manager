build:
	@go build -o bin/backend cmd/server/main.go

test:
	@go test -v ./...

run: build
	@./bin/backend

STORAGE_SOURCE_FILE=./../hackathon-mvp/packages/hardhat/artifacts/contracts/StorageContract.sol/StorageContract.json
STORAGE_TARGET_FILE=./../hackathon-mvp/packages/hardhat/artifacts/contracts/StorageContract.sol/StorageContract.abigen.json

# Extract the abi array and save it as the root element in a new file
abiconv:
	jq '.abi' $(STORAGE_SOURCE_FILE) > $(STORAGE_TARGET_FILE)

abigen: abiconv
	@abigen --abi $(STORAGE_TARGET_FILE) --pkg contract --out ./internal/web3/contract/StorageContract.go
