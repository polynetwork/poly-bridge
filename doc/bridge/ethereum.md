## Ethereum(ERC20 and ERC721) Chain Listener


With the specified wrapper contract addresses, the ethereum chain listener will monitor the matched emitted contract events to find the desired `PolyWrapperLock` event and `PolyWrapperSpeedUp` event. Currently, two contract addresses were specified for wrapper smart contract with the second will use the cross chain asset as fee instead. The `PolyWrapperSpeedUp ` event will replace the transaction fee with a newer one normally to speed up the process. Besides the contract for asset of type ERC20, the NFT contract was deployed too and to be monitored in a similar way as the ERC20. 

#### Wrapper Transaction

Field|Description
:--:|:--:
Hash        | Txid[2:]
User        | event.Sender[2:]
DstChainId  | event.ToChainId
DstUser     | event.ToAddress
FeeTokenHash| event.FromAsset (Zero for the new version)
FeeAmount   | event.Fee
ServerId    | event.Id
BlockHeight | block.Height


#### CCM transaction

In a similar way, the proxy lock/unlock event will be collected while listening, for both the ERC20 and NFT contract. In the process, the eccm locked/unlocked events will collected first if matching with the specified the ECCM contract address.

Source Transaction

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | event.TxHash[2:]
State       | 1
Fee         | execution.GasConsumed
Time        | block.Time
Height      | block.Height
User        | event.Sender[2:]
DstChainId  | event.ToChainId
Contract    | event.ProxyOrAssetContract[2:]
Key         | event.TxId
Param       | event.RawData
SrcTransfer | srcTransfer
Standard    | TokenTypeErc20 or TokenTypeErc721

Source Transfer (Proxy lock event)

Field|Description
:--:|:--:
ChainId     | CurrentChainId 
TxHash      | event.TxHash[2:]
Time        | block.Time
From        | event.Sender[2:]
To          | event.ProxyOrAssetContract[2:]
Asset       | eventSrc.FromAssetHash[2:]
Amount      | eventSrc.Amount (eventSrc.TokenId for NFT)
DstChainId  | eventSrc.ToChainId
DstUser     | eventSrc.ToAddress
DstAsset    | eventSrc.ToAssetHash

Destination Transaction

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | event.TxHash[2:]
State       | 1
Fee         | execution.GasConsumed
Time        | block.Time
Height      | block.Height
SrcChainId  | event.FromChainId
Contract    | event.ToContract
PolyHash    | event.CrossChainTxHash
DstTransfer | dstTransfer
Standard    | TokenTypeErc20 or TokenTypeErc721

Destination Transfer (Proxy unlock event)

Field|Description
:--:|:--:
ChainId | CurrentChainId
TxHash  | event.TxHash[2:]
Time    | block.Time
From    | event.ToContract
To      | eventDst.ToAddress[2:]
Asset   | eventDst.ToAssetHash[2:]
Amount  | eventDst.Amount(eventDst.TokenId for NFT)


























