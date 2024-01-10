## Status Check

`crosschaineffect` is used to do periodical checks on the processing of cross chain transactions. Some logging will be triggered at bad situations and be monitored by the logging center to trigger alarms in IM channels.


#### Update Poly transaction source Hash

One step is status check is to find the poly transactions with fake source hash, which is emitted from the transaction events before sumbitting to polynetowrk, it will be updated with the real source transaction hash by matching the fake source hash id with the source transaction `key` field.

#### Transaction status check

Unfinished cross chain transactions will be regularly checked.

#### Update cross chain transaction status

Based on the collected data from chains, the cross chain transaction status will be updated. Normally, the confirmations of wrapper transactions and destination transactions(on target chain)  will be verified.

#### Update average cross chain transction time expense 

The average time expense of cross chain transactions will be updated.

#### Chain listener status check

Chain status collected from chains will be used to verify if the listeners are still running.



