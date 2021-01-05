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
    "TotalCount": 4,
    "Tokens": [
        {
            "Hash": "0000000000000000000200000000000000000000",
            "ChainId": 2,
            "Name": "Ethereum",
            "TokenBasicName": "Ethereum",
            "TokenBasic": {
                "Name": "Ethereum",
                "Precision": 0,
                "Price": 0,
                "Ind": 0,
                "Time": 1609718410,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000000",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000400000000000000000000",
                    "DstToken": {
                        "Hash": "0000000000000000000400000000000000000000",
                        "ChainId": 4,
                        "Name": "NETH",
                        "TokenBasicName": "Ethereum",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                },
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000000",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000800000000000000000000",
                    "DstToken": {
                        "Hash": "0000000000000000000800000000000000000000",
                        "ChainId": 8,
                        "Name": "NETH",
                        "TokenBasicName": "Ethereum",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "0000000000000000000200000000000000000001",
            "ChainId": 2,
            "Name": "ENeo",
            "TokenBasicName": "Neo",
            "TokenBasic": {
                "Name": "Neo",
                "Precision": 0,
                "Price": 0,
                "Ind": 0,
                "Time": 1609718410,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000001",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000400000000000000000001",
                    "DstToken": {
                        "Hash": "0000000000000000000400000000000000000001",
                        "ChainId": 4,
                        "Name": "nNeo",
                        "TokenBasicName": "Neo",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                },
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000001",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000800000000000000000001",
                    "DstToken": {
                        "Hash": "0000000000000000000800000000000000000001",
                        "ChainId": 8,
                        "Name": "nNeo",
                        "TokenBasicName": "Neo",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "0000000000000000000200000000000000000002",
            "ChainId": 2,
            "Name": "USDT",
            "TokenBasicName": "USDT",
            "TokenBasic": {
                "Name": "USDT",
                "Precision": 0,
                "Price": 0,
                "Ind": 0,
                "Time": 1609675216,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000002",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000400000000000000000002",
                    "DstToken": {
                        "Hash": "0000000000000000000400000000000000000002",
                        "ChainId": 4,
                        "Name": "USDT",
                        "TokenBasicName": "USDT",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                },
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000002",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000800000000000000000002",
                    "DstToken": {
                        "Hash": "0000000000000000000800000000000000000002",
                        "ChainId": 8,
                        "Name": "USDT",
                        "TokenBasicName": "USDT",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                }
            ]
        },
        {
            "Hash": "0000000000000000000200000000000000000004",
            "ChainId": 2,
            "Name": "BNB",
            "TokenBasicName": "BNB",
            "TokenBasic": {
                "Name": "BNB",
                "Precision": 0,
                "Price": 0,
                "Ind": 0,
                "Time": 1609718410,
                "PriceMarkets": null,
                "Tokens": null
            },
            "TokenMaps": [
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000004",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000400000000000000000004",
                    "DstToken": {
                        "Hash": "0000000000000000000400000000000000000004",
                        "ChainId": 4,
                        "Name": "BNB",
                        "TokenBasicName": "BNB",
                        "TokenBasic": null,
                        "TokenMaps": null
                    }
                },
                {
                    "SrcTokenHash": "0000000000000000000200000000000000000004",
                    "SrcToken": null,
                    "DstTokenHash": "0000000000000000000800000000000000000004",
                    "DstToken": {
                        "Hash": "0000000000000000000800000000000000000004",
                        "ChainId": 8,
                        "Name": "BNB",
                        "TokenBasicName": "BNB",
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
   "Hash": "0000000000000000000200000000000000000002"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/token/' \
--data-raw '{
   "Hash": "0000000000000000000200000000000000000002"
}'
```

Example Response
```
{
    "Hash": "0000000000000000000200000000000000000002",
    "ChainId": 2,
    "Name": "USDT",
    "TokenBasicName": "USDT",
    "TokenBasic": {
        "Name": "USDT",
        "Precision": 0,
        "Price": 0,
        "Ind": 0,
        "Time": 1609844520,
        "PriceMarkets": null,
        "Tokens": null
    },
    "TokenMaps": [
        {
            "SrcTokenHash": "0000000000000000000200000000000000000002",
            "SrcToken": null,
            "DstTokenHash": "0000000000000000000400000000000000000002",
            "DstToken": {
                "Hash": "0000000000000000000400000000000000000002",
                "ChainId": 4,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        },
        {
            "SrcTokenHash": "0000000000000000000200000000000000000002",
            "SrcToken": null,
            "DstTokenHash": "0000000000000000000800000000000000000002",
            "DstToken": {
                "Hash": "0000000000000000000800000000000000000002",
                "ChainId": 8,
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
    "Hash": "0000000000000000000200000000000000000002"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokenmap/' \
--data-raw '{
    "ChainId": 2,
    "Hash": "0000000000000000000200000000000000000002"
}'
```

Example Response
```
{
    "TotalCount": 2,
    "TokenMaps": [
        {
            "SrcTokenHash": "0000000000000000000200000000000000000002",
            "SrcToken": {
                "Hash": "0000000000000000000200000000000000000002",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "0000000000000000000400000000000000000002",
            "DstToken": {
                "Hash": "0000000000000000000400000000000000000002",
                "ChainId": 4,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        },
        {
            "SrcTokenHash": "0000000000000000000200000000000000000002",
            "SrcToken": {
                "Hash": "0000000000000000000200000000000000000002",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "0000000000000000000800000000000000000002",
            "DstToken": {
                "Hash": "0000000000000000000800000000000000000002",
                "ChainId": 8,
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
    "Hash": "0000000000000000000200000000000000000002"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokenmapreverse/' \
--data-raw '{
    "ChainId": 2,
    "Hash": "0000000000000000000200000000000000000002"
}'
```

Example Response
```
{
    "TotalCount": 2,
    "TokenMaps": [
        {
            "SrcTokenHash": "0000000000000000000400000000000000000002",
            "SrcToken": {
                "Hash": "0000000000000000000400000000000000000002",
                "ChainId": 4,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "0000000000000000000200000000000000000002",
            "DstToken": {
                "Hash": "0000000000000000000200000000000000000002",
                "ChainId": 2,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            }
        },
        {
            "SrcTokenHash": "0000000000000000000800000000000000000002",
            "SrcToken": {
                "Hash": "0000000000000000000800000000000000000002",
                "ChainId": 8,
                "Name": "USDT",
                "TokenBasicName": "USDT",
                "TokenBasic": null,
                "TokenMaps": null
            },
            "DstTokenHash": "0000000000000000000200000000000000000002",
            "DstToken": {
                "Hash": "0000000000000000000200000000000000000002",
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
    "TotalCount": 5,
    "TokenBasics": [
        {
            "Name": "BNB",
            "Precision": 0,
            "Price": 0,
            "Ind": 0,
            "Time": 1609844520,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000200000000000000000004",
                    "ChainId": 2,
                    "Name": "BNB",
                    "TokenBasicName": "BNB",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000400000000000000000004",
                    "ChainId": 4,
                    "Name": "BNB",
                    "TokenBasicName": "BNB",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000800000000000000000004",
                    "ChainId": 8,
                    "Name": "BNB",
                    "TokenBasicName": "BNB",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "Ethereum",
            "Precision": 0,
            "Price": 0,
            "Ind": 0,
            "Time": 1609844520,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000200000000000000000000",
                    "ChainId": 2,
                    "Name": "Ethereum",
                    "TokenBasicName": "Ethereum",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000400000000000000000000",
                    "ChainId": 4,
                    "Name": "NETH",
                    "TokenBasicName": "Ethereum",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000800000000000000000000",
                    "ChainId": 8,
                    "Name": "NETH",
                    "TokenBasicName": "Ethereum",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "GAS",
            "Precision": 0,
            "Price": 0,
            "Ind": 0,
            "Time": 1609844520,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000400000000000000000003",
                    "ChainId": 4,
                    "Name": "GAS",
                    "TokenBasicName": "GAS",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "Neo",
            "Precision": 0,
            "Price": 0,
            "Ind": 0,
            "Time": 1609844520,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000200000000000000000001",
                    "ChainId": 2,
                    "Name": "ENeo",
                    "TokenBasicName": "Neo",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000400000000000000000001",
                    "ChainId": 4,
                    "Name": "nNeo",
                    "TokenBasicName": "Neo",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000800000000000000000001",
                    "ChainId": 8,
                    "Name": "nNeo",
                    "TokenBasicName": "Neo",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ]
        },
        {
            "Name": "USDT",
            "Precision": 0,
            "Price": 0,
            "Ind": 0,
            "Time": 1609844520,
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "0000000000000000000200000000000000000002",
                    "ChainId": 2,
                    "Name": "USDT",
                    "TokenBasicName": "USDT",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000400000000000000000002",
                    "ChainId": 4,
                    "Name": "USDT",
                    "TokenBasicName": "USDT",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "0000000000000000000800000000000000000002",
                    "ChainId": 8,
                    "Name": "USDT",
                    "TokenBasicName": "USDT",
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
    "ChainId": 2,
    "Hash": "0000000000000000000200000000000000000002"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/getfee/' \
--data-raw '{
    "ChainId": 2,
    "Hash": "0000000000000000000200000000000000000002"
}'
```

Example Response
```
{
    "ChainId": 2,
    "Hash": "0000000000000000000200000000000000000002",
    "Amount": 23538314.17111664
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
    "Hash": "0000000000000000000000000000000000000000"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/checkfee/' \
--data-raw '{
    "Hash": "0000000000000000000000000000000000000000"
}'
```

Example Response
```
{
    "Hash": "0000000000000000000000000000000000000000",
    "hasPay": true,
    "Amount": 23538314.17111664
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
    "PageNo":1,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 1,
    "TotalPage": 0,
    "TotalCount": 0,
    "Transactions": null
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
    "State":"Finished",
    "PageNo":1,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactoinsofstate/' \
--data-raw '{
    "State":"Finished",
    "PageNo":1,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 1,
    "TotalPage": 0,
    "TotalCount": 0,
    "Transactions": null
}
```