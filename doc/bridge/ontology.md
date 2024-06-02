## Ontology Chain Listener


Ontology chain listener will go through the event notifications emitted in the new block. The specified wrapper contract address and CCM contract address will be used to detect the contract events. The parsed method in the event notification will be used to identify the cross chain transaction type which is one of wrapper transaction, source transaction and destination transaction.

#### Wrapper Transaction
The transaction with execution notification matching the defined wrapper contract and with a valid parsed method(`PolyWrapperLock`) will be treated as a wrapper transaction.

Field|Description
:--:|:--:
Hash        | event.TxHash
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

If matching the CCM contract, and the parsed method is `makeFromOntProof`, the transaction will be treated as a source transaction.

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | event.TxHash
State       | event.State
Fee         | event.GasConsumed
Time        | block.Time
Height      | block.Height
User        | srcTransfer.From
DstChainId  | notify.State.Value[2].Value
Contract    | notify.State.Value[5].Value
Key         | notify.State.Value[4].Value
Param       | notify.State.Value[6].Value
SrcTransfer | srcTransfer

Method `lock` indicates it's a Source Transfer.

Field|Description
:--:|:--:
ChainId     | CurrentChainId 
TxHash      | event.TxHash
Time        | block.Time
From        | notifySrc.State.Value[2].Value
To          | notify.State.Value[5].Value
Asset       | notifySrc.State.Value[1].Value
Amount      | notifySrc.State.Value[6].Value
DstChainId  | notifySrc.State.Value[3].Value
DstUser     | notifySrc.State.Value[5].Value
DstAsset    | notifySrc.State.Value[4].Value


If matching the CCM contract, and the parsed method is `verifyToOntProof`, the transaction will be treated as a destination transaction.

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | event.TxHash
State       | event.State
Fee         | event.GasConsumed
Time        | block.Time
Height      | block.Height
SrcChainId  | notify.State.Value[3].Value
Contract    | notify.State.Value[5].Value
PolyHash    | notify.State.Value[1].Value
DstTransfer | dstTransfer



Method `unlock` indicates it's a Destination Transfer.

Field|Description
:--:|:--:
ChainId | CurrentChainId
TxHash  | event.TxHash
Time    | block.Time
From    | notify.State.Value[5].Value
To      | notifyDst.State.Value[2].Value
Asset   | notifyDst.State.Value[1].Value
Amount  | notifyDst.State.Value[3].Value


























