module poly-bridge

go 1.14

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/goleveldb v1.0.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/devfans/cogroup v1.1.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/hashicorp/golang-lru v0.5.4
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/joeqian10/neo3-gogogo v0.3.8
	github.com/ontio/ontology v1.14.0-beta.0.20210818114002-fedaf66010a7
	github.com/ontio/ontology-crypto v1.1.0
	github.com/ontio/ontology-go-sdk v1.12.3
	github.com/polynetwork/bridge-common v0.0.0-20210730081758-77b97bbeb305
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200827085015-12374709b707
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/polynetwork/poly-io-test v0.0.0-20200819093740-8cf514b07750 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.33.7
	github.com/urfave/cli v1.22.4
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.39.2-0.20200814061308-474a0dbbe4ba
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20210126025041-8d21ec4f0324
)
