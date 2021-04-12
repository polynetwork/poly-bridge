## Deploy tool


### 准备
给几条链的admin账户转native token，这里以console模式为例:
```bash
eth.accounts

personal.unlockAccount("0xa**b");

eth.getBalance("0xa**b");

eth.sendTransaction({from: "0xc**d4",to: "0x7a**6a", value: "7**0"});
```
这里需要注意，`personal.unlockAccount`需要在节点启动的时候配置`--allow-insecure-unlock`

#### 部署erc20token

1.部署erc20合约:
```shell script
./deploy_tool --chain=2 deployFee
./deploy_tool --chain=79 deployFee
./deploy_tool --chain=7 deployFee
```
这本合约，可以作为wrapper的feeToken(参数`--feeToken=true`, 因为默认为true，所以忽略), 同时继承了mintable属性，我们可以在dev环境下准备相应的token做测试.

2.mint erc20 token:
```shell script
./deploy_tool --chain=2 mintFee --to=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --amount=100000000000000000000000000
./deploy_tool --chain=79 mintFee --to=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --amount=100000000000000000000000000
```
这个fee后面可以用来支付wrapper手续费

3.transfer erc20 token:

本地测试数据
```dtd
user1 0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 
user2 0x9B5a65b6946084AFB2d7c2851318665daA0519E9
user3 0xa252dCBF98D02218b4E5B7B00d8FE7646592394E
user4 0xB9933ff0CB5C5B42b12972C9826703E10BFDd863
```

给用户一定量的native token
```shell script
./deploy_tool --chain=2 nativeBalance --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C
./deploy_tool --chain=79 nativeBalance --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C

./deploy_tool --chain=2 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 --amount=1000000000000000000
./deploy_tool --chain=2 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x9B5a65b6946084AFB2d7c2851318665daA0519E9 --amount=1000000000000000000

./deploy_tool --chain=79 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 --amount=1000000000000000000
./deploy_tool --chain=79 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x9B5a65b6946084AFB2d7c2851318665daA0519E9 --amount=1000000000000000000

./deploy_tool --chain=79 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0xc9d93747EAbc4d738fAE96bcF64d9EC04ffD8C95 --amount=2000000000000000000
```

为relayer发送账户准备native token
```shell script
./deploy_tool --chain=2 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x69611B922c985fD793AFA56CE8Cfe7d8aFffeFDd --amount=100000000000000000000
./deploy_tool --chain=79 transferNative --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0xc4A519453135aE569c582154A5E632668E6DADc4 --amount=100000000000000000000
```

该部分测试必须在部署完feeToken, 并mint一部分给管理员后操作.每个用户10000个feeToken

```shell script
./deploy_tool --chain=2 transferFee --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 --amount=10000000000000000000000
./deploy_tool --chain=2 transferFee --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x9B5a65b6946084AFB2d7c2851318665daA0519E9 --amount=10000000000000000000000

./deploy_tool --chain=79 transferFee --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 --amount=10000000000000000000000
./deploy_tool --chain=79 transferFee --from=0x31c0dd87B33Dcd66f9a255Cf4CF39287F8AE593C --to=0x9B5a65b6946084AFB2d7c2851318665daA0519E9 --amount=10000000000000000000000

./deploy_tool --chain=2 feeBalance --from=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986
./deploy_tool --chain=2 feeBalance --from=0x9B5a65b6946084AFB2d7c2851318665daA0519E9

./deploy_tool --chain=79 feeBalance --from=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986
./deploy_tool --chain=79 feeBalance --from=0x9B5a65b6946084AFB2d7c2851318665daA0519E9

./deploy_tool --chain=2 approve --from=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 --amount=1000000000000000000000000
./deploy_tool --chain=2 allowance --from=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986

./deploy_tool --chain=79 approve --from=0x9B5a65b6946084AFB2d7c2851318665daA0519E9 --amount=1000000000000000000000000
./deploy_tool --chain=79 allowance --from=0x9B5a65b6946084AFB2d7c2851318665daA0519E9

```

#### 链基础合约

1.以太
```shell script
./deploy_tool --chain=2 deployECCD
./deploy_tool --chain=2 deployECCM
./deploy_tool --chain=2 deployCCMP

./deploy_tool --chain=2 transferECCDOwnership
./deploy_tool --chain=2 transferECCMOwnership

./deploy_tool --chain=2 registerSideChain
./deploy_tool --chain=2 approveSideChain

./deploy_tool --chain=2 syncSideGenesis
./deploy_tool --chain=2 syncPolyGenesis
```

