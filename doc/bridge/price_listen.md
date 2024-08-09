## Price Listener
Coin prices are monitored to and updated into the databases. These values will be used to estimate the transaction fee when target chain fee coin differs from the initial fee coin, especially during submitting cross chain transaction on to the target chain.

#### Coin Markets
The prices source markets includes `coinmarketcap`, `binance` and `self(hosted by ourselves)`. Some coins have multiple prices sources, the average price will be used at this condition.

