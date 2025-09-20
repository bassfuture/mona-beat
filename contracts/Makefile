# Makefile for CaptureNFT project

# Load environment variables
include .env
export

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build      - Compile contracts"
	@echo "  test       - Run tests"
	@echo "  test-v     - Run tests with verbose output"
	@echo "  clean      - Clean build artifacts"
	@echo "  deploy-local - Deploy to local network"
	@echo "  deploy-sepolia - Deploy to Sepolia testnet"
	@echo "  verify     - Verify contract on Etherscan"
	@echo "  format     - Format code"
	@echo "  lint       - Run linter"

# Build contracts
.PHONY: build
build:
	forge build

# Run tests
.PHONY: test
test:
	forge test

# Run tests with verbose output
.PHONY: test-v
test-v:
	forge test -vvv

# Clean build artifacts
.PHONY: clean
clean:
	forge clean

# Deploy to local network (anvil)
.PHONY: deploy-local
deploy-local:
	forge script script/Deploy.s.sol:DeployScript --rpc-url http://localhost:8545 --private-key $(PRIVATE_KEY) --broadcast

# Deploy to Sepolia testnet
.PHONY: deploy-sepolia
deploy-sepolia:
	forge script script/Deploy.s.sol:DeployScript --rpc-url $(SEPOLIA_RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(ETHERSCAN_API_KEY)

# Verify contract on Etherscan
.PHONY: verify
verify:
	forge verify-contract $(CAPTURE_NFT_ADDRESS) src/CaptureNFT.sol:CaptureNFT --etherscan-api-key $(ETHERSCAN_API_KEY)

# Format code
.PHONY: format
format:
	forge fmt

# Run linter
.PHONY: lint
lint:
	forge fmt --check