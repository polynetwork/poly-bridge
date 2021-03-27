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

查询服务的状态

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

### POST tokenbasics

获取目前可以跨链的资产信息。

TokenBasic是币的本位信息，各链上不同的usdt，都对应usdt为本位。
Token为本位币在各个链上的详细信息，如eth上的usdt，bsc上的usdt。
TokenMap为币之间的跨链映射关系，如eth上的usdt可以跨链到bsc上的usdt，则有eth的usdt到bsc的usdt的映射，如bsc上的usdt不可以跨链到eth上的usdt，则没有bsc的usdt到eth的usdt的映射。

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
    "TotalCount": 20,
    "TokenBasics": [
        {
            "Name": "CWS",
            "Precision": 18,
            "Price": "46.35938146",
            "Ind": 1,
            "Time": 1616803238,
            "Property": 1,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "ac0104cca91d167873b8601d2e71eb3d4d8c33e0",
                    "ChainId": 2,
                    "Name": "CWS",
                    "Property": 1,
                    "TokenBasicName": "CWS",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "bcf39f0edda668c58371e519af37ca705f2bfcbd",
                    "ChainId": 6,
                    "Name": "CWS",
                    "Property": 1,
                    "TokenBasicName": "CWS",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "DAI",
            "Precision": 18,
            "Price": "1.00509312",
            "Ind": 1,
            "Time": 1616803238,
            "Property": 1,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "643f3914fb8ede03d932c79732746a8c11ae470a",
                    "ChainId": 7,
                    "Name": "pDAI",
                    "Property": 1,
                    "TokenBasicName": "DAI",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "6b175474e89094c44da98b954eedeac495271d0f",
                    "ChainId": 2,
                    "Name": "DAI",
                    "Property": 0,
                    "TokenBasicName": "DAI",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "7b956c0c11fcffb9c9227ca1925ba4c3486b36f1",
                    "ChainId": 3,
                    "Name": "DAI",
                    "Property": 1,
                    "TokenBasicName": "DAI",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "8f339abc2a2a8a4d0364c7e35f892c40fbfb4bc0",
                    "ChainId": 6,
                    "Name": "pDAI",
                    "Property": 1,
                    "TokenBasicName": "DAI",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "ETH",
            "Precision": 18,
            "Price": "1704.503055",
            "Ind": 1,
            "Time": 1616803238,
            "Property": 1,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000000000000000000000000",
                    "ChainId": 2,
                    "Name": "ETH",
                    "Property": 0,
                    "TokenBasicName": "ETH",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "8c0859c191d8f100e4a3c0d8c0066c36a0c1f894",
                    "ChainId": 7,
                    "Name": "pETH",
                    "Property": 1,
                    "TokenBasicName": "ETH",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "b9478391eec218defa96f7b9a7938cf44e7a2fd5",
                    "ChainId": 6,
                    "Name": "pETH",
                    "Property": 1,
                    "TokenBasicName": "ETH",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "df19600d334bb13c6a9e3e9777aa8ec6ed6a4a79",
                    "ChainId": 3,
                    "Name": "ETH",
                    "Property": 1,
                    "TokenBasicName": "ETH",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "FLM",
            "Precision": 8,
            "Price": "0.45932281",
            "Ind": 1,
            "Time": 1616803238,
            "Property": 1,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "4d9eab13620fe3569ba3b0e56e2877739e4145e3",
                    "ChainId": 4,
                    "Name": "Flamingo",
                    "Property": 1,
                    "TokenBasicName": "FLM",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "a0e910ce120d6220ceb3ad0000dbb4843eb912f5",
                    "ChainId": 7,
                    "Name": "Flamingo",
                    "Property": 1,
                    "TokenBasicName": "FLM",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "nNEO",
            "Precision": 8,
            "Price": "41.74652049",
            "Ind": 1,
            "Time": 1616803238,
            "Property": 1,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "6514a5ebff7944099591ae3e8a5c0979c83b2571",
                    "ChainId": 7,
                    "Name": "pNEO",
                    "Property": 1,
                    "TokenBasicName": "nNEO",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "9a576d927dda934b8ce69f35ec2c1025ceb10e6f",
                    "ChainId": 3,
                    "Name": "nNEO",
                    "Property": 1,
                    "TokenBasicName": "nNEO",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "b119b3b8e5e6eeffbe754b20ee5b8a42809931fb",
                    "ChainId": 6,
                    "Name": "pNEO",
                    "Property": 1,
                    "TokenBasicName": "nNEO",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "f46719e2d16bf50cddcef9d4bbfece901f73cbb6",
                    "ChainId": 4,
                    "Name": "nNEO",
                    "Property": 1,
                    "TokenBasicName": "nNEO",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        }
    ]
}
```

### POST tokens

获取一条链上，目前支持跨链的资产列表。

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

获取一条链上的一个跨链资产的信息。

Request 
```
http://localhost:8080/v1/token/
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
curl --location --request POST 'http://localhost:8080/v1/token/' \
--data-raw '{
   "ChainId": 2,
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

查询一条链上一个资产的跨链到目标链以及资产的映射关系。

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

查询一条链上一个资产可以被跨链过来的源链以及资产的映射关系。

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

### POST getfee

获取到一个指定链的跨链需要收取的源链上资产金额。
用户在做一次跨链操作时，在源链交易上收取了用户的费用。

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

检查一个源链交易是否支付了满足要求的手续费。

Request 
```
http://localhost:8080/v1/checkfee/
```

BODY raw
```
{
    "Checks":[{"Hash":"0000000000000000000000000000000000000000000000000000000000000024","ChainId":6},{"Hash":"0000000000000000000000000000000000000000000000000000000000000e72","ChainId":79}]
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/checkfee/' \
--data-raw '{
    "Checks":[{"Hash":"0000000000000000000000000000000000000000000000000000000000000024","ChainId":6},{"Hash":"0000000000000000000000000000000000000000000000000000000000000e72","ChainId":79}]
}'
```

Example Response
```
{
    "TotalCount": 2,
    "CheckFees": [
        {
            "ChainId": 6,
            "Hash": "0000000000000000000000000000000000000000000000000000000000000024",
            "PayState": -1,
            "Amount": "0",
            "MinProxyFee": "0"
        },
        {
            "ChainId": 79,
            "Hash": "0000000000000000000000000000000000000000000000000000000000000e72",
            "PayState": -1,
            "Amount": "0.3881888898738662",
            "MinProxyFee": "0.5590046679671925"
        }
    ]
}
```

### POST transactions

获取跨链交易列表。

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

获取指定地址上的跨链交易列表。

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
获取指定hash的跨链交易。

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

获取指定状态的跨链交易。

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