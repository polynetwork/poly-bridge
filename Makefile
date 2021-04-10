# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build -tags $(env)
GOTEST=$(GOCMD) test

# Set and confirm environment `BRIDGE`, it should be one of devnet/testnet/mainnet
env=$(BRIDGE)
BaseDir=build/$(env)

.PHONY: all test clean

nft_bridge_http:
	@mkdir -p $(BaseDir)/nft_bridge_http/conf
	@mkdir -p $(BaseDir)/nft_bridge_http/logs
	@$(GOBUILD) -o $(BaseDir)/nft_bridge_http/http_service nft_http/*.go

eth_listen:
	@mkdir -p $(BaseDir)/eth_listen/logs
	@mkdir -p $(BaseDir)/bsc_listen/logs
	@mkdir -p $(BaseDir)/heco_listen/logs
	@$(GOBUILD) -o $(BaseDir)/eth_listen/listener crosschainlisten/ethereumlisten/cmd/main.go
	@cp $(BaseDir)/eth_listen/listener $(BaseDir)/bsc_listen/listener
	@cp $(BaseDir)/eth_listen/listener $(BaseDir)/heco_listen/listener

poly_listen:
	@mkdir -p $(BaseDir)/poly_listen/logs
	@$(GOBUILD) -o $(BaseDir)/poly_listen/listener crosschainlisten/polylisten/cmd/main.go

nft_asset_tool:
	@mkdir -p $(BaseDir)/nft_asset_tool/logs
	@$(GOBUILD) -o $(BaseDir)/nft_asset_tool/asset_tool nft_asset_tool/*.go

nft_deploy_tool:
	@mkdir -p $(BaseDir)/nft_deploy_tool/logs
	@$(GOBUILD) -o $(BaseDir)/nft_deploy_tool/deploy_tool nft_deploy_tool/*.go

all:
	make nft_http eth_listen poly_listen nft_asset_tool nft_deploy_tool

clean:
	rm -rf $(BaseDir)/nft_bridge_http/http_service
	rm -rf $(BaseDir)/eth_listen/listener
	rm -rf $(BaseDir)/bsc_listen/listener
	rm -rf $(BaseDir)/heco_listen/listener
	rm -rf $(BaseDir)/poly_listen/listener