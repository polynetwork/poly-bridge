# Poly Bridge

This file involves essential APIs provided by PolyBridge and describes how to deploy these APIs. 

## API

* [GET /](#get-/)
* [POST tokens](#post-tokens)
* [POST token](#post-token)
* [POST tokenbasics](#post-tokenbasics)
* [POST tokenbasicsinfo](#post-tokenbasicsinfo)
* [POST tokenmap](#post-tokenmap)
* [POST tokenmapreverse](#post-tokenmapreverse)
* [POST getfee](#post-getfee)
* [POST checkfee](#post-checkfee)
* [POST transactions](#post-transactions)
* [POST transactionswithfilter](#post-transactionswithfilter)
* [POST transactionsofaddress](#post-transactionsofaddress)
* [POST transactionofhash](#post-transactionofhash)
* [POST transactionsofstate](#post-transactionsofstate)
* [POST transactionofcurve](#post-transactionofcurve)
* [POST transactionsofunfinished](#post-transactionsofunfinished)
* [POST transactionsofasset](#post-transactionsofasset)
* [POST expecttime](#post-expecttime)

## Test Node
[testnet](https://bridge.poly.network/testnet/v1/)

[mainnet](https://bridge.poly.network/v1/)

## Status Codes of Transaction

Code|Message
:--:|:--:
0|finished
1|pendding
2|source done
3|source confirmed
4|poly confirmed

## Charges of Cross-Chain Transaction

### Accounting
The fee charged by agent = The transaction fee on target chain * 120% （the rate, 120%, is configurable）

e.g., if BNB of chain BSC is transferred onto ETH, 
the fee = (eth.gas_limit * eth.gas_price) * (eth to USDT) / (BNB to USDT)

**Attention:** The asset charged is BNB. 

### Checking

hasPay = Charges > The transaction fee on target chain * 20% （the rate, 20%, is configurable）

e.g., if checking the fee charged during BNB transaction between BSC and ETH,

hasPay = BNB charged * (BNB to USDT) > (eth.gas_limit * eth.gas_price) * (eth to USDT) * 20%

## API Info

Status querying is shown in the following. 

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

This API enables you to check assets that can be transferred cross chains currently.

TokenBasic shows standard coin, e.g., USDT. 
Token refers to the detailed information of standard coin in each chains, e.g., USDT in ETH and BSC chains.
TokenMap refers to the mapping relations among different assets, e.g., it will show the mapping of USDT between ETH and BSC if USDT can be transferred from ETH to BSC. 

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


### POST tokenbasicsinfo

This API enables you to check information of statistic and assets that can be transferred cross chains currently.

TokenBasic shows standard coin, e.g., USDT.
Token refers to the detailed information of standard coin in each chains, e.g., USDT in ETH and BSC chains.
TokenMap refers to the mapping relations among different assets, e.g., it will show the mapping of USDT between ETH and BSC if USDT can be transferred from ETH to BSC.

Request 
```
http://localhost:8080/v1/tokenbasicsinfo/
```

BODY raw
```
{
  "PageSize":10,
  "PageNo":5,
  "Order": "total_count"
}

# Valid order keys: "total_amount", "total_count", "name", "price"
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/tokenbasicsinfo/' \
--data-raw '{
}'
```

Example Response
```
{
    "PageSize": 1,
    "PageNo": 1,
    "TotalPage": 2,
    "TotalCount": 52,
    "TokenBasics": [
        {
            "Name": "PKR",
            "Precision": 18,
            "Price": "1",
            "Ind": 1,
            "Time": 0,
            "Property": 1,
            "Meta": "",
            "PriceMarkets": null,
            "Tokens": [
                {
                    "Hash": "c05f2a6c6aac12bc662ed54a3a89f77ec7eef56d",
                    "ChainId": 79,
                    "Name": "PKR",
                    "Property": 1,
                    "TokenBasicName": "PKR",
                    "Precision": 18,
                    "AvailableAmount": "50015008100000000000",
                    "TokenBasic": null,
                    "TokenMaps": null
                },
                {
                    "Hash": "e65ef1aea76e62ab030e98c962255f37294f1648",
                    "ChainId": 2,
                    "Name": "PKR",
                    "Property": 1,
                    "TokenBasicName": "PKR",
                    "Precision": 18,
                    "AvailableAmount": "1000049473771230000000000",
                    "TokenBasic": null,
                    "TokenMaps": null
                }
            ],
            "TotalAmount": "999999981275279900016640",
            "TotalVolume": "999999",
            "TotalCount": 2,
            "SocialTwitter": "",
            "SocialTelegram": "",
            "SocialWebsite": "",
            "SocialOther": ""
        }
    ]
}
```

### POST tokens

This API lists tokens that are transferable across chains currently on assigned chain. 

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

This API shows assigned asset status on assigned chain.

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

This API returns the mapping relations of assigned token between chains. 

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

This API returns the source chains supporting cross chain transaction of assigned token onto this chain and mapping relations of assets. 

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

This API returns transaction fee which will be charged on the source chain in cross-chain transaction.
And if SwapTokenHash is specified, the transferable amount will be returned.

Request 
```
http://localhost:8080/v1/getfee/
```

BODY raw
```
{
    "SrcChainId": 7, 
    "Hash": "0000000000000000000000000000000000000000", 
    "SwapTokenHash": "6ef070cb10fc9f66d04a4c387928b268f55b9198", 
    "DstChainId": 5
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/getfee/' \
--data-raw '{
    "SrcChainId": 7, 
    "Hash": "0000000000000000000000000000000000000000", 
    "SwapTokenHash": "6ef070cb10fc9f66d04a4c387928b268f55b9198", 
    "DstChainId": 5
}'
```

Example Response
```
{
    "SrcChainId": 7,
    "Hash": "0000000000000000000000000000000000000000",
    "DstChainId": 5,
    "UsdtAmount": "3.7505409912",
    "TokenAmount": "0.23261657",
    "TokenAmountWithPrecision": "232616561965745400",
    "SwapTokenHash": "6ef070cb10fc9f66d04a4c387928b268f55b9198",
    "Balance": "12.45323704",
    "BalanceWithPrecision": "1245323704"
}
```

### POST checkfee

This API is used to check whether the source transaction pays required fee.

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

This API returns cross chain transaction lists.

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

### POST transactionswithfilter

This API returns cross chain transaction lists, supporting parameter: SrcChainId, DstChainId, and Assets.

Request 
```
http://localhost:8080/v1/transactionswithfilter/
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":10,
    "SrcChainId": 2,
    "DstChainId: 79,
    "Assets": ["155040625d7ae3e9cada9a73e3e44f76d3ed1409"],
    "Addresses":["ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"]
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionswithfilter/' \
--data-raw '{
    "PageNo":0,
    "PageSize":10,
    "SrcChainId": 2,
    "DstChainId: 79,
    "Assets": ["155040625d7ae3e9cada9a73e3e44f76d3ed1409"],
    "Addresses":["ad79c606bd4ef330ac45df9d2ace4e7e7c6db13f", "ARpuQar5CPtxEoqfcg1fxGWnwDdp7w3jj8"]
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

### POST transactionsofaddress

This API returns the cross-chain history of the specified address.

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

This API returns the details of the specified hash, and you can view the cross-chain progress through the TransactionState in the response.

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

This API returns the details of the specified hash, and you can view the cross-chain progress through the TransactionState in the response.

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

### POST transactionofcurve

This API returns the detailed information of the whole transaction according to the hash in curve.

Request 
```
http://localhost:8080/v1/transactionofcurve/
```

BODY raw
```
{
    "Hash":"6cd2b1997f3ef4cecce76a378c1fb0ebc75573f1aeb2ad7bcf49351eecff145e"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionofcurve/' \
--data-raw '{
    "Hash":"6cd2b1997f3ef4cecce76a378c1fb0ebc75573f1aeb2ad7bcf49351eecff145e"
}'
```

Example Response
```
{
    "Hash": "7fac27c0bdeea45cb0d9e7f666bf2a0c77e7c54ce4f08ded30e8c6afef069536",
    "User": "9f6a12f81458594f693fb78789a9c6640996da0c",
    "SrcChainId": 7,
    "BlockHeight": 4151127,
    "Time": 1619347047,
    "DstChainId": 7,
    "Amount": "18635869611217890000",
    "FeeAmount": "308010000000000",
    "TransferAmount": "18635561601217890000",
    "DstUser": "34d4a23a1fc0c694f0d74ddaf9d8d564cfe2d430",
    "ServerId": 1,
    "State": 0,
    "Token": null,
    "TransactionState": [
        {
            "Hash": "7fac27c0bdeea45cb0d9e7f666bf2a0c77e7c54ce4f08ded30e8c6afef069536",
            "ChainId": 7,
            "Blocks": 21,
            "NeedBlocks": 21,
            "Time": 1619347047
        },
        {
            "Hash": "4bd9c760148e8e4cf4dfca50d357b8ea139ebceb7ca2a57db65eae0b2fb25c49",
            "ChainId": 0,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1619347134
        },
        {
            "Hash": "6cd2b1997f3ef4cecce76a378c1fb0ebc75573f1aeb2ad7bcf49351eecff145e",
            "ChainId": 10,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1619347147
        },
        {
            "Hash": "edcba7668ed7b9f4663d1326c0e42d2ad22978871727ace4200b20b38875f7fc",
            "ChainId": 0,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1619347214
        },
        {
            "Hash": "bf6cc7b030ce3d1cb6288c83b34e4522f338412730907e1a95a03abb31ee1ecb",
            "ChainId": 7,
            "Blocks": 1,
            "NeedBlocks": 1,
            "Time": 1619347227
        }
    ]
}
```

### POST transactionsofunfinished

This API returns the unfinished transaction.

Request 
```
http://localhost:8080/v1/transactionsofunfinished/
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
curl --location --request POST 'http://localhost:8080/v1/transactionsofunfinished/' \
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
    "TotalPage": 3,
    "TotalCount": 29,
    "Transactions": [
        {
            "WrapperTransaction": {
                "Hash": "28301075a7fcd1e54b06ef591fd02870d139e2981fe279fa4c8b3fd4d2edd254",
                "User": "d7c7cb2532ead8b81e7d0e920a457cbc2fa30788",
                "SrcChainId": 7,
                "BlockHeight": 4521078,
                "Time": 1620456900,
                "DstChainId": 7,
                "DstUser": "d7c7cb2532ead8b81e7d0e920a457cbc2fa30788",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "3080100",
                "State": 2
            },
            "SrcTransaction": {
                "Hash": "28301075a7fcd1e54b06ef591fd02870d139e2981fe279fa4c8b3fd4d2edd254",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456900,
                "Height": 4521078,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "28301075a7fcd1e54b06ef591fd02870d139e2981fe279fa4c8b3fd4d2edd254",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "0298c2b32eae4da002a15f36fdf7615bea3da047",
                    "Amount": "0.02543336",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "0298c2b32eae4da002a15f36fdf7615bea3da047",
                "ChainId": 7,
                "Name": "USDT",
                "Property": 0,
                "TokenBasicName": "USDT",
                "TokenBasic": {
                    "Name": "USDT",
                    "Precision": 6,
                    "Price": "0.99989002",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "f800402fe1ccd9d7c4de31563cbe43a1edbadd22302c2bbcb707006d47d907df",
                "User": "12055de6152a50788025c9ce6c475a4d40b1f5b8",
                "SrcChainId": 6,
                "BlockHeight": 7235584,
                "Time": 1620456884,
                "DstChainId": 7,
                "DstUser": "12055de6152a50788025c9ce6c475a4d40b1f5b8",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "0.00001302",
                "State": 2
            },
            "SrcTransaction": {
                "Hash": "f800402fe1ccd9d7c4de31563cbe43a1edbadd22302c2bbcb707006d47d907df",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620456884,
                "Height": 7235584,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "f800402fe1ccd9d7c4de31563cbe43a1edbadd22302c2bbcb707006d47d907df",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "e9e7cea3dedca5984780bafc599bd69add087d56",
                    "Amount": "0.0062545451255415",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "e9e7cea3dedca5984780bafc599bd69add087d56",
                "ChainId": 6,
                "Name": "USDT",
                "Property": 0,
                "TokenBasicName": "USDT",
                "TokenBasic": {
                    "Name": "USDT",
                    "Precision": 6,
                    "Price": "0.99989002",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                "User": "6bc7e8d6d0361f71e82f49d0648265e0c229b138",
                "SrcChainId": 7,
                "BlockHeight": 4521069,
                "Time": 1620456873,
                "DstChainId": 6,
                "DstUser": "6bc7e8d6d0361f71e82f49d0648265e0c229b138",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 2
            },
            "SrcTransaction": {
                "Hash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456873,
                "Height": 4521069,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620456873,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "334.85554514",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "f30311fa06a4936dcf07f3c858c77cb709a41ff3f908785dcc5261e20502532a",
                "User": "768f2a7ccdfde9ebdfd5cea8b635dd590cb3a3f1",
                "SrcChainId": 7,
                "BlockHeight": 4521063,
                "Time": 1620456855,
                "DstChainId": 2,
                "DstUser": "768f2a7ccdfde9ebdfd5cea8b635dd590cb3a3f1",
                "ServerId": 0,
                "FeeTokenHash": "c38072aa3f8e049de541223a9c9772132bb48634",
                "FeeAmount": "16570610.9199985",
                "State": 2
            },
            "SrcTransaction": {
                "Hash": "f30311fa06a4936dcf07f3c858c77cb709a41ff3f908785dcc5261e20502532a",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456855,
                "Height": 4521063,
                "DstChainId": 2,
                "SrcTransfer": {
                    "TxHash": "f30311fa06a4936dcf07f3c858c77cb709a41ff3f908785dcc5261e20502532a",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620456855,
                    "Asset": "c38072aa3f8e049de541223a9c9772132bb48634",
                    "Amount": "80058865931.1425628907491116",
                    "DstChainId": 2,
                    "DstAsset": "95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "c38072aa3f8e049de541223a9c9772132bb48634",
                "ChainId": 7,
                "Name": "SHIB",
                "Property": 1,
                "TokenBasicName": "SHIB",
                "TokenBasic": {
                    "Name": "SHIB",
                    "Precision": 18,
                    "Price": "5.35e-06",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "36c534ff96666cb03f3c7d7a8740ac4a172fc9853177c283d81a373a6a14c9ba",
                "User": "e1238cbe00a07d663caa42aaef4ef6bdd90475fa",
                "SrcChainId": 7,
                "BlockHeight": 4521061,
                "Time": 1620456849,
                "DstChainId": 2,
                "DstUser": "fd24a793936c25c36e846ebdfccfc71a7add3c35",
                "ServerId": 0,
                "FeeTokenHash": "c38072aa3f8e049de541223a9c9772132bb48634",
                "FeeAmount": "16570610.9199985",
                "State": 2
            },
            "SrcTransaction": {
                "Hash": "36c534ff96666cb03f3c7d7a8740ac4a172fc9853177c283d81a373a6a14c9ba",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456849,
                "Height": 4521061,
                "DstChainId": 2,
                "SrcTransfer": {
                    "TxHash": "36c534ff96666cb03f3c7d7a8740ac4a172fc9853177c283d81a373a6a14c9ba",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620456849,
                    "Asset": "c38072aa3f8e049de541223a9c9772132bb48634",
                    "Amount": "125441658.9372152242016528",
                    "DstChainId": 2,
                    "DstAsset": "95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "c38072aa3f8e049de541223a9c9772132bb48634",
                "ChainId": 7,
                "Name": "SHIB",
                "Property": 1,
                "TokenBasicName": "SHIB",
                "TokenBasic": {
                    "Name": "SHIB",
                    "Precision": 18,
                    "Price": "5.35e-06",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "7197df27e89546e3c5e86644d0eb3653baef555bbc321e36c93768500cdf79d1",
                "User": "d756577b49888c92f6c9761ab678d971292df4d6",
                "SrcChainId": 7,
                "BlockHeight": 4521059,
                "Time": 1620456843,
                "DstChainId": 7,
                "DstUser": "d756577b49888c92f6c9761ab678d971292df4d6",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "308010000000000",
                "State": 3
            },
            "SrcTransaction": {
                "Hash": "7197df27e89546e3c5e86644d0eb3653baef555bbc321e36c93768500cdf79d1",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456843,
                "Height": 4521059,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "7197df27e89546e3c5e86644d0eb3653baef555bbc321e36c93768500cdf79d1",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "2ec96bb06e6af8c8ac20f93c34ea2ab663e40d62",
                    "Amount": "",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": null
        },
        {
            "WrapperTransaction": {
                "Hash": "39ecc77b1b393591a9817d27f52404cbf7d35e6ff082ba2b109b1fc8d03cfdc2",
                "User": "c71e8407912243e5afe0af814a89f0ffadd2263f",
                "SrcChainId": 7,
                "BlockHeight": 4521058,
                "Time": 1620456840,
                "DstChainId": 7,
                "DstUser": "c71e8407912243e5afe0af814a89f0ffadd2263f",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "308010000000000",
                "State": 3
            },
            "SrcTransaction": {
                "Hash": "39ecc77b1b393591a9817d27f52404cbf7d35e6ff082ba2b109b1fc8d03cfdc2",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456840,
                "Height": 4521058,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "39ecc77b1b393591a9817d27f52404cbf7d35e6ff082ba2b109b1fc8d03cfdc2",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "2ec96bb06e6af8c8ac20f93c34ea2ab663e40d62",
                    "Amount": "",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": null
        },
        {
            "WrapperTransaction": {
                "Hash": "e24180b159898db559f4398417be5a3347e548c4627151c50ab0cadebe69fff1",
                "User": "ca75886eea9a95bebb68d3e2b5afa60d9ed19f31",
                "SrcChainId": 6,
                "BlockHeight": 7235568,
                "Time": 1620456836,
                "DstChainId": 6,
                "DstUser": "ca75886eea9a95bebb68d3e2b5afa60d9ed19f31",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "0.00154001",
                "State": 3
            },
            "SrcTransaction": {
                "Hash": "e24180b159898db559f4398417be5a3347e548c4627151c50ab0cadebe69fff1",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620456836,
                "Height": 7235568,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "e24180b159898db559f4398417be5a3347e548c4627151c50ab0cadebe69fff1",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "e9e7cea3dedca5984780bafc599bd69add087d56",
                    "Amount": "1.4",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "e9e7cea3dedca5984780bafc599bd69add087d56",
                "ChainId": 6,
                "Name": "USDT",
                "Property": 0,
                "TokenBasicName": "USDT",
                "TokenBasic": {
                    "Name": "USDT",
                    "Precision": 6,
                    "Price": "0.99989002",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "83171883591f0e263703e3f00e49e7ce957857baef25152aec69e5acbdc26be6",
                "User": "4b8fe418da8a18f6a40b66ccc4c56b2811b40387",
                "SrcChainId": 6,
                "BlockHeight": 7235567,
                "Time": 1620456833,
                "DstChainId": 7,
                "DstUser": "4b8fe418da8a18f6a40b66ccc4c56b2811b40387",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "0.00001302",
                "State": 3
            },
            "SrcTransaction": {
                "Hash": "83171883591f0e263703e3f00e49e7ce957857baef25152aec69e5acbdc26be6",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620456833,
                "Height": 7235567,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "83171883591f0e263703e3f00e49e7ce957857baef25152aec69e5acbdc26be6",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "e9e7cea3dedca5984780bafc599bd69add087d56",
                    "Amount": "3.1238929283800599",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "e9e7cea3dedca5984780bafc599bd69add087d56",
                "ChainId": 6,
                "Name": "USDT",
                "Property": 0,
                "TokenBasicName": "USDT",
                "TokenBasic": {
                    "Name": "USDT",
                    "Precision": 6,
                    "Price": "0.99989002",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "98cc84657c5db906e0c2153871ae17632ade9e536c9c79015c4c440bfbf41c7b",
                "User": "b9074d4f009e4b22e60405e851f80e4d94ef509c",
                "SrcChainId": 7,
                "BlockHeight": 4521051,
                "Time": 1620456819,
                "DstChainId": 7,
                "DstUser": "b9074d4f009e4b22e60405e851f80e4d94ef509c",
                "ServerId": 1,
                "FeeTokenHash": "0000000000000000000000000000000000000000",
                "FeeAmount": "3080100",
                "State": 4
            },
            "SrcTransaction": {
                "Hash": "98cc84657c5db906e0c2153871ae17632ade9e536c9c79015c4c440bfbf41c7b",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456819,
                "Height": 4521051,
                "DstChainId": 10,
                "SrcTransfer": {
                    "TxHash": "98cc84657c5db906e0c2153871ae17632ade9e536c9c79015c4c440bfbf41c7b",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 0,
                    "Asset": "0298c2b32eae4da002a15f36fdf7615bea3da047",
                    "Amount": "6",
                    "DstChainId": 10,
                    "DstAsset": "0000000000000000000000000000000000000000"
                }
            },
            "PolyTransaction": {
                "Hash": "dd1a7a51f8128bcd57b69f763df6ce582cb6088c85c3160caf3df5380542f0c5",
                "ChainId": 0,
                "State": 1,
                "Time": 1620456906,
                "Fee": "0",
                "Height": 7631391,
                "SrcChainId": 7,
                "SrcHash": "98cc84657c5db906e0c2153871ae17632ade9e536c9c79015c4c440bfbf41c7b",
                "DstChainId": 10,
                "Key": ""
            },
            "DstTransaction": null,
            "Token": {
                "Hash": "0298c2b32eae4da002a15f36fdf7615bea3da047",
                "ChainId": 7,
                "Name": "USDT",
                "Property": 0,
                "TokenBasicName": "USDT",
                "TokenBasic": {
                    "Name": "USDT",
                    "Precision": 6,
                    "Price": "0.99989002",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        }
    ]
}
```

### POST transactionsofasset

This API returns detailed information of transaction according to asset.

Request 
```
http://localhost:8080/v1/transactionsofasset/
```

BODY raw
```
{
    "Asset":"25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
    "Chain":7,
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofasset/' \
--data-raw '{
    "Asset":"25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
    "Chain":7,
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1814,
    "TotalCount": 18137,
    "Transactions": [
        {
            "WrapperTransaction": {
                "Hash": "61aa5d4ac42dbcf106d32766cd85bc24e353335fbf3a35f4fc749858e9427c34",
                "User": "88a949145a5748ed1a7895bcf10dbae6185cb9a4",
                "SrcChainId": 7,
                "BlockHeight": 4521139,
                "Time": 1620457083,
                "DstChainId": 6,
                "DstUser": "88a949145a5748ed1a7895bcf10dbae6185cb9a4",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 3
            },
            "SrcTransaction": {
                "Hash": "61aa5d4ac42dbcf106d32766cd85bc24e353335fbf3a35f4fc749858e9427c34",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620457083,
                "Height": 4521139,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "61aa5d4ac42dbcf106d32766cd85bc24e353335fbf3a35f4fc749858e9427c34",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620457083,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "374.5552575812177631",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": null,
            "DstTransaction": null,
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                "User": "6bc7e8d6d0361f71e82f49d0648265e0c229b138",
                "SrcChainId": 7,
                "BlockHeight": 4521069,
                "Time": 1620456873,
                "DstChainId": 6,
                "DstUser": "6bc7e8d6d0361f71e82f49d0648265e0c229b138",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456873,
                "Height": 4521069,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620456873,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "334.85554514",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "717507bc6776bb07a1ed4e0330bf8a9f4f32de2cb625ea83fefdb743198f5f50",
                "ChainId": 0,
                "State": 1,
                "Time": 1620456957,
                "Fee": "0",
                "Height": 7631433,
                "SrcChainId": 7,
                "SrcHash": "4901b95f1dbe215b4335e74571e3875b351cca1044b86e469843fafa264dd711",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "dde1f41abd274c8814f9807ab7425714b84de95ceabe92588beff7cb3d9872e0",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620456971,
                "Height": 7235613,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "dde1f41abd274c8814f9807ab7425714b84de95ceabe92588beff7cb3d9872e0",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620456971,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "a010c77b8ad3bc6e5ecd1b6c22f88d00e8aca93cbc6de522305617deb85f814d",
                "User": "8baf070d2907ed1b8ed7fe23f6a45b973ef9983d",
                "SrcChainId": 7,
                "BlockHeight": 4520947,
                "Time": 1620456507,
                "DstChainId": 6,
                "DstUser": "a24c8577b58b6028f52ed938c557353c405e7745",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "a010c77b8ad3bc6e5ecd1b6c22f88d00e8aca93cbc6de522305617deb85f814d",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620456507,
                "Height": 4520947,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "a010c77b8ad3bc6e5ecd1b6c22f88d00e8aca93cbc6de522305617deb85f814d",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620456507,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "47.8983521071394288",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "54afe25bdc3396337c48aeefac8d494a1cc1b14e90a512d8a08dddc3ede76eef",
                "ChainId": 0,
                "State": 1,
                "Time": 1620456627,
                "Fee": "0",
                "Height": 7631270,
                "SrcChainId": 7,
                "SrcHash": "a010c77b8ad3bc6e5ecd1b6c22f88d00e8aca93cbc6de522305617deb85f814d",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "0e61b190a4007e66493111b02ef879383f3e571d5e9ba6733ca1cb4b50dfd94a",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620456641,
                "Height": 7235503,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "0e61b190a4007e66493111b02ef879383f3e571d5e9ba6733ca1cb4b50dfd94a",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620456641,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "4f15d7683ce21784dbcb4d5e0f443be29ea00507c799a13068a974885e420bfa",
                "User": "84f53f9274b9aa0f7715dd72ad70b6f974dc4b9c",
                "SrcChainId": 7,
                "BlockHeight": 4520616,
                "Time": 1620455514,
                "DstChainId": 6,
                "DstUser": "84f53f9274b9aa0f7715dd72ad70b6f974dc4b9c",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "4f15d7683ce21784dbcb4d5e0f443be29ea00507c799a13068a974885e420bfa",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620455514,
                "Height": 4520616,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "4f15d7683ce21784dbcb4d5e0f443be29ea00507c799a13068a974885e420bfa",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620455514,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "65.2025000183634665",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "d94e094d37ee1e6cfaa240be2f5c90d0521b434d4c8b64975d6f93c07351a4bb",
                "ChainId": 0,
                "State": 1,
                "Time": 1620455618,
                "Fee": "0",
                "Height": 7630798,
                "SrcChainId": 7,
                "SrcHash": "4f15d7683ce21784dbcb4d5e0f443be29ea00507c799a13068a974885e420bfa",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "9d25269ada3b2458f3b5cf1817031870d2d35515f7119b13741cd4c8920237a3",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620455630,
                "Height": 7235166,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "9d25269ada3b2458f3b5cf1817031870d2d35515f7119b13741cd4c8920237a3",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620455630,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "8971defd59e7d6a854bce259ac12881ad2c32bac421cda34a4adacf6d821b288",
                "User": "6d280bfe1f036ee5cf47b0720764b7ea6f7b35e2",
                "SrcChainId": 7,
                "BlockHeight": 4520478,
                "Time": 1620455100,
                "DstChainId": 6,
                "DstUser": "6d280bfe1f036ee5cf47b0720764b7ea6f7b35e2",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "8971defd59e7d6a854bce259ac12881ad2c32bac421cda34a4adacf6d821b288",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620455100,
                "Height": 4520478,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "8971defd59e7d6a854bce259ac12881ad2c32bac421cda34a4adacf6d821b288",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620455100,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "274.34262",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "d729d9c178d3a92978e9b2abf4a9e0e2bd12804c9b04f8c0dd29cd6d9644b1dc",
                "ChainId": 0,
                "State": 1,
                "Time": 1620455192,
                "Fee": "0",
                "Height": 7630612,
                "SrcChainId": 7,
                "SrcHash": "8971defd59e7d6a854bce259ac12881ad2c32bac421cda34a4adacf6d821b288",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "7554f044fdd45ef4a88893afffa16d075cc3165890ca3d051cd71dbfa4129634",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620455204,
                "Height": 7235024,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "7554f044fdd45ef4a88893afffa16d075cc3165890ca3d051cd71dbfa4129634",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620455204,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "8b8d17ac15785b219cfaf7e05e13467d84b2c9a7240679c30b3166c39c2e32fa",
                "User": "de880bc9984bd36f278da4853534fd26b10384b2",
                "SrcChainId": 7,
                "BlockHeight": 4520477,
                "Time": 1620455097,
                "DstChainId": 6,
                "DstUser": "de880bc9984bd36f278da4853534fd26b10384b2",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "8b8d17ac15785b219cfaf7e05e13467d84b2c9a7240679c30b3166c39c2e32fa",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620455097,
                "Height": 4520477,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "8b8d17ac15785b219cfaf7e05e13467d84b2c9a7240679c30b3166c39c2e32fa",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620455097,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "9.67554514",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "36ef12f31b901bd841d3e152b460a509c87bb532a308dd3024a66f9b084c2c7b",
                "ChainId": 0,
                "State": 1,
                "Time": 1620455192,
                "Fee": "0",
                "Height": 7630612,
                "SrcChainId": 7,
                "SrcHash": "8b8d17ac15785b219cfaf7e05e13467d84b2c9a7240679c30b3166c39c2e32fa",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "1785c9426dbd3d87e23afaed7dea35ae030065d1b4d24e3f0ee1082f18b60fb3",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620455207,
                "Height": 7235025,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "1785c9426dbd3d87e23afaed7dea35ae030065d1b4d24e3f0ee1082f18b60fb3",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620455207,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "166d785eec858ebf965cd9e8bad9c77629b4165155f60e28278d6f706d212924",
                "User": "f561059babbcf7aab6f9a5a6398ae776ec983511",
                "SrcChainId": 7,
                "BlockHeight": 4520375,
                "Time": 1620454791,
                "DstChainId": 6,
                "DstUser": "e6cdee98c5a85ce9bd2a7f824f611514ae7d41d4",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "166d785eec858ebf965cd9e8bad9c77629b4165155f60e28278d6f706d212924",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620454791,
                "Height": 4520375,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "166d785eec858ebf965cd9e8bad9c77629b4165155f60e28278d6f706d212924",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620454791,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "124.7348594562967002",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "a4d228cbe31145c12ba6a068e324134ca8cbc776a07ce53b3a8620d791c96cd6",
                "ChainId": 0,
                "State": 1,
                "Time": 1620454875,
                "Fee": "0",
                "Height": 7630479,
                "SrcChainId": 7,
                "SrcHash": "166d785eec858ebf965cd9e8bad9c77629b4165155f60e28278d6f706d212924",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "cef821e6e9373eea7037b619a1fbd4f52e10897094669b6a3689216af2490e47",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620454889,
                "Height": 7234919,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "cef821e6e9373eea7037b619a1fbd4f52e10897094669b6a3689216af2490e47",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620454889,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "41fde7b38e794940bc9cb2a6e535d44b204a7242bc06ffb2fe7a2e715178c2da",
                "User": "c9593471d89e085e46666bddc006f7541d748950",
                "SrcChainId": 7,
                "BlockHeight": 4520234,
                "Time": 1620454368,
                "DstChainId": 6,
                "DstUser": "616c3fb96affe52d6b2604cfb70abb5f0c742693",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "41fde7b38e794940bc9cb2a6e535d44b204a7242bc06ffb2fe7a2e715178c2da",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620454368,
                "Height": 4520234,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "41fde7b38e794940bc9cb2a6e535d44b204a7242bc06ffb2fe7a2e715178c2da",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620454368,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "1252.4287822699258383",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "556174aa2ed39d4b3f7c3ba7a8d43e7bf492cc6f126e8b4a6527acff819feda0",
                "ChainId": 0,
                "State": 1,
                "Time": 1620454461,
                "Fee": "0",
                "Height": 7630319,
                "SrcChainId": 7,
                "SrcHash": "41fde7b38e794940bc9cb2a6e535d44b204a7242bc06ffb2fe7a2e715178c2da",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "1e0860c4fa7f5880b5fb59043953e93ce6678aa4b2483b4e9360eb3c16bf30b1",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620454493,
                "Height": 7234787,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "1e0860c4fa7f5880b5fb59043953e93ce6678aa4b2483b4e9360eb3c16bf30b1",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620454493,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "4317faf14d464221eeb9d571201be44937bce7cb50d84d746be9d0b1d622f80a",
                "User": "3e09367670ff638db22099fa582f01ea9b6e71d2",
                "SrcChainId": 7,
                "BlockHeight": 4520198,
                "Time": 1620454260,
                "DstChainId": 6,
                "DstUser": "8879011ac32294d458c66d86b837e695a6f96808",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "4317faf14d464221eeb9d571201be44937bce7cb50d84d746be9d0b1d622f80a",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620454260,
                "Height": 4520198,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "4317faf14d464221eeb9d571201be44937bce7cb50d84d746be9d0b1d622f80a",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620454260,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "4688.9557",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "c1b12880fcddc8c782524ec8e71d202affc5ebebb09c61a9ddc5a1febd03caa6",
                "ChainId": 0,
                "State": 1,
                "Time": 1620454354,
                "Fee": "0",
                "Height": 7630261,
                "SrcChainId": 7,
                "SrcHash": "4317faf14d464221eeb9d571201be44937bce7cb50d84d746be9d0b1d622f80a",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "4a4507a2ea5cc166c91ce491a1c79613b03d0897f25f4ec96b7a66e86a52a666",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620454367,
                "Height": 7234745,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "4a4507a2ea5cc166c91ce491a1c79613b03d0897f25f4ec96b7a66e86a52a666",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620454367,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        },
        {
            "WrapperTransaction": {
                "Hash": "3a089f3aefa672f648ee63df945d9b45d660e687e948ab74e1373c328e6ad567",
                "User": "6904f1f14f5fe9e6fa994e380d94a4d591ca2b89",
                "SrcChainId": 7,
                "BlockHeight": 4520192,
                "Time": 1620454242,
                "DstChainId": 6,
                "DstUser": "6904f1f14f5fe9e6fa994e380d94a4d591ca2b89",
                "ServerId": 0,
                "FeeTokenHash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "FeeAmount": "0.32445486",
                "State": 0
            },
            "SrcTransaction": {
                "Hash": "3a089f3aefa672f648ee63df945d9b45d660e687e948ab74e1373c328e6ad567",
                "ChainId": 7,
                "Standard": 0,
                "State": 1,
                "Time": 1620454242,
                "Height": 4520192,
                "DstChainId": 6,
                "SrcTransfer": {
                    "TxHash": "3a089f3aefa672f648ee63df945d9b45d660e687e948ab74e1373c328e6ad567",
                    "ChainId": 7,
                    "Standard": 0,
                    "Time": 1620454242,
                    "Asset": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                    "Amount": "1533.427630404333319",
                    "DstChainId": 6,
                    "DstAsset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668"
                }
            },
            "PolyTransaction": {
                "Hash": "bfe8a6a3fd2c69d9e390d364fe48984e7073924ad3db4e35a6f90bbd32aa8e53",
                "ChainId": 0,
                "State": 1,
                "Time": 1620454347,
                "Fee": "0",
                "Height": 7630256,
                "SrcChainId": 7,
                "SrcHash": "3a089f3aefa672f648ee63df945d9b45d660e687e948ab74e1373c328e6ad567",
                "DstChainId": 6,
                "Key": ""
            },
            "DstTransaction": {
                "Hash": "1cde8c7faa2b0740d44c2326544fa1e7eda88b2a4013c5d1975354f891442638",
                "ChainId": 6,
                "Standard": 0,
                "State": 1,
                "Time": 1620454361,
                "Height": 7234743,
                "SrcChainId": 7,
                "DstTransfer": {
                    "TxHash": "1cde8c7faa2b0740d44c2326544fa1e7eda88b2a4013c5d1975354f891442638",
                    "ChainId": 6,
                    "Standard": 0,
                    "Time": 1620454361,
                    "Asset": "aee4164c1ee46ed0bbc34790f1a3d1fc87796668",
                    "Amount": ""
                }
            },
            "Token": {
                "Hash": "25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
                "ChainId": 7,
                "Name": "MDX",
                "Property": 1,
                "TokenBasicName": "MDX",
                "TokenBasic": {
                    "Name": "MDX",
                    "Precision": 18,
                    "Price": "2.98384891",
                    "Ind": 0,
                    "Time": 1620453742,
                    "Property": 1,
                    "PriceMarkets": null,
                    "Tokens": null
                },
                "TokenMaps": null
            }
        }
    ]
}
```


### POST expecttime

This API returns the expected elapsed time for token to transfer from source chain to target chain.

Request 
```
http://localhost:8080/v1/expecttime/
```

BODY raw
```
{
    "SrcChainId":7,
    "DstChainId":2
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/expecttime/' \
--data-raw '{
                "SrcChainId":7,
                "DstChainId":2
            }'
```

Example Response
```
{
    "SrcChainId": 7,
    "DstChainId": 2,
    "Time": 90
}
```
### POST gettokenasset

This API returns the token balance and total supply according to token or name.

Request
```
http://localhost:8080/v1/gettokenasset/
```

BODY raw
```
{
    "nameOrHash":"HKR"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/gettokenasset/' \
--data-raw '{
    "nameOrHash":"HKR"
}'
```

Example Response
```
[
    {
        "BasicName": "HKR",
        "TokenAsset": [
            {
                "ChainName": "",
                "Hash": "0",
                "TotalSupply": 0,
                "Balance": 0,
                "ErrReason": ""
            },
            {
                "ChainName": "",
                "Hash": "0",
                "TotalSupply": 0,
                "Balance": 0,
                "ErrReason": ""
            }
        ],
        "Precision": 18
    }
]
```

