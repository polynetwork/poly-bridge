### Poly Bridge Overview


<div align=center><img width="800" height="570" src="doc/bridge/PolyBridge.jpg"/></div>
Poly Bridge


Poly bridge monitors the target chains, parse the blocks, and find the cross chain transactions, then write into the state database(mysql), with which the bridge server will be able to expose APIs to check cross chain transactions status.


For each target chain, the listener process, at a specified interval, check the latest block height from time to time. If a new block is found and added on the chain, the listener will verify if the block get enough confirmations, then parse the block to find the wrapper transactions(ERC20, NFT) and ECCM event(Lock/Unlock Event), then write into the database.


Data parsed from the chain blocks include:
##### Wrapper Transaction

Wrapper transactions (ERC20, NFT, etc), parsed from the blocks will be inserted or updated into the database table with hash the primary key. 

##### Proxy Lock Events

Proxy lock events indicating a cross chain source transaction has be processed by the source chain relayer, will be inserted or updated into the database table.

##### Proxy Unlock Events

Proxy unlock events indicating a cross chain target transaction has be processed by the target chain relayer, will be inserted or updated into the database table.

##### Poly Transaction

Poly Transactions happen on the poly chain. Its status will be updated into the database as well to track the cross chain transaction status. 

## HTTP API

### Transaction Status

Status Code|Description
:--:|:--:
0|finished
1|pendding
2|source done
3|source confirmed
4|poly confirmed

### Server Address
* [testnet](https://bridge.poly.network/testnet/v1/)
* [mainnet](https://bridge.poly.network/v1/)

### Endpoints
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
* [POST transactionofcurve](#post-transactionofcurve)
* [POST transactionsofunfinished](#post-transactionsofunfinished)
* [POST transactionsofasset](#post-transactionsofasset)
* [POST expecttime](#post-expecttime)


<div align=center><img width="800" height="570" src="doc/bridge/PolyCrossChain.jpg"/></div>
Poly Cross Chain