2.bsc
```shell script
./deploy_tool --chain=79 deployECCD
./deploy_tool --chain=79 deployECCM
./deploy_tool --chain=79 deployCCMP

./deploy_tool --chain=79 transferECCDOwnership
./deploy_tool --chain=79 transferECCMOwnership

./deploy_tool --chain=79 registerSideChain
./deploy_tool --chain=79 approveSideChain

./deploy_tool --chain=79 syncSideGenesis
./deploy_tool --chain=79 syncPolyGenesis
```

3.heco
```shell script
./deploy_tool --chain=7 deployECCD
./deploy_tool --chain=7 deployECCM
./deploy_tool --chain=7 deployCCMP

./deploy_tool --chain=7 transferECCDOwnership
./deploy_tool --chain=7 transferECCMOwnership

./deploy_tool --chain=7 registerSideChain
./deploy_tool --chain=7 approveSideChain

./deploy_tool --chain=7 syncSideGenesis
./deploy_tool --chain=7 syncPolyGenesis
```

#### NFTLockProxy合约

```shell script
./deploy_tool --chain=2 deployNFTLockProxy
./deploy_tool --chain=2 proxySetCCMP

./deploy_tool --chain=79 deployNFTLockProxy
./deploy_tool --chain=79 proxySetCCMP

./deploy_tool --chain=7 deployNFTLockProxy
./deploy_tool --chain=7 proxySetCCMP

./deploy_tool --chain=2 bindProxy --dstChain=79
./deploy_tool --chain=79 bindProxy --dstChain=2

./deploy_tool --chain=2 bindProxy --dstChain=7
./deploy_tool --chain=7 bindProxy --dstChain=2

./deploy_tool --chain=79 bindProxy --dstChain=7
./deploy_tool --chain=7 bindProxy --dstChain=79

```

#### NFTWrap合约
```shell script
./deploy_tool --chain=2 deployNFTWrapper
./deploy_tool --chain=2 setWrapLockProxy
./deploy_tool --chain=2 setFeeCollector

./deploy_tool --chain=79 deployNFTWrapper
./deploy_tool --chain=79 setWrapLockProxy
./deploy_tool --chain=79 setFeeCollector

./deploy_tool --chain=7 deployNFTWrapper
./deploy_tool --chain=7 setWrapLockProxy
./deploy_tool --chain=7 setFeeCollector
```

#### NFT资产合约及绑定

这里以eth和bsc的跨链为例
```shell script

#todo: erc721合约_safeMint已修改.在mint的时候不要进入到onReceive方法，因为现在的lock proxy的onReceive方法中只接收来自proxy的行为 

./deploy_tool --chain=2 deployNFT --name=digitalCat1 --symbol=cat1
./deploy_tool --chain=79 deployNFT --name=digitalCat1 --symbol=cat1

./deploy_tool --chain=2 bindNFT --asset=0xa85c9FC8F2c9060d674E0CA97F703a0A30619305 --dstChain=79 --dstAsset=0x455B51D882571E244d03668f1a458ca74E70d196
./deploy_tool --chain=79 bindNFT --asset=0x455B51D882571E244d03668f1a458ca74E70d196 --dstChain=2 --dstAsset=0xa85c9FC8F2c9060d674E0CA97F703a0A30619305
```

#### NFT跨链

```shell script

# mint nft token
./deploy_tool --chain=2 mintNFT --asset=0xa85c9FC8F2c9060d674E0CA97F703a0A30619305 --to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 --tokenId=11

# wrapper lock nft
./deploy_tool --chain=2 lockNFT --dstChain=79 \
--asset=0xa85c9FC8F2c9060d674E0CA97F703a0A30619305 \
--from=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 \
--to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 \
--amount=1000000000000000 \
--tokenId=11 --lockId=11


./deploy_tool --chain=79 lockNFT --dstChain=2 \
--asset=0x455B51D882571E244d03668f1a458ca74E70d196 \
--from=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 \
--to=0x5Fb03EB21303D39967a1a119B32DD744a0fA8986 \
--amount=1000000000000000 \
--tokenId=11 --lockId=11

./deploy_tool --chain=2 owner --asset=0xa85c9FC8F2c9060d674E0CA97F703a0A30619305 --tokenId=11
./deploy_tool --chain=79 owner --asset=0x455B51D882571E244d03668f1a458ca74E70d196 --tokenId=11

```