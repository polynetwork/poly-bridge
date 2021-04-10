# Poly NFT Bridge

PolyNFTBridge的API。

## API

* [GET /](#get)
* [POST assetshow](#post-assetshow)
* [POST assets](#post-assets)
* [POST asset](#post-asset)
* [POST items](#post-items)
* [POST getfee](#post-getfee)
* [POST transactionsofaddress](#post-transactionsofaddress)
* [POST transactionofhash](#post-transactionofhash)

## Test Node
[testnet](https://bridge.poly.network/nft/testnet/v1/)
[mainnet](https://bridge.poly.network/nft/v1/)

## 交易状态码

状态码|描述
:--:|:--:
0|finished
1|pendding
2|source done
3|source confirmed
4|poly confirmed

## 跨链交易手续费

### 手续费计算
代理收取的手续费 = 目标链交易的手续费 * 120% （120%可配置）

以BSC上的BNB跨链到以太手续费来计算：
fee = (eth.gas_limit * eth.gas_price) * (eth的USDT价格) / (BNB的USDT价格)

收取的手续费的资产为BNB

### 手续费检查

hasPay = 收取的手续费 > 目标链交易的手续费 * 20% （20%可配置）

以BSC上的BNB跨链到以太的过程来检查手续费：

hasPay = 收取的BNB * (BNB的USDT价格) > (eth.gas_limit * eth.gas_price) * (eth的USDT价格) * 20%

## API Info

### GET

Request 
```
http://localhost:8080/nft/v1
```

Example Request
```
curl --location --request GET 'http://localhost:8080/nft/v1'
```

Example Response
```
{
    "Version": "v1",
    "URL": "http://192.**.**.36:8081/nft"
}
```

### POST assetshow

Request 
```
http://localhost:8080/nft/v1/assetshow
```

BODY raw
```
{
    "ChainId": 2,
    "Size": 10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/assetshow' \
--data-raw '{
    "ChainId": 2,
    "Size": 10
}'
```

Example Response
```
{
    "TotalCount": 1,
    "Assets": [
        {
            "Asset": {
                "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
                "ChainId": 2,
                "Name": "cat1",
                "BaseUri": "http://106.75.250.160:10060/minio/"
            },
            "Items": [
                {
                    "TokenId": 1,
                    "Url": "http://106.75.250.160:10060/minio/1"
                },
                {
                    "TokenId": 2,
                    "Url": "http://106.75.250.160:10060/minio/2"
                }
            ]
        }
    ]
}
```

### POST assets

Request 
```
http://localhost:8080/nft/v1/assets
```

BODY raw
```
{
    "ChainId": 2
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/assets' \
--data-raw '{
    "ChainId": 2
}'
```

Example Response
```
{
    "TotalCount": 1,
    "Assets": [
        {
            "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
            "ChainId": 6,
            "Name": "cat1",
            "BaseUri": "http://106.75.250.160:10060/minio/",
            "DstAssets": [
                {
                    "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
                    "ChainId": 2,
                    "Name": "cat1",
                    "BaseUri": "http://106.75.250.160:10060/minio/"
                }
            ]
        }
    ]
}
```

### POST asset
   
Request 
```
http://localhost:8080/nft/v1/asset/
```

BODY raw
```
{
   "ChainId": 2,
   "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/asset/' \
--data-raw '{
    "ChainId": 2,
   "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001"
}'
```

Example Response
```
{
    "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
    "ChainId": 2,
    "Name": "cat1",
    "BaseUri": "http://106.75.250.160:10060/minio/",
    "DstAssets": [
        {
            "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
            "ChainId": 6,
            "Name": "cat1",
            "BaseUri": "http://106.75.250.160:10060/minio/"
        }
    ]
}
```

### POST items

Request 
```
http://localhost:8080/nft/v1/items/
```

BODY raw
```
{
    "ChainId": 2,
    "Asset": "03d84da9432f7cb5364a8b99286f97c59f738001",
    "Address": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
    "TokenId": "",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/items/' \
--data-raw '{
    "ChainId": 2,
    "Asset": "03d84da9432f7cb5364a8b99286f97c59f738001",
    "Address": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
    "TokenId": "",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 2,
    "Items": [
        {
            "TokenId": 1,
            "Url": "http://106.75.250.160:10060/minio/1"
        },
        {
            "TokenId": 2,
            "Url": "http://106.75.250.160:10060/minio/2"
        }
    ]
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/items/' \
--data-raw '{
    "ChainId": 2,
    "Asset": "03d84da9432f7cb5364a8b99286f97c59f738001",
    "Address": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
    "TokenId": "1",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 2,
    "Items": [
        {
            "TokenId": 1,
            "Url": "http://106.75.250.160:10060/minio/1"
        }
    ]
}
```

### POST getfee

Request 
```
http://localhost:8080/nft/v1/getfee/
```

BODY raw
```
{
    "SrcChainId": 2,
    "Hash": "0000000000000000000000000000000000000000",
    "DstChainId":6
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/getfee/' \
--data-raw '{
    "SrcChainId": 2,
    "Hash": "0000000000000000000000000000000000000000",
    "DstChainId":6
}'
```

Example Response
```
{
    "SrcChainId": 2,
    "Hash": "0000000000000000000000000000000000000000",
    "DstChainId": 6,
    "UsdtAmount": "0.5848838989488",
    "TokenAmount": "0.00027696",
    "TokenAmountWithPrecision": "276953315960434.16"
}
```

### POST transactionsofaddress

Request 
```
http://localhost:8080/nft/v1/transactionsofaddress/
```

BODY raw
```
{
    "Addresses":["5fb03eb21303d39967a1a119b32dd744a0fa8986"],
    "PageNo":0,
    "PageSize":10,
    "State": -1
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/transactionsofaddress/' \
--data-raw '{
    "Addresses":["5fb03eb21303d39967a1a119b32dd744a0fa8986"],
    "PageNo":0,
    "PageSize":10,
    "State": -1
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 2,
    "Transactions": [
        {
            "Hash": "6fdd31b04e4593aa9a5cd7fe5d6b80dd00b0e0106bbff3264427053cb0e91635",
            "User": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
            "SrcChainId": 6,
            "BlockHeight": 10445,
            "Time": 1617775450,
            "DstChainId": 2,
            "DstUser": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
            "TokenId": "2",
            "ServerId": 2,
            "FeeToken": {
                "Hash": "0000000000000000000000000000000000000000",
                "ChainId": 6,
                "Name": "BNB",
                "Property": 1,
                "TokenBasicName": "BNB",
                "TokenBasic": {
                    "Name": "BNB",
                    "Precision": 18,
                    "Price": "406.1693743",
                    "Ind": 1,
                    "Time": 1617764506,
                    "Property": 0,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            },
            "FeeAmount": "0.01",
            "State": 2,
            "SrcAsset": {
                "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
                "ChainId": 6,
                "Name": "cat1",
                "BaseUri": "http://106.75.250.160:10060/minio/"
            },
            "DstAsset": {
                "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
                "ChainId": 2,
                "Name": "cat1",
                "BaseUri": "http://106.75.250.160:10060/minio/"
            },
            "TransactionState": [
                {
                    "Hash": "6fdd31b04e4593aa9a5cd7fe5d6b80dd00b0e0106bbff3264427053cb0e91635",
                    "ChainId": 6,
                    "Blocks": 1,
                    "NeedBlocks": 1,
                    "Time": 1617775450
                },
                {
                    "Hash": "59adef4817eaa9a2486114812b7b6b9fe86fb0e51d6619528a4778fb16779d0f",
                    "ChainId": 0,
                    "Blocks": 1,
                    "NeedBlocks": 1,
                    "Time": 1617775521
                },
                {
                    "Hash": "f981d72309f365e7c46e1e4b07fb4403946893dd380daddbc1f4fa390307253d",
                    "ChainId": 2,
                    "Blocks": 1,
                    "NeedBlocks": 1,
                    "Time": 1617775530
                }
            ]
        },
        {
            "Hash": "fd6925f73492775f4709d68d00de5dc5ce4894c56eab6a04eec2b16e81359835",
            "User": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
            "SrcChainId": 2,
            "BlockHeight": 4831,
            "Time": 1617775170,
            "DstChainId": 6,
            "DstUser": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
            "TokenId": "2",
            "ServerId": 2,
            "FeeToken": {
                "Hash": "0000000000000000000000000000000000000000",
                "ChainId": 2,
                "Name": "ETH",
                "Property": 1,
                "TokenBasicName": "ETH",
                "TokenBasic": {
                    "Name": "ETH",
                    "Precision": 18,
                    "Price": "2111.85014",
                    "Ind": 1,
                    "Time": 1617764506,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            },
            "FeeAmount": "0.01",
            "State": 2,
            "SrcAsset": {
                "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
                "ChainId": 2,
                "Name": "cat1",
                "BaseUri": "http://106.75.250.160:10060/minio/"
            },
            "DstAsset": {
                "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
                "ChainId": 6,
                "Name": "cat1",
                "BaseUri": "http://106.75.250.160:10060/minio/"
            },
            "TransactionState": [
                {
                    "Hash": "fd6925f73492775f4709d68d00de5dc5ce4894c56eab6a04eec2b16e81359835",
                    "ChainId": 2,
                    "Blocks": 1,
                    "NeedBlocks": 1,
                    "Time": 1617775170
                },
                {
                    "Hash": "772619300f3577a6545107e0b40b47bf5e73bf173640480fec6dc73bb8ee4835",
                    "ChainId": 0,
                    "Blocks": 1,
                    "NeedBlocks": 1,
                    "Time": 1617775399
                },
                {
                    "Hash": "d0bb95464df21f83de0bae4f214f950d9170012821faa69d6e1726118a36f7e8",
                    "ChainId": 6,
                    "Blocks": 1,
                    "NeedBlocks": 1,
                    "Time": 1617775420
                }
            ]
        }
    ]
}
```

### POST transactionofhash

Request 
```
http://localhost:8080/nft/v1/transactionofhash/
```

BODY raw
```
{
    "Hash":"6fdd31b04e4593aa9a5cd7fe5d6b80dd00b0e0106bbff3264427053cb0e91635"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/nft/v1/transactionofhash/' \
--data-raw '{
    "Hash":"6fdd31b04e4593aa9a5cd7fe5d6b80dd00b0e0106bbff3264427053cb0e91635"
}'
```

Example Response
```
{
    "Hash": "6fdd31b04e4593aa9a5cd7fe5d6b80dd00b0e0106bbff3264427053cb0e91635",
    "User": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
    "SrcChainId": 6,
    "BlockHeight": 10445,
    "Time": 1617775450,
    "DstChainId": 2,
    "DstUser": "5fb03eb21303d39967a1a119b32dd744a0fa8986",
    "TokenId": "2",
    "ServerId": 2,
    "FeeToken": {
        "Hash": "0000000000000000000000000000000000000000",
        "ChainId": 6,
        "Name": "BNB",
        "Property": 1,
        "TokenBasicName": "BNB",
        "TokenBasic": {
            "Name": "BNB",
            "Precision": 18,
            "Price": "406.1693743",
            "Ind": 1,
            "Time": 1617764506,
            "Property": 0,
            "PriceMarkets": null,
            "Tokens": null
        },
        "TokenMaps": null
    },
    "FeeAmount": "0.01",
    "State": 2,
    "SrcAsset": {
        "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
        "ChainId": 6,
        "Name": "cat1",
        "BaseUri": "http://106.75.250.160:10060/minio/"
    },
    "DstAsset": {
        "Hash": "03d84da9432f7cb5364a8b99286f97c59f738001",
        "ChainId": 2,
        "Name": "cat1",
        "BaseUri": "http://106.75.250.160:10060/minio/"
    },
    "TransactionState": [
        {
            "Hash": "6fdd31b04e4593aa9a5cd7fe5d6b80dd00b0e0106bbff3264427053cb0e91635",
            "ChainId": 6,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1617775450
        },
        {
            "Hash": "59adef4817eaa9a2486114812b7b6b9fe86fb0e51d6619528a4778fb16779d0f",
            "ChainId": 0,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1617775521
        },
        {
            "Hash": "f981d72309f365e7c46e1e4b07fb4403946893dd380daddbc1f4fa390307253d",
            "ChainId": 2,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1617775530
        }
    ]
}
```