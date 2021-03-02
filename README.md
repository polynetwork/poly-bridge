# Poly Bridge

PolyBridge的API。

## API

* [GET /](#get-/)
* [POST tokens](#post-tokens)
* [POST token](#post-token)
* [POST tokenbasics](#post-tokenbasics)
* [POST tokenmap](#post-tokenmap)
* [POST tokenmapreverse](#post-tokenmapreverse)
* [POST getfee](#post-getfee)
* [POST checkfee](#post-checkfee)
* [POST transactions](#post-transactions)
* [POST transactionsofaddress](#post-transactionsofaddress)
* [POST transactionofhash](#post-transactionofhash)
* [POST transactionsofstate](#post-transactionsofstate)

## Test Node
[testnet](https://bridge.poly.network/testnet/v1/)
[mainnet](https://bridge.poly.network/v1/)

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

### GET /

Request 
```
http://localhost:8080/v1
```

Example Request
```
curl --location --request GET 'http://localhost:8080/v1'
```

Example Response
```
{
    "Version": "v1",
    "URL": "http://localhost:8080/v1"
}
```

### POST tokens

Request 
```
http://localhost:8080/v1/tokens
```

BODY raw
```
{
    "ChainId": 2
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokens' \
--data-raw '{
    "ChainId": 2
}'
```

Example Response
```
{
    "TotalCount": 6,
    "Tokens": [
        {
            "Hash": "0000000000000000000000000000000000000000",
            "ChainId": 2,
            "Name": "Ethereum",
            "TokenBasicName": "Ethereum",
            "TokenBasic": {
                "Name": "Ethereum",
                "Precision": 18,
                "Price": "1227.921489",
                "Ind": 1,
                "Time": 1610668843,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "0000000000000000000000000000000000000000",
                    "SrcToken": null,
                    "DstTokenHash": "23535b6fd46b8f867ed010bab4c2bd8ef0d0c64f",
                    "DstToken": {
                        "Hash": "23535b6fd46b8f867ed010bab4c2bd8ef0d0c64f",
                        "ChainId": 4,
                        "Name": "pnWETH",
                        "TokenBasicName": "Ethereum",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "09c6a1b0b32a8b2c327532518c68f9b0c54255b8",
            "ChainId": 2,
            "Name": "BNB",
            "TokenBasicName": "BNB",
            "TokenBasic": {
                "Name": "BNB",
                "Precision": 18,
                "Price": "41.84074196",
                "Ind": 1,
                "Time": 1610668843,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "09c6a1b0b32a8b2c327532518c68f9b0c54255b8",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000000000000000000000000",
                    "DstToken": {
                        "Hash": "0000000000000000000000000000000000000000",
                        "ChainId": 79,
                        "Name": "BNB",
                        "TokenBasicName": "BNB",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "557563dc4ed3fd256eba55b9622f53331ab97c2f",
            "ChainId": 2,
            "Name": "WBTC",
            "TokenBasicName": "WBTC",
            "TokenBasic": {
                "Name": "WBTC",
                "Precision": 8,
                "Price": "39108.08931",
                "Ind": 1,
                "Time": 1610668843,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "557563dc4ed3fd256eba55b9622f53331ab97c2f",
                    "SrcToken": null,
                    "DstTokenHash": "a3ce15f11d4427b6bad5630036f368a98e923e95",
                    "DstToken": {
                        "Hash": "a3ce15f11d4427b6bad5630036f368a98e923e95",
                        "ChainId": 79,
                        "Name": "WBTC",
                        "TokenBasicName": "WBTC",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "7e269f2f33a97c64192e9889faeec72a6fcdb397",
            "ChainId": 2,
            "Name": "eNEO",
            "TokenBasicName": "Neo",
            "TokenBasic": {
                "Name": "Neo",
                "Precision": 8,
                "Price": "23.0547776",
                "Ind": 1,
                "Time": 1610668843,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "7e269f2f33a97c64192e9889faeec72a6fcdb397",
                    "SrcToken": null,
                    "DstTokenHash": "17da3881ab2d050fea414c80b3fa8324d756f60e",
                    "DstToken": {
                        "Hash": "17da3881ab2d050fea414c80b3fa8324d756f60e",
                        "ChainId": 4,
                        "Name": "nNeo",
                        "TokenBasicName": "Neo",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "ChainId": 2,
            "Name": "USDT",
            "TokenBasicName": "USDT",
            "TokenBasic": {
                "Name": "USDT",
                "Precision": 6,
                "Price": "0.99896425",
                "Ind": 1,
                "Time": 1610668843,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                    "SrcToken": null,
                    "DstTokenHash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
                    "DstToken": {
                        "Hash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
                        "ChainId": 4,
                        "Name": "pnUSDT",
                        "TokenBasicName": "USDT",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                },
                {
                    "SrcTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                    "SrcToken": null,
                    "DstTokenHash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
                    "DstToken": {
                        "Hash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
                        "ChainId": 79,
                        "Name": "USDT",
                        "TokenBasicName": "USDT",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "b60e03e6973b1d0b90a763f5b64c48ca7cb8c2d1",
            "ChainId": 2,
            "Name": "WING",
            "TokenBasicName": "WING",
            "TokenBasic": {
                "Name": "WING",
                "Precision": 9,
                "Price": "12.91266349",
                "Ind": 1,
                "Time": 1610668843,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "b60e03e6973b1d0b90a763f5b64c48ca7cb8c2d1",
                    "SrcToken": null,
                    "DstTokenHash": "0a7bf54d2684885d731dc63917a3178a2a1a8d4a",
                    "DstToken": {
                        "Hash": "0a7bf54d2684885d731dc63917a3178a2a1a8d4a",
                        "ChainId": 79,
                        "Name": "WING",
                        "TokenBasicName": "WING",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        }
    ]
}
```

### POST token
   
Request 
```
http://localhost:8080/v1/token/
```

BODY raw
```
{
   "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/token/' \
--data-raw '{
   "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb"
}'
```

Example Response
```
{
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
    "ChainId": 2,
    "Name": "USDT",
    "TokenBasicName": "USDT",
    "TokenBasic": {
        "Name": "USDT",
        "Precision": 6,
        "Price": "0.99896425",
        "Ind": 1,
        "Time": 1610668843,
        "PriceMarkets": null,
        "Tokens": null
    },
    "TokenMaps": [
        {
            "SrcTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "SrcToken": null,
            "DstTokenHash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
            "DstToken": {
                "Hash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
                "ChainId": 4,
                "Name": "pnUSDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        },
        {
            "SrcTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "SrcToken": null,
            "DstTokenHash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
            "DstToken": {
                "Hash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
                "ChainId": 79,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        }
    ]
}
```

### POST tokenmap

Request 
```
http://localhost:8080/v1/tokenmap/
```

BODY raw
```
{
    "ChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokenmap/' \
--data-raw '{
    "ChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb"
}'
```

Example Response
```
{
    "TotalCount": 2,
    "TokenMaps": [
        {
            "SrcTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "SrcToken": {
                "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
            "DstToken": {
                "Hash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
                "ChainId": 4,
                "Name": "pnUSDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        },
        {
            "SrcTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "SrcToken": {
                "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
            "DstToken": {
                "Hash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
                "ChainId": 79,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        }
    ]
}
```

### POST tokenmapreverse

Request 
```
http://localhost:8080/v1/tokenmapreverse/
```

BODY raw
```
{
    "ChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokenmapreverse/' \
--data-raw '{
    "ChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb"
}'
```

Example Response
```
{
    "TotalCount": 2,
    "TokenMaps": [
        {
            "SrcTokenHash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
            "SrcToken": {
                "Hash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
                "ChainId": 4,
                "Name": "pnUSDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "DstToken": {
                "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        },
        {
            "SrcTokenHash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
            "SrcToken": {
                "Hash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
                "ChainId": 79,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
            "DstToken": {
                "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        }
    ]
}
```

### POST tokenbasics

Request 
```
http://localhost:8080/v1/tokenbasics/
```

BODY raw
```
{
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokenbasics/' \
--data-raw '{
}'
```

Example Response
```
{
    "TotalCount": 6,
    "TokenBasics": [
        {
            "Name": "BNB",
            "Precision": 18,
            "Price": "41.84074196",
            "Ind": 1,
            "Time": 1610668843,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000000000000000000000000",
                    "ChainId": 79,
                    "Name": "BNB",
                    "TokenBasicName": "BNB",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "09c6a1b0b32a8b2c327532518c68f9b0c54255b8",
                    "ChainId": 2,
                    "Name": "BNB",
                    "TokenBasicName": "BNB",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "Ethereum",
            "Precision": 18,
            "Price": "1227.921489",
            "Ind": 1,
            "Time": 1610668843,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000000000000000000000000",
                    "ChainId": 2,
                    "Name": "Ethereum",
                    "TokenBasicName": "Ethereum",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "23535b6fd46b8f867ed010bab4c2bd8ef0d0c64f",
                    "ChainId": 4,
                    "Name": "pnWETH",
                    "TokenBasicName": "Ethereum",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "Neo",
            "Precision": 8,
            "Price": "23.0547776",
            "Ind": 1,
            "Time": 1610668843,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "17da3881ab2d050fea414c80b3fa8324d756f60e",
                    "ChainId": 4,
                    "Name": "nNeo",
                    "TokenBasicName": "Neo",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "7e269f2f33a97c64192e9889faeec72a6fcdb397",
                    "ChainId": 2,
                    "Name": "eNEO",
                    "TokenBasicName": "Neo",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "USDT",
            "Precision": 6,
            "Price": "0.99896425",
            "Ind": 1,
            "Time": 1610668843,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "23f5075740c2c99c569ffd0768c383a92d1a4ad7",
                    "ChainId": 79,
                    "Name": "USDT",
                    "TokenBasicName": "USDT",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
                    "ChainId": 2,
                    "Name": "USDT",
                    "TokenBasicName": "USDT",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "b8f78d43ea9fe006c85a26b9aff67bcf69dd4fe1",
                    "ChainId": 4,
                    "Name": "pnUSDT",
                    "TokenBasicName": "USDT",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "WBTC",
            "Precision": 8,
            "Price": "39108.08931",
            "Ind": 1,
            "Time": 1610668843,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "557563dc4ed3fd256eba55b9622f53331ab97c2f",
                    "ChainId": 2,
                    "Name": "WBTC",
                    "TokenBasicName": "WBTC",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "a3ce15f11d4427b6bad5630036f368a98e923e95",
                    "ChainId": 79,
                    "Name": "WBTC",
                    "TokenBasicName": "WBTC",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "WING",
            "Precision": 9,
            "Price": "12.91266349",
            "Ind": 1,
            "Time": 1610668843,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0a7bf54d2684885d731dc63917a3178a2a1a8d4a",
                    "ChainId": 79,
                    "Name": "WING",
                    "TokenBasicName": "WING",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "b60e03e6973b1d0b90a763f5b64c48ca7cb8c2d1",
                    "ChainId": 2,
                    "Name": "WING",
                    "TokenBasicName": "WING",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        }
    ]
}
```

### POST getfee

Request 
```
http://localhost:8080/v1/getfee/
```

BODY raw
```
{
    "SrcChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
    "DstChainId":79
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/getfee/' \
--data-raw '{
    "SrcChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
    "DstChainId":79
}'
```

Example Response
```
{
    "SrcChainId": 2,
    "Hash": "ad3f96ae966ad60347f31845b7e4b333104c52fb",
    "DstChainId": 79,
    "UsdtAmount": "0.1205013368",
    "TokenAmount": "0.1206262755",
    "TokenAmountWithPrecision": "120626.2755"
}
```

### POST checkfee

Request 
```
http://localhost:8080/v1/checkfee/
```

BODY raw
```
{
    "Hashs": ["000000000000000000000000000000000000000000000000000000000000175c"]
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/checkfee/' \
--data-raw '{
    "Hashs": ["000000000000000000000000000000000000000000000000000000000000175c"]
}'
```

Example Response
```
{
    "TotalCount": 1,
    "CheckFees": [
        {
            "Hash": "000000000000000000000000000000000000000000000000000000000000175c",
            "PayState": 1,
            "Amount": "12.27921489",
            "MinProxyFee": "0"
        }
    ]
}
```

### POST transactions

Request 
```
http://localhost:8080/v1/transactions/
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactions/' \
--data-raw '{
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
    "TotalCount": 1,
    "Transactions": [
        {
            "Hash": "85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002",
            "User": "ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f",
            "SrcChainId": 2,
            "BlockHeight": 9469807,
            "Time": 1610695305,
            "DstChainId": 79,
            "FeeTokenHash": "0000000000000000000000000000000000000000",
            "FeeAmount": 10000000000000000,
            "State": 0
        }
    ]
}
```

### POST transactionsofaddress

Request 
```
http://localhost:8080/v1/transactionsofaddress/
```

BODY raw
```
{
    "Addresses":["ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"],
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofaddress/' \
--data-raw '{
    "Addresses":["ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"],
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
    "TotalCount": 1,
    "Transactions": [
        {
            "Hash": "85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002",
            "User": "ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f",
            "SrcChainId": 2,
            "BlockHeight": 9469807,
            "Time": 1610695305,
            "DstChainId": 79,
            "FeeAmount": "10000000000000000",
            "TransferAmount": "90000000000000000",
            "DstUser": "6e43f9988f2771f1a2b140cb3faad424767d39fc",
            "State": 0,
            "Token": {
                "Hash": "0000000000000000000000000000000000000000",
                "ChainId": 2,
                "Name": "Ethereum",
                "TokenBasicName": "Ethereum",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "TransactionState": [
                {
                    "Hash": "85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002",
                    "ChainId": 2,
                    "Blocks": 10,
                    "NeedBlocks": 10,
                    "Time": 1610695305
                },
                {
                    "Hash": "a58b5705c2117e390c7add98d55e762342c26508a9b787befa228e5c10a2b14f",
                    "ChainId": 0,
                    "Blocks": 10,
                    "NeedBlocks": 10,
                    "Time": 1610697074
                },
                {
                    "Hash": "5e201266b11f107dafa8e323b4be3b1c7f062bc1f1926ce36cf8832497342e37",
                    "ChainId": 79,
                    "Blocks": 10,
                    "NeedBlocks": 10,
                    "Time": 1610697089
                }
            ]
        }
    ]
}
```

### POST transactionofhash

Request 
```
http://localhost:8080/v1/transactionofhash/
```

BODY raw
```
{
    "Hash":"85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionofhash/' \
--data-raw '{
    "Hash":"85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002"
}'
```

Example Response
```
{
    "Hash": "85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002",
    "User": "ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f",
    "SrcChainId": 2,
    "BlockHeight": 9469807,
    "Time": 1610695305,
    "DstChainId": 79,
    "FeeAmount": "10000000000000000",
    "TransferAmount": "90000000000000000",
    "DstUser": "6e43f9988f2771f1a2b140cb3faad424767d39fc",
    "State": 0,
    "Token": {
        "Hash": "0000000000000000000000000000000000000000",
        "ChainId": 2,
        "Name": "Ethereum",
        "TokenBasicName": "Ethereum",
        "TokenBasic": null,
        "TokenMaps": null
    },
    "TransactionState": [
        {
            "Hash": "85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002",
            "ChainId": 2,
            "Blocks": 10,
            "NeedBlocks": 10,
            "Time": 1610695305
        },
        {
            "Hash": "a58b5705c2117e390c7add98d55e762342c26508a9b787befa228e5c10a2b14f",
            "ChainId": 0,
            "Blocks": 10,
            "NeedBlocks": 10,
            "Time": 1610697074
        },
        {
            "Hash": "5e201266b11f107dafa8e323b4be3b1c7f062bc1f1926ce36cf8832497342e37",
            "ChainId": 79,
            "Blocks": 10,
            "NeedBlocks": 10,
            "Time": 1610697089
        }
    ]
}
```

### POST transactionsofstate

Request 
```
http://localhost:8080/v1/transactionsofstate/
```

BODY raw
```
{
    "State":0,
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofstate/' \
--data-raw '{
    "State":0,
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
    "TotalCount": 1,
    "Transactions": [
        {
            "Hash": "85d1b5a97ae1a16e4507bc20e55c17426af6fcf5c35ef177e333148b601f1002",
            "User": "ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f",
            "SrcChainId": 2,
            "BlockHeight": 9469807,
            "Time": 1610695305,
            "DstChainId": 79,
            "FeeTokenHash": "0000000000000000000000000000000000000000",
            "FeeAmount": 10000000000000000,
            "State": 0
        }
    ]
}
```