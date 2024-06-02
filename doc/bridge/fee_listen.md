## Fee Listener

Chain suggested fee is monitored for several chains, with the value updated into the database to be used in bridge page for fee suggestions.

#### Ethereum/HECO/BSC/OK Fee

Via sdk, the suggested gas price is retrieved, with defined gas limit, the max fee is got. Based on the max fee, the min fee and the proxy fee is calculated.

#### NEO/Ontology Fee

Rather than fetching from the chain, the initial gas price is hard coded as 1.