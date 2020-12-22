module poly-swap

go 1.14

require (
    github.com/astaxie/beego v1.12.1
	github.com/ethereum/go-ethereum v1.9.15
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/urfave/cli v1.22.1
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace (
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20200824102609-fddf06a45f66
)
