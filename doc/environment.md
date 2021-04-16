## NFT cross chain environment prepare

#### deploy proxy contracts

```shell script
./deploy_tool --chain=2 deployNFTLockProxy
./deploy_tool --chain=2 proxySetCCMP

./deploy_tool --chain=6 deployNFTLockProxy
./deploy_tool --chain=6 proxySetCCMP

./deploy_tool --chain=7 deployNFTLockProxy
./deploy_tool --chain=7 proxySetCCMP

./deploy_tool --chain=2 bindProxy --dstChain=6
./deploy_tool --chain=6 bindProxy --dstChain=2

./deploy_tool --chain=2 bindProxy --dstChain=7
./deploy_tool --chain=7 bindProxy --dstChain=2

./deploy_tool --chain=6 bindProxy --dstChain=7
./deploy_tool --chain=7 bindProxy --dstChain=6
```

#### deploy wrap contract
```shell script
./deploy_tool --chain=2 deployNFTWrapper
./deploy_tool --chain=2 setWrapLockProxy
./deploy_tool --chain=2 setFeeCollector

./deploy_tool --chain=6 deployNFTWrapper
./deploy_tool --chain=6 setWrapLockProxy
./deploy_tool --chain=6 setFeeCollector

./deploy_tool --chain=7 deployNFTWrapper
./deploy_tool --chain=7 setWrapLockProxy
./deploy_tool --chain=7 setFeeCollector
```

#### bind NFT asset contracts
e.g:

nft contract on etherem: A
nft contract on bsc    : B
nft contract on heco   : C

```shell script
./deploy_tool --chain=2 bindNFT --asset=A --dstChain=6 --dstAsset=B
./deploy_tool --chain=6 bindNFT --asset=B --dstChain=2 --dstAsset=A

./deploy_tool --chain=2 bindNFT --asset=A --dstChain=7 --dstAsset=C
./deploy_tool --chain=7 bindNFT --asset=C --dstChain=2 --dstAsset=A

./deploy_tool --chain=6 bindNFT --asset=B --dstChain=7 --dstAsset=C
./deploy_tool --chain=7 bindNFT --asset=C --dstChain=6 --dstAsset=B
```