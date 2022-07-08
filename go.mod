module poly-bridge

go 1.14

require (
	github.com/Zilliqa/gozilliqa-sdk v1.2.1-0.20210927032600-4c733f2cb879
	github.com/antihax/optional v1.0.0
	github.com/beego/beego/v2 v2.0.1
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/goleveldb v1.0.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/devfans/cogroup v1.1.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gateio/gateapi-go/v6 v6.23.2
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/hashicorp/golang-lru v0.5.4
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/joeqian10/neo3-gogogo v1.1.2
	github.com/novifinancial/serde-reflection/serde-generate/runtime/golang v0.0.0-20211013011333-6820d5b97d8c
	github.com/ontio/ontology v1.14.0-beta.0.20210818114002-fedaf66010a7
	github.com/ontio/ontology-crypto v1.2.1
	github.com/ontio/ontology-go-sdk v1.12.4
	github.com/polynetwork/bridge-common v0.0.0-20210730081758-77b97bbeb305
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200827085015-12374709b707
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/polynetwork/ripple-sdk v0.0.0-20220616022641-d64d4aa053fe
	github.com/rubblelabs/ripple v0.0.0-20220222071018-38c1a8b14c18
	github.com/starcoinorg/starcoin-go v0.0.0-20220105024102-530daedc128b
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.33.7
	github.com/urfave/cli v1.22.4
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.39.2-0.20200814061308-474a0dbbe4ba
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20210126025041-8d21ec4f0324
	github.com/rubblelabs/ripple v0.0.0-20220222071018-38c1a8b14c18 => github.com/siovanus/ripple v0.0.0-20220406100637-81f6afe283d9
)
