## Poly Network Listener

Poly network listener monitor the poly chain blocks to track the state update of cross chain transactions on the middle chain.

The smart contract events emitted will be checked one by one. When the contract address matches the defined poly contract address, and the parsed method is either `makeProof` or `btcTxToRelay`, the execution notification will be parsed and recorded.

Field|Description
:--:|:--:
ChainId     | CurrentChainId
Hash        | event.TxHash
State       | event.State
Fee         | 0
Time        | block.Time
Height      | block.Height
SrcChainId  | notify.States[1]
DstChainId  | notify.States[2]
SrcHash     | notify.States[3]

