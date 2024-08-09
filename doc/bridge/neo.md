## NEO Chain Listener


Neo chain listener will go through the transactions included in the new block, and only check against the transactions with type `InvocationTransaction`. When such a transaction was found, the transaction execution notification will be checked . 

#### Wrapper Transaction
The transaction with execution notification matching the defined wrapper contract and with a valid parsed method(`PolyWrapperLock`) will be treated as a wrapper transaction.

Field|Description
:--:|:--:
Hash        | Txid[2:]
User        | notify.State.Value[2].Value
DstChainId  | notify.State.Value[3].Value
DstUser     | notify.State.Value[4].Value
FeeTokenHash| notify.State.Value[1].Value
FeeAmount   | notify.State.Value[6].Value
ServerId    | notify.State.Value[7].Value
Status      | STATE_SOURCE_DONE
Time        | block.Time
BlockHeight | block.Height
SrcChainId  | CurrentChainId


#### CCM transaction

If matching the CCM contract, and the parsed method is `CrossChainLockEvent`, the transaction will be treated as a source transaction.

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | Txid[2:]
State       | 1
Fee         | execution.GasConsumed
Time        | block.Time
Height      | block.Height
User        | srcTransfer.From
DstChainId  | notify.State.Value[3].Value
Contract    | notify.State.Value[2].Value
Key         | notify.State.Value[4].Value
Param       | notify.State.Value[5].Value
SrcTransfer | srcTransfer

Method `Lock` and `LockEvent` indicate it's a Source Transfer.

Field|Description
:--:|:--:
ChainId     | CurrentChainId 
TxHash      | Txid[2:]
Time        | block.Time
From        | notifySrc.State.Value[2].Value
To          | notify.State.Value[2].Value
Asset       | notifySrc.State.Value[1].Value
Amount      | notifySrc.State.Value[6].Value
DstChainId  | notifySrc.State.Value[3].Value
DstUser     | notifySrc.State.Value[5].Value
DstAsset    | notifySrc.State.Value[4].Value


If matching the CCM contract, and the parsed method is `CrossChainUnLockEvent`, the transaction will be treated as a destination transaction.

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | Txid[2:]
State       | 1
Fee         | execution.GasConsumed
Time        | block.Time
Height      | block.Height
SrcChainId  | notify.State.Value[1].Value
Contract    | notify.State.Value[2].Value
PolyHash    | notify.State.Value[3].Value
DstTransfer | dstTransfer



Method `Unlock` and `UnlockEvent` indicate it's a Destination Transfer.

Field|Description
:--:|:--:
ChainId | CurrentChainId
TxHash  | TxId[2:]
Time    | block.Time
From    | notify.State.Value[2].Value
To      | notifyDst.State.Value[2].Value
Asset   | notifyDst.State.Value[1].Value
Amount  | notifyDst.State.Value[3].Value


























