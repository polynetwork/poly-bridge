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
* [POST transactoins](#post-transactoins)
* [POST transactoinsofuser](#post-transactoinsofuser)
* [POST transactoinsofstate](#post-transactoinsofstate)

## Test Node
[40.115.153.174:30330](http://40.115.153.174:30330/v1/)

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

### POST transactoins

Request 
```
http://localhost:8080/v1/transactoins/
```

BODY raw
```
{
    "PageNo":1,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactoins/' \
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
            "State": "finished"
        }
    ]
}
```

### POST transactoinsofuser

Request 
```
http://localhost:8080/v1/transactoinsofuser/
```

BODY raw
```
{
    "Addresses":["8bc7e7304120b88d111431f6a4853589d10e8132", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"],
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactoinsofuser/' \
--data-raw '{
    "Addresses":["8bc7e7304120b88d111431f6a4853589d10e8132", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"],
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
    "Transactions": [
        {
            "Hash": "336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729",
            "User": "ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f",
            "SrcChainId": 2,
            "BlockHeight": 9329385,
            "Time": 1608885420,
            "DstChainId": 4,
            "FeeTokenHash": "0000000000000000000000000000000000000000",
            "FeeAmount": "1000000000000000000000000000000",
            "Amount": "9000000000000000000000000000000",
            "DstUser": "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8",
            "State": "Finished",
            "TransactionState": [
                {
                    "Hash": "336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c729",
                    "ChainId": 2,
                    "Blocks": 0,
                    "Time": 1608885420
                },
                {
                    "Hash": "d2e8e325265ed314d9f538c2cb3f8e0a71ca2adad8b31db98278a4af6aecc1df",
                    "ChainId": 0,
                    "Blocks": 10,
                    "Time": 1609143919
                },
                {
                    "Hash": "1eecbe19ea85bd3ea57c3e7496b89f5263aa43d57449e21726f98435904ca7c4",
                    "ChainId": 4,
                    "Blocks": 0,
                    "Time": 1608896077
                }
            ]
        },
        {
            "Hash": "336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c739",
            "User": "ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f",
            "SrcChainId": 2,
            "BlockHeight": 9329385,
            "Time": 1608885420,
            "DstChainId": 4,
            "FeeTokenHash": "0000000000000000000000000000000000000000",
            "FeeAmount": "1000000000000000000000000000000",
            "Amount": "9000000000000000000000000000000",
            "DstUser": "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8",
            "State": "source done",
            "TransactionState": [
                {
                    "Hash": "336cd94f1ec80280c684606b8c9358f1ad0e9e7e7ce69f0da35c21a66fa0c739",
                    "ChainId": 2,
                    "Blocks": 0,
                    "Time": 1608885420
                }
            ]
        }
    ]
}
```

### POST transactoinsofstate

Request 
```
http://localhost:8080/v1/transactoinsofstate/
```

BODY raw
```
{
    "State":"finished",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactoinsofstate/' \
--data-raw '{
    "State":"finished",
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
            "State": "finished"
        }
    ]
}
```