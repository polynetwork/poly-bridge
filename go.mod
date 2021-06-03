module poly-bridge

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/goleveldb v1.0.0
	github.com/ethereum/go-ethereum v1.9.15
	github.com/hashicorp/golang-lru v0.5.4
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/joeqian10/neo3-gogogo v0.3.4
	github.com/ontio/ontology v1.11.1-0.20200812075204-26cf1fa5dd47
	github.com/ontio/ontology-crypto v1.0.9
	github.com/ontio/ontology-go-sdk v1.11.4
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200827085015-12374709b707
	github.com/polynetwork/eth-contracts v0.0.0-20200814062128-70f58e22b014
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/polynetwork/poly-io-test v0.0.0-20200819093740-8cf514b07750 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.4
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.39.2-0.20200814061308-474a0dbbe4ba
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20210126025041-8d21ec4f0324
)
