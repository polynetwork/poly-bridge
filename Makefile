# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build -tags $(env)
GOTEST=$(GOCMD) test

# Set and confirm environment `BRIDGE`, it should be one of devnet/testnet/mainnet
env=$(BRIDGE)
BaseDir=build/$(env)

.PHONY: all test clean

bridge_http:
	@mkdir -p $(BaseDir)/bridge_http
	@$(GOBUILD) -o $(BaseDir)/bridge_http/http_service main.go

bridge_server:
	@mkdir -p $(BaseDir)/bridge_server
	@$(GOBUILD) -o $(BaseDir)/bridge_server/bridge_server cmd/main.go

nft_bridge_http:
	@mkdir -p $(BaseDir)/nft_bridge_http/conf
	@mkdir -p $(BaseDir)/nft_bridge_http/logs
	@$(GOBUILD) -o $(BaseDir)/nft_bridge_http/http_service nft_http/*.go

bridge_tool:
	@mkdir -p $(BaseDir)/bridge_tools
	@$(GOBUILD) -o $(BaseDir)/bridge_tools/bridge_tool bridge_tools/*.go

deploy_tool:
	@mkdir -p $(BaseDir)/deploy_tool
	@$(GOBUILD) -o $(BaseDir)/deploy_tool/deploy_tool chain_tool/*.go

all:
	make nft_bridge_http bridge_tool deploy_tool

clean:
	@rm -rf $(BaseDir)/nft_bridge_http/http_service
	@rm -rf $(BaseDir)/bridge_tools/bridge_tool
	@rm -rf $(BaseDir)/deploy_tool/deploy_tool