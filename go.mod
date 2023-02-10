module poly-bridge

go 1.14

require (
	github.com/Zilliqa/gozilliqa-sdk v1.2.1-0.20210927032600-4c733f2cb879
	github.com/antihax/optional v1.0.0
	github.com/beego/beego/v2 v2.0.1
	github.com/devfans/cogroup v1.1.0
	github.com/devfans/zion-sdk v0.0.24
	github.com/ethereum/go-ethereum v1.10.17
	github.com/gateio/gateapi-go/v6 v6.23.2
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/joeqian10/neo-gogogo v1.4.0
	github.com/joeqian10/neo3-gogogo v1.2.1
	github.com/novifinancial/serde-reflection/serde-generate/runtime/golang v0.0.0-20211013011333-6820d5b97d8c
	github.com/ontio/ontology v1.11.1-0.20200812075204-26cf1fa5dd47
	github.com/ontio/ontology-crypto v1.2.1
	github.com/ontio/ontology-go-sdk v1.11.4
	github.com/polynetwork/bridge-common v0.0.0-20230112035625-bfadb7cf548b
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/ripple-sdk v0.0.0-20220616022641-d64d4aa053fe
	github.com/portto/aptos-go-sdk v0.0.0-20221025115549-5c74acafa193
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rubblelabs/ripple v0.0.0-20220222071018-38c1a8b14c18
	github.com/starcoinorg/starcoin-go v0.0.0-20220105024102-530daedc128b
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.35.9
	github.com/urfave/cli v1.22.4
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
//github.com/cosmos/iavl v0.16.0
)

replace (
	//github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.46.8
	//github.com/gogo/protobuf v1.3.3 => github.com/cosmos/gogoproto v1.4.4
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20210126025041-8d21ec4f0324
	github.com/polynetwork/kai-relayer => github.com/dogecoindev/kai-relayer v0.0.0-20210609112229-34bf794e78e7
	//github.com/rubblelabs/ripple v0.0.0-20220222071018-38c1a8b14c18 => github.com/siovanus/ripple v0.0.0-20220406100637-81f6afe283d9
	github.com/polynetwork/ripple-sdk => github.com/siovanus/ripple-sdk v0.0.0-20230113103625-9795771fe0fc
	github.com/rubblelabs/ripple => github.com/siovanus/ripple v0.0.0-20230113075118-4a31480c1af2

)
