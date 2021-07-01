
# cross chain explorer

## 安装环境
- GoLang版本：go1.13.3及以上
- 数据库：Mysql5.6及以上

## 编译
1. 下载代码
```
git clone https://github.com/polynetwork/explorer.git
```
2. 进入项目目录
```
cd cmd
go build main.go
```

## 数据库初始化

在数据库中导入以下sql
```
CREATE SCHEMA IF NOT EXISTS `cross_chain_explorer` DEFAULT CHARACTER SET utf8;
USE `cross_chain_explorer`;

DROP TABLE IF EXISTS `chain_info`;
CREATE TABLE `chain_info` (
 `xname` VARCHAR(32) NOT NULL COMMENT '链名称',
 `id`  INT(4) NOT NULL COMMENT '链id',
 `xtype` INT(4) NOT NULL COMMENT '链类型',
 `height` INT(12) NOT NULL COMMENT '解析的区块高度',
 `txin` 	INT(12) NOT NULL COMMENT '链的入金数量',
 `txout`	INT(12) NOT NULL COMMENT '链的出金数量',
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `chain_contract`;
CREATE TABLE `chain_contract` (
  `id` INT(4) NOT NULL COMMENT '链id',
  `contract` VARCHAR(128) NOT NULL COMMENT '跨链合约地址'
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `chain_token`;
CREATE TABLE `chain_token` (
  `id` INT(4) NOT NULL COMMENT '链id',
  `xtoken` VARCHAR(32) NOT NULL COMMENT '跨链通用token名称',
  `hash` VARCHAR(128) NOT NULL COMMENT 'token地址',
  `xname` VARCHAR(32) NOT NULL COMMENT 'token名称',
  `xtype` VARCHAR(32) NOT NULL COMMENT 'token类型',
  `xprecision` VARCHAR(32)  NOT NULL COMMENT 'token精度',
  `xdesc` VARCHAR(1024) COMMENT 'token描述'
) ENGINE=INNODB DEFAULT CHARSET=utf8;

INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("poly",0,0,22732,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("btc",1,1,0,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("eth",2,2,10650091,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("ontology",3,3,9300490,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("neo",4,4,6023777,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("switcheo",5,5,202650,0,0);

INSERT INTO `chain_contract`(`id`,`contract`) VALUES(0, "0300000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(2, "838bf9e95cb12dd76a54c9f9d2e3082eaf928270");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(3, "0900000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(4, "82a3401fb9a60db42c6fa2ea2b6d62e872d6257f");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology Gas", "0200000000000000000000000000000000000000", "ong", "OEP4", "1000000000","ong");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology", "0100000000000000000000000000000000000000", "ont", "OEP4", "1","ont");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ethereum", "0000000000000000000000000000000000000000", "ether", "ether", "1000000000000000000","ether");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "DeepBrain Chain", "b951ecbbc5fe37a9c280a76cb0ce0014827294cf", "DeepBrain Coin", "NEP5", "100000000","DBC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "DeepBrain Chain", "64626331", "DeepBrain Coin", "Cosmos", "100000000","dbc1");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Switcheo", "ab38352559b8b203bde5fddfa0b07d8b2525e132", "Switcheo", "NEP5", "100000000","SWTH");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Switcheo", "73777468", "Switcheo", "Cosmos", "100000000","swth");

DROP TABLE IF EXISTS `poly_validators`;
CREATE TABLE `poly_validators` (
  `height` INT(12) NOT NULL COMMENT '交易的高度',
  `validators`  VARCHAR(8192) COMMENT '验证节点',
  PRIMARY KEY (`height`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `mchain_tx`;
CREATE TABLE `mchain_tx` (
 `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
 `chain_id` INT(4) NOT NULL COMMENT '链ID',
 `state` INT(1) NOT NULL COMMENT '交易状态',
 `tt` INT(11) NOT NULL COMMENT '交易时间',
 `fee` BIGINT(8) NOT NULL COMMENT '交易手续费',
 `height` INT(12) NOT NULL COMMENT '交易的高度',
 `fchain` INT(4) NOT NULL COMMENT '源链的id',
 `ftxhash` VARCHAR(128) NOT NULL COMMENT '源链的交易hash',
 `tchain` INT(4) NOT NULL COMMENT '目标链的id',
 `xkey` VARCHAR(8192) COMMENT '比特币交易',
 PRIMARY KEY (`txhash`),
 UNIQUE (`ftxhash`),
 INDEX (`tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `fchain_tx`;
CREATE TABLE `fchain_tx` (
 `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
 `chain_id` INT(4) NOT NULL COMMENT '链ID',
 `state` INT(1) NOT NULL COMMENT '交易状态',
 `tt` INT(4) NOT NULL COMMENT '交易时间',
 `fee` BIGINT(8) NOT NULL COMMENT '交易手续费',
 `height` INT(12) NOT NULL COMMENT '交易的高度',
 `xuser` VARCHAR(128) NOT NULL COMMENT '用户',
 `tchain` INT(4) NOT NULL COMMENT '目标链的id',
 `contract` VARCHAR(128) NOT NULL COMMENT '执行的合约',
 `xkey` VARCHAR(8192) NOT NULL COMMENT '目标链的参数',
 `xparam` VARCHAR(8192) NOT NULL COMMENT '合约参数',
 PRIMARY KEY (`txhash`),
 INDEX (`tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `fchain_transfer`;
CREATE TABLE `fchain_transfer` (
  `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
  `chain_id` INT(4) NOT NULL COMMENT '链ID',
  `tt` INT(4) NOT NULL COMMENT '交易时间',
  `asset` VARCHAR(128) NOT NULL COMMENT '资产hash',
  `xfrom` VARCHAR(128) NOT NULL COMMENT '发送用户',
  `xto` VARCHAR(128) NOT NULL COMMENT '接受用户',
  `amount` BIGINT(8) NOT NULL COMMENT '收到的金额',
  `tochainid` INT(4) NOT NULL COMMENT '目标链的id',
  `toasset` VARCHAR(1024) NOT NULL COMMENT '目标链的资产hash',
  `touser` VARCHAR(128) NOT NULL COMMENT '目标链的接受用户',
  PRIMARY KEY (`txhash`),
  INDEX (`asset`, `tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tchain_tx`;
CREATE TABLE `tchain_tx` (
 `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
 `chain_id` INT(4) NOT NULL COMMENT '链ID',
 `state` INT(1) NOT NULL COMMENT '交易状态',
 `tt` INT(4) NOT NULL COMMENT '交易时间',
 `fee` BIGINT(8) NOT NULL COMMENT '交易手续费',
 `height` INT(12) NOT NULL COMMENT '交易的高度',
 `fchain` INT(4) NOT NULL COMMENT '源链的id',
 `contract` VARCHAR(128) NOT NULL COMMENT '执行的合约',
 `rtxhash` VARCHAR(128) NOT NULL COMMENT '中继链的交易hash',
 PRIMARY KEY (`txhash`),
 UNIQUE (`rtxhash`),
 INDEX (`tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tchain_transfer`;
CREATE TABLE `tchain_transfer` (
  `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
  `chain_id` INT(4) NOT NULL COMMENT '链ID',
  `tt` INT(4) NOT NULL COMMENT '交易时间',
  `asset` VARCHAR(128) NOT NULL COMMENT '资产hash',
  `xfrom` VARCHAR(128) NOT NULL COMMENT '发送用户',
  `xto` VARCHAR(128) NOT NULL COMMENT '接受用户',
  `amount` BIGINT(8) NOT NULL COMMENT '收到的金额',
  PRIMARY KEY (`txhash`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `asset_statistic`;
CREATE TABLE `asset_statistic` (
  `xname` VARCHAR(16)  COMMENT '资产名称',
  `addressnum`   INT(4) NOT NULL COMMENT '资产的总地址数',
  `amount`       BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `amount_btc`  BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `amount_usd`  BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `txnum`       INT(4) NOT NULL COMMENT '总的交易个数',
  `latestupdate` INT(4)  NOT NULL COMMENT '统计数据的时间点',
  PRIMARY KEY (`xname`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));
```

## 配置
以下为跨链浏览器的所有配置信息，主要包括redis，mysql，server，ontology，btc，eth等服务运行需要的配置信息，在运行项目前应根据需要修改相关配置。
 ```json
{
  "redis": {
    "addr": "localhost:6379",
    "proto": "tcp",
    "pool_size": 50,
    "min_idle_conns": 10,
    "dial_timeout": 2,
    "read_timeout": 2,
    "write_timeout": 2,
    "idle_timeout": 10,
    "expiration": 60
  },
  "mysql": {
    "url": "127.0.0.1:3306",
    "user": "root",
    "pwd": "root",
    "dbName": "cross_chain_explorer"
  },
  "server": {
    "rest_port": 30334,
    "version": "1.0.0",
    "http_max_connections": 10000,
    "master": 1,
    "http_cert_path": "",
    "http_key_path": "http_key_path",
    "statistic_time_slot": 10,
    "loglevel": 1
  },
  "coinmarketcap": {
    "url": "https://pro-api.coinmarketcap.com/v1/cryptocurrency/",
    "appkey": "8c175886-3aec-4276-8961-1f0ce06ab69e"
  },
  "neo": {
    "name": "neo",
    "chainId": 4,
    "rawurl": ["http://seed10.ngd.network:11332","http://seed9.ngd.network:11332"],
    "block_duration": 1
  },
  "ontology": {
    "name": "ontology",
    "chainId": 3,
    "rawurl": ["http://dappnode4.ont.io:20336","http://dappnode2.ont.io:20336"],
    "block_duration": 1
  },
  "ethereum": {
    "name": "eth",
    "chainId": 2,
    "rawurl": ["http://onto-eth.ont.io:10331"],
    "block_duration": 1,
    "proxy": "a4a0919ba5a2a89cb2eb1cb4018a16841b2e943f",
    "btcx": "92705a16815a3d1aec3ce9cc273c5aa302961fcc"
  },
  "alliance": {
    "name": "poly",
    "chainId": 0,
    "rawurl": ["http://13.92.155.62:20336"],
    "block_duration": 1
  },
  "btc": {
    "name": "btc",
    "chainId": 1,
    "user": ["omnicorerpc"],
    "passwd": ["EzriglUqnFC!"],
    "rawurl": ["http://18.140.187.37:18332"],
    "block_duration": 1
  },
  "cosmos": {
    "name": "switcheo",
    "chainId": 5,
    "rawurl": ["http://175.41.151.35:26657", "http://54.255.5.46:26657"],
    "block_duration": 1
  }
}
```
以上配置大部分不需要重新配置，但需要根据实际情况配置mysql和主从。

主从配置：

+ master配置为1则为主
+ master配置为0则为从


## 链ID

链名称|链ID
:--:|:--:
poly|0
btc|1
eth|2
ontology|3
neo|4
cosmos|5

## 测试网节点

40.115.153.174:30334

## API

### 1. getexplorerinfo
查询信息

POST
```
http://{{host}}/api/v1/getexplorerinfo
```


#### 参数:
start: 开始的日期
end: 结束日期
开始日期和结束日期用于返回链上该日期区间上的跨链交易总数。

```json
{
    "start":1592274867,
    "end":1593534067
}
```

#### example:

```
curl -i http://172.168.3.26:30334/api/v1/getexplorerinfo -X POST -d "{"start":"1592274867","end":"1593534067"}"
```

```json
{
    "action": "getexplorerinfo",
    "code": 1,
    "desc": "success",
    "result": "{\"chains\":[{\"chainid\":0,\"chainname\":\"poly\",\"blockheight\":646454,\"in\":6,\"incrosschaintxstatus\":null,\"out\":6,\"outcrosschaintxstatus\":null,\"addresses\":0,\"contracts\":[{\"chainid\":0,\"contract\":\"0300000000000000000000000000000000000000\"}],\"tokens\":[]},{\"chainid\":1,\"chainname\":\"btc\",\"blockheight\":632686,\"in\":1,\"incrosschaintxstatus\":[{\"timestamp\":1592265600,\"txnumber\":0},{\"timestamp\":1592352000,\"txnumber\":0},{\"timestamp\":1592438400,\"txnumber\":0},{\"timestamp\":1592524800,\"txnumber\":0},{\"timestamp\":1592611200,\"txnumber\":0},{\"timestamp\":1592697600,\"txnumber\":0},{\"timestamp\":1592784000,\"txnumber\":0},{\"timestamp\":1592870400,\"txnumber\":0},{\"timestamp\":1592956800,\"txnumber\":0},{\"timestamp\":1593043200,\"txnumber\":0},{\"timestamp\":1593129600,\"txnumber\":0},{\"timestamp\":1593216000,\"txnumber\":0},{\"timestamp\":1593302400,\"txnumber\":0},{\"timestamp\":1593388800,\"txnumber\":0},{\"timestamp\":1593475200,\"txnumber\":0},{\"timestamp\":1593561600,\"txnumber\":0}],\"out\":1,\"outcrosschaintxstatus\":[{\"timestamp\":1592265600,\"txnumber\":0},{\"timestamp\":1592352000,\"txnumber\":0},{\"timestamp\":1592438400,\"txnumber\":0},{\"timestamp\":1592524800,\"txnumber\":0},{\"timestamp\":1592611200,\"txnumber\":0},{\"timestamp\":1592697600,\"txnumber\":0},{\"timestamp\":1592784000,\"txnumber\":0},{\"timestamp\":1592870400,\"txnumber\":0},{\"timestamp\":1592956800,\"txnumber\":0},{\"timestamp\":1593043200,\"txnumber\":0},{\"timestamp\":1593129600,\"txnumber\":0},{\"timestamp\":1593216000,\"txnumber\":0},{\"timestamp\":1593302400,\"txnumber\":0},{\"timestamp\":1593388800,\"txnumber\":0},{\"timestamp\":1593475200,\"txnumber\":0},{\"timestamp\":1593561600,\"txnumber\":0}],\"addresses\":2,\"contracts\":[],\"tokens\":[{\"chainid\":1,\"hash\":\"0000000000000000000000000000000000000011\",\"name\":\"btc\",\"type\":\"BTC\",\"precision\":100000000,\"desc\":\"btc\"}]},{\"chainid\":2,\"chainname\":\"eth\",\"blockheight\":8205950,\"in\":0,\"incrosschaintxstatus\":null,\"out\":3,\"outcrosschaintxstatus\":[{\"timestamp\":1592265600,\"txnumber\":0},{\"timestamp\":1592352000,\"txnumber\":0},{\"timestamp\":1592438400,\"txnumber\":0},{\"timestamp\":1592524800,\"txnumber\":0},{\"timestamp\":1592611200,\"txnumber\":0},{\"timestamp\":1592697600,\"txnumber\":0},{\"timestamp\":1592784000,\"txnumber\":0},{\"timestamp\":1592870400,\"txnumber\":0},{\"timestamp\":1592956800,\"txnumber\":0},{\"timestamp\":1593043200,\"txnumber\":0},{\"timestamp\":1593129600,\"txnumber\":0},{\"timestamp\":1593216000,\"txnumber\":0},{\"timestamp\":1593302400,\"txnumber\":0},{\"timestamp\":1593388800,\"txnumber\":0},{\"timestamp\":1593475200,\"txnumber\":0},{\"timestamp\":1593561600,\"txnumber\":0}],\"addresses\":1,\"contracts\":[{\"chainid\":2,\"contract\":\"ba6f835ecae18f5fc5ebc074e5a0b94422a13126\"}],\"tokens\":[{\"chainid\":2,\"hash\":\"bbe0da0f3d5132a5c245d7760d2700e2192fba39\",\"name\":\"btc\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"btc\"},{\"chainid\":2,\"hash\":\"63692d2ba64a5869114068b7b08dffed94f378d8\",\"name\":\"oep4\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"oep4\"},{\"chainid\":2,\"hash\":\"a8177ee8a6e496c701cfec0cbd8f723cc851153d\",\"name\":\"ong\",\"type\":\"ERC20\",\"precision\":1000000000,\"desc\":\"ong\"},{\"chainid\":2,\"hash\":\"514092ef689ebae8eebbca97fd6987e94b033ccb\",\"name\":\"ont\",\"type\":\"ERC20\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":2,\"hash\":\"0000000000000000000000000000000000000000\",\"name\":\"ether\",\"type\":\"ether\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":2,\"hash\":\"d1cb2bda2146c0878b41b5c0164e4420aef72584\",\"name\":\"erc20\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"erc20\"},{\"chainid\":2,\"hash\":\"20f307ea523e69d195b3a370fe6496eb50ce281a\",\"name\":\"neo\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"neo\"}]},{\"chainid\":3,\"chainname\":\"ontology\",\"blockheight\":13087900,\"in\":0,\"incrosschaintxstatus\":null,\"out\":0,\"outcrosschaintxstatus\":null,\"addresses\":0,\"contracts\":[{\"chainid\":3,\"contract\":\"0900000000000000000000000000000000000000\"}],\"tokens\":[{\"chainid\":3,\"hash\":\"b7f398711664de1dd685d9ba3eee3b6b830a7d83\",\"name\":\"btc\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"btc\"},{\"chainid\":3,\"hash\":\"99981b7485df558eb63f45ee19dcb0458b83ed25\",\"name\":\"oep4\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"oep4\"},{\"chainid\":3,\"hash\":\"0000000000000000000000000000000000000002\",\"name\":\"ong\",\"type\":\"OEP4\",\"precision\":1000000000,\"desc\":\"ong\"},{\"chainid\":3,\"hash\":\"0000000000000000000000000000000000000001\",\"name\":\"ont\",\"type\":\"OEP4\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":3,\"hash\":\"08014516ad7cbaecd4f488f80772e41d1611e179\",\"name\":\"ether\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":3,\"hash\":\"7e0c97ff0879b17ef09ef77c91056d81f923e135\",\"name\":\"erc20\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"erc20\"}]},{\"chainid\":4,\"chainname\":\"neo\",\"blockheight\":4486917,\"in\":4,\"incrosschaintxstatus\":[{\"timestamp\":1592265600,\"txnumber\":0},{\"timestamp\":1592352000,\"txnumber\":0},{\"timestamp\":1592438400,\"txnumber\":0},{\"timestamp\":1592524800,\"txnumber\":0},{\"timestamp\":1592611200,\"txnumber\":0},{\"timestamp\":1592697600,\"txnumber\":0},{\"timestamp\":1592784000,\"txnumber\":0},{\"timestamp\":1592870400,\"txnumber\":0},{\"timestamp\":1592956800,\"txnumber\":0},{\"timestamp\":1593043200,\"txnumber\":0},{\"timestamp\":1593129600,\"txnumber\":0},{\"timestamp\":1593216000,\"txnumber\":0},{\"timestamp\":1593302400,\"txnumber\":0},{\"timestamp\":1593388800,\"txnumber\":0},{\"timestamp\":1593475200,\"txnumber\":0},{\"timestamp\":1593561600,\"txnumber\":0}],\"out\":0,\"outcrosschaintxstatus\":null,\"addresses\":1,\"contracts\":[{\"chainid\":4,\"contract\":\"978286951e0011221de3fffe6a9e6dd160925837\"}],\"tokens\":[{\"chainid\":4,\"hash\":\"a63d7dffa7718902fda0f64e57f3c5e0c33fd3ff\",\"name\":\"ont\",\"type\":\"NEP5\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":4,\"hash\":\"74fac41ad5ad23921a3400e953e1cafb41240d08\",\"name\":\"ether\",\"type\":\"NEP5\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":4,\"hash\":\"c074a05e9dcf0141cbe6b4b3475dd67baf4dcb60\",\"name\":\"neo\",\"type\":\"NEP5\",\"precision\":100000000,\"desc\":\"neo\"},{\"chainid\":4,\"hash\":\"74f2dc36a68fdc4682034178eb2220729231db76\",\"name\":\"neogas\",\"type\":\"NEP5\",\"precision\":100000000,\"desc\":\"neogas\"}]},{\"chainid\":5,\"chainname\":\"cosmos\",\"blockheight\":137572,\"in\":1,\"incrosschaintxstatus\":[{\"timestamp\":1592265600,\"txnumber\":0},{\"timestamp\":1592352000,\"txnumber\":0},{\"timestamp\":1592438400,\"txnumber\":0},{\"timestamp\":1592524800,\"txnumber\":0},{\"timestamp\":1592611200,\"txnumber\":0},{\"timestamp\":1592697600,\"txnumber\":0},{\"timestamp\":1592784000,\"txnumber\":0},{\"timestamp\":1592870400,\"txnumber\":0},{\"timestamp\":1592956800,\"txnumber\":0},{\"timestamp\":1593043200,\"txnumber\":0},{\"timestamp\":1593129600,\"txnumber\":0},{\"timestamp\":1593216000,\"txnumber\":0},{\"timestamp\":1593302400,\"txnumber\":0},{\"timestamp\":1593388800,\"txnumber\":0},{\"timestamp\":1593475200,\"txnumber\":0},{\"timestamp\":1593561600,\"txnumber\":0},{\"timestamp\":1593648000,\"txnumber\":0}],\"out\":1,\"outcrosschaintxstatus\":[{\"timestamp\":1592265600,\"txnumber\":0},{\"timestamp\":1592352000,\"txnumber\":0},{\"timestamp\":1592438400,\"txnumber\":0},{\"timestamp\":1592524800,\"txnumber\":0},{\"timestamp\":1592611200,\"txnumber\":0},{\"timestamp\":1592697600,\"txnumber\":0},{\"timestamp\":1592784000,\"txnumber\":0},{\"timestamp\":1592870400,\"txnumber\":0},{\"timestamp\":1592956800,\"txnumber\":0},{\"timestamp\":1593043200,\"txnumber\":0},{\"timestamp\":1593129600,\"txnumber\":0},{\"timestamp\":1593216000,\"txnumber\":0},{\"timestamp\":1593302400,\"txnumber\":0},{\"timestamp\":1593388800,\"txnumber\":0},{\"timestamp\":1593475200,\"txnumber\":0},{\"timestamp\":1593561600,\"txnumber\":0},{\"timestamp\":1593648000,\"txnumber\":0}],\"addresses\":1,\"contracts\":[],\"tokens\":[{\"chainid\":5,\"hash\":\"62746378\",\"name\":\"btc\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"btc\"},{\"chainid\":5,\"hash\":\"6f65703478\",\"name\":\"oep4\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"oep4\"},{\"chainid\":5,\"hash\":\"6f6e6778\",\"name\":\"ong\",\"type\":\"Cosmos\",\"precision\":1000000000,\"desc\":\"ong\"},{\"chainid\":5,\"hash\":\"6f6e7478\",\"name\":\"ont\",\"type\":\"Cosmos\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":5,\"hash\":\"65746878\",\"name\":\"ether\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":5,\"hash\":\"657263323078\",\"name\":\"erc20\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"erc20\"},{\"chainid\":5,\"hash\":\"6e656f78\",\"name\":\"neo\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"neo\"},{\"chainid\":5,\"hash\":\"67617378\",\"name\":\"neogas\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"neogas\"},{\"chainid\":5,\"hash\":\"7374616b65\",\"name\":\"atom\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"atom\"}]}],\"crosstxnumber\":6,\"tokens\":[{\"name\":\"btc\",\"tokens\":[{\"chainid\":1,\"hash\":\"0000000000000000000000000000000000000011\",\"name\":\"btc\",\"type\":\"BTC\",\"precision\":100000000,\"desc\":\"btc\"},{\"chainid\":2,\"hash\":\"bbe0da0f3d5132a5c245d7760d2700e2192fba39\",\"name\":\"btc\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"btc\"},{\"chainid\":3,\"hash\":\"b7f398711664de1dd685d9ba3eee3b6b830a7d83\",\"name\":\"btc\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"btc\"},{\"chainid\":5,\"hash\":\"62746378\",\"name\":\"btc\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"btc\"}]},{\"name\":\"oep4\",\"tokens\":[{\"chainid\":2,\"hash\":\"63692d2ba64a5869114068b7b08dffed94f378d8\",\"name\":\"oep4\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"oep4\"},{\"chainid\":3,\"hash\":\"99981b7485df558eb63f45ee19dcb0458b83ed25\",\"name\":\"oep4\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"oep4\"},{\"chainid\":5,\"hash\":\"6f65703478\",\"name\":\"oep4\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"oep4\"}]},{\"name\":\"ong\",\"tokens\":[{\"chainid\":2,\"hash\":\"a8177ee8a6e496c701cfec0cbd8f723cc851153d\",\"name\":\"ong\",\"type\":\"ERC20\",\"precision\":1000000000,\"desc\":\"ong\"},{\"chainid\":3,\"hash\":\"0000000000000000000000000000000000000002\",\"name\":\"ong\",\"type\":\"OEP4\",\"precision\":1000000000,\"desc\":\"ong\"},{\"chainid\":5,\"hash\":\"6f6e6778\",\"name\":\"ong\",\"type\":\"Cosmos\",\"precision\":1000000000,\"desc\":\"ong\"}]},{\"name\":\"ont\",\"tokens\":[{\"chainid\":2,\"hash\":\"514092ef689ebae8eebbca97fd6987e94b033ccb\",\"name\":\"ont\",\"type\":\"ERC20\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":3,\"hash\":\"0000000000000000000000000000000000000001\",\"name\":\"ont\",\"type\":\"OEP4\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":4,\"hash\":\"a63d7dffa7718902fda0f64e57f3c5e0c33fd3ff\",\"name\":\"ont\",\"type\":\"NEP5\",\"precision\":1,\"desc\":\"ont\"},{\"chainid\":5,\"hash\":\"6f6e7478\",\"name\":\"ont\",\"type\":\"Cosmos\",\"precision\":1,\"desc\":\"ont\"}]},{\"name\":\"ether\",\"tokens\":[{\"chainid\":2,\"hash\":\"0000000000000000000000000000000000000000\",\"name\":\"ether\",\"type\":\"ether\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":3,\"hash\":\"08014516ad7cbaecd4f488f80772e41d1611e179\",\"name\":\"ether\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":4,\"hash\":\"74fac41ad5ad23921a3400e953e1cafb41240d08\",\"name\":\"ether\",\"type\":\"NEP5\",\"precision\":100000000,\"desc\":\"ether\"},{\"chainid\":5,\"hash\":\"65746878\",\"name\":\"ether\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"ether\"}]},{\"name\":\"erc20\",\"tokens\":[{\"chainid\":2,\"hash\":\"d1cb2bda2146c0878b41b5c0164e4420aef72584\",\"name\":\"erc20\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"erc20\"},{\"chainid\":3,\"hash\":\"7e0c97ff0879b17ef09ef77c91056d81f923e135\",\"name\":\"erc20\",\"type\":\"OEP4\",\"precision\":100000000,\"desc\":\"erc20\"},{\"chainid\":5,\"hash\":\"657263323078\",\"name\":\"erc20\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"erc20\"}]},{\"name\":\"neo\",\"tokens\":[{\"chainid\":2,\"hash\":\"20f307ea523e69d195b3a370fe6496eb50ce281a\",\"name\":\"neo\",\"type\":\"ERC20\",\"precision\":100000000,\"desc\":\"neo\"},{\"chainid\":4,\"hash\":\"c074a05e9dcf0141cbe6b4b3475dd67baf4dcb60\",\"name\":\"neo\",\"type\":\"NEP5\",\"precision\":100000000,\"desc\":\"neo\"},{\"chainid\":5,\"hash\":\"6e656f78\",\"name\":\"neo\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"neo\"}]},{\"name\":\"neogas\",\"tokens\":[{\"chainid\":4,\"hash\":\"74f2dc36a68fdc4682034178eb2220729231db76\",\"name\":\"neogas\",\"type\":\"NEP5\",\"precision\":100000000,\"desc\":\"neogas\"},{\"chainid\":5,\"hash\":\"67617378\",\"name\":\"neogas\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"neogas\"}]},{\"name\":\"atom\",\"tokens\":[{\"chainid\":5,\"hash\":\"7374616b65\",\"name\":\"atom\",\"type\":\"Cosmos\",\"precision\":100000000,\"desc\":\"atom\"}]}]}",
    "version": "1.0.0"
}
```

### 2. getcrosstxlist
查询跨链交易列表

POST
```
http://{{host}}/api/v1/getcrosstxlist
```

#### 参数:
start,end: 指定返回交易索引

```json
{
    "pageNo":1,
    "pageSize":10
}
```

#### example:

```
curl -i http://172.168.3.26:30334/api/v1/getcrosstxlist -X POST -d "{"start":"0","end":"5"}"
```

```json
{
    "action": "getcrosstxlist",
    "code": 1,
    "desc": "success",
    "result": "[{\"txhash\":\"ef727cc5a28cfeb35e4a9afddf0efcce1a77f66378f48266d011164712aec621\",\"state\":1,\"timestamp\":1593603465,\"fee\":0,\"blockheight\":624365},{\"txhash\":\"30186df5508fa795ec14c5a37783fbde87f47f907968739958df10114ea9d905\",\"state\":1,\"timestamp\":1593601774,\"fee\":0,\"blockheight\":624083},{\"txhash\":\"52a0380eaff506b02a76c740b77ea0b2aa44b8b27dee718aaeafdb22d347e71a\",\"state\":1,\"timestamp\":1593599958,\"fee\":0,\"blockheight\":623808},{\"txhash\":\"a5103df0ae2fad743d185edc77d700349d18abd7717696cb199a047c25cffc9f\",\"state\":1,\"timestamp\":1593598287,\"fee\":0,\"blockheight\":623589},{\"txhash\":\"e0a992365c6dcbac683ae3b051b6796b7e3c72a70dc9bb9164b81d2989bf8ded\",\"state\":1,\"timestamp\":1593592776,\"fee\":0,\"blockheight\":622686}]",
    "version": "1.0.0"
}
```

state的解释：
1. 已经完成，成功
0. 正在进行中
3. 失败

### 3. getcrosstx
查询跨链交易详细信息

GET
```
http://{{host}}/api/v1/getcrosstx/:txhash
```

#### 参数:
```
txhash : "980107a8ca1c2db41497391cc3487c0a4898de442036d24ecdb36553bef74ba4"
```

#### example:

```
curl -i http://172.168.3.26:30334/api/v1/getcrosstx/00020fb0b090681648b50734dc835b26b0552adfa3186259089ed3e1ac0e7af1
```

```json
{
    "action": "getcrosstx",
    "code": 1,
    "desc": "success",
    "result": "{\"fchaintr\":{\"chainid\":2,\"chainname\":\"ontology\",\"txhash\":\"980107a8ca1c2db41497391cc3487c0a4898de442036d24ecdb36553bef74ba4\",\"state\":1,\"tt\":1566463480,\"fee\":10000,\"height\":31608,\"tchainid\":3,\"tchainname\":\"neo\",\"key\":\"\",\"contract\":\"0200000000000000000000000000000000000000\",\"value\":\"AJLHY6wbJqE2j1VoRNPp3ZdacpLE58qNNy 123\",\"type\":0,\"typename\":\"unkown\",\"transfer\":{\"from\":\"AJLHY6wbJqE2j1VoRNPp3ZdacpLE58qNNy\",\"to\":\"AJLHY6wbJqE2j1VoRNPp3ZdacpLE58qNNy\",\"token\":\"0200000000000000000000000000000000000000\",\"amount\":123}},\"fchaintr_valid\":true,\"mchaintx\":{\"chainid\":1,\"chainname\":\"alliance\",\"txhash\":\"6c19f93157bc5f8ed7850b669f66c4457b3256fd1849b2be0a9f8da9aba86101\",\"state\":1,\"tt\":1566463472,\"fee\":10000,\"height\":11908,\"fchainid\":2,\"fchainname\":\"ontology\",\"ftxhash\":\"980107a8ca1c2db41497391cc3487c0a4898de442036d24ecdb36553bef74ba4\",\"tchainid\":3,\"tchainname\":\"neo\",\"key\":\"xx\"},\"mchaintx_valid\":true,\"tchaintx\":{\"chainid\":3,\"chainname\":\"neo\",\"txhash\":\"00020fb0b090681648b50734dc835b26b0552adfa3186259089ed3e1ac0e7af1\",\"state\":1,\"tt\":1566463488,\"fee\":10000,\"height\":311608,\"fchainid\":2,\"fchainname\":\"ontology\",\"mtxhash\":\"6c19f93157bc5f8ed7850b669f66c4457b3256fd1849b2be0a9f8da9aba86101\"},\"tchaintx_valid\":true}",
    "version": "1.0.0"
}
```


### 4. gettokentxlist
查询一个币种上的跨链交易列表

POST
```
http://{{host}}/api/v1/gettokentxlist
```

#### 参数:
```json
{
    "token":"0000000000000000000000000000000000000000",
    "pageNo":1,
    "pageSize":10
}
```

#### example:

```
curl -i http://172.168.3.26:30334/api/v1/gettokentxlist -X POST -d "{"token":"0000000000000000000000000000000000000000"}"
```

```json
{
    "action": "gettokentxlist",
    "code": 1,
    "desc": "success",
    "result": "[{\"txhash\":\"ec378a2ef0aa62f2451e6ad3591997984d8a0a44e0cfe93087917430c07f3d60\",\"from\":\"344cfc3b8635f72f14200aaf2168d9f75df86fd3\",\"to\":\"75ed27ee68f0d6bdd4e41e38388c5a9028fb6707\",\"amount\":\"100000\",\"timestamp\":1593592641,\"blockheight\":8205354,\"direct\":1},{\"txhash\":\"ec19bf1277f2d110d675f50a212f23ef75421db8efc5bf8d034b172f289e6062\",\"from\":\"344cfc3b8635f72f14200aaf2168d9f75df86fd3\",\"to\":\"75ed27ee68f0d6bdd4e41e38388c5a9028fb6707\",\"amount\":\"12345678912345\",\"timestamp\":1593598059,\"blockheight\":8205820,\"direct\":1},{\"txhash\":\"2b0ea80bcbf0aad255e7f1977e7f3df14d17929ad8ac772cd38a86269f712949\",\"from\":\"344cfc3b8635f72f14200aaf2168d9f75df86fd3\",\"to\":\"75ed27ee68f0d6bdd4e41e38388c5a9028fb6707\",\"amount\":\"112345678900\",\"timestamp\":1593599648,\"blockheight\":8205924,\"direct\":1}]",
    "version": "1.0.0"
}
```

direct的解释：
1. 从该链到其他链，outgo
2. 从其他链到本链，income

### 5. getaddresstxlist
查询一个地址上的跨链交易列表

POST
```
http://{{host}}/api/v1/getaddresstxlist
```

#### 参数:
```json
{
    "address":"344cfc3b8635f72f14200aaf2168d9f75df86fd3",
    "chain":2,
    "pageNo":1,
    "pageSize":10
}
```

#### example:

```
curl -i http://172.168.3.26:30334/api/v1/getaddresstxlist -X POST -d "{"address":"344cfc3b8635f72f14200aaf2168d9f75df86fd3", "chain":"2"}"
```

```json
{
    "action": "getaddresstxlist",
    "code": 1,
    "desc": "success",
    "result": "[{\"txhash\":\"ec378a2ef0aa62f2451e6ad3591997984d8a0a44e0cfe93087917430c07f3d60\",\"from\":\"344cfc3b8635f72f14200aaf2168d9f75df86fd3\",\"to\":\"75ed27ee68f0d6bdd4e41e38388c5a9028fb6707\",\"amount\":\"100000\",\"timestamp\":1593592641,\"blockheight\":8205354,\"direct\":1},{\"txhash\":\"ec19bf1277f2d110d675f50a212f23ef75421db8efc5bf8d034b172f289e6062\",\"from\":\"344cfc3b8635f72f14200aaf2168d9f75df86fd3\",\"to\":\"75ed27ee68f0d6bdd4e41e38388c5a9028fb6707\",\"amount\":\"12345678912345\",\"timestamp\":1593598059,\"blockheight\":8205820,\"direct\":1},{\"txhash\":\"2b0ea80bcbf0aad255e7f1977e7f3df14d17929ad8ac772cd38a86269f712949\",\"from\":\"344cfc3b8635f72f14200aaf2168d9f75df86fd3\",\"to\":\"75ed27ee68f0d6bdd4e41e38388c5a9028fb6707\",\"amount\":\"112345678900\",\"timestamp\":1593599648,\"blockheight\":8205924,\"direct\":1}]",
    "version": "1.0.0"
}
```

direct的解释：
1. 从该链到其他链，outgo
2. 从其他链到本链，income

### 6. getassetstatistic
查询跨链资产的统计信息

GET
```
http://{{host}}/api/v1/getassetstatistic
```

#### 参数:
```
无
```

#### example:

```
curl -i http://172.168.3.26:30334/api/v1/getassetstatistic
```

```json
{"action":"getassetstatistic","code":1,"desc":"success","result":"{\"amount_btc_total\":\"2302\",\"amount_usd_total\":\"26929296\",\"asset_statistics\":[{\"name\":\"Switcheo\",\"addressnumber\":2749,\"addressnumber_precent\":\"99.93%\",\"amount\":\"472418432.99673509\",\"amount_btc\":\"2302.43175628\",\"amount_btc_precent\":\"100.00%\",\"amount_usd\":\"26929296.96655629\",\"Amount_usd_precent\":\"100.00%\",\"txnumber\":3466,\"txnumber_precent\":\"99.94%\",\"latestupdate\":1598846400},{\"name\":\"DeepBrain Chain\",\"addressnumber\":2,\"addressnumber_precent\":\"0.07%\",\"amount\":\"0.00002\",\"amount_btc\":\"0\",\"amount_btc_precent\":\"0.00%\",\"amount_usd\":\"0.00000001\",\"Amount_usd_precent\":\"0.00%\",\"txnumber\":2,\"txnumber_precent\":\"0.06%\",\"latestupdate\":1598846400},{\"name\":\"Ethereum\",\"addressnumber\":0,\"addressnumber_precent\":\"0.00%\",\"amount\":\"0\",\"amount_btc\":\"0\",\"amount_btc_precent\":\"0.00%\",\"amount_usd\":\"0\",\"Amount_usd_precent\":\"0.00%\",\"txnumber\":0,\"txnumber_precent\":\"0.00%\",\"latestupdate\":0},{\"name\":\"Ontology\",\"addressnumber\":0,\"addressnumber_precent\":\"0.00%\",\"amount\":\"0\",\"amount_btc\":\"0\",\"amount_btc_precent\":\"0.00%\",\"amount_usd\":\"0\",\"Amount_usd_precent\":\"0.00%\",\"txnumber\":0,\"txnumber_precent\":\"0.00%\",\"latestupdate\":0},{\"name\":\"Ontology Gas\",\"addressnumber\":0,\"addressnumber_precent\":\"0.00%\",\"amount\":\"0\",\"amount_btc\":\"0\",\"amount_btc_precent\":\"0.00%\",\"amount_usd\":\"0\",\"Amount_usd_precent\":\"0.00%\",\"txnumber\":0,\"txnumber_precent\":\"0.00%\",\"latestupdate\":0}]}","version":"1.0.0"}
```

### 7. gettransferstatistic

GET
```
http://{{host}}/api/v1/gettransferstatistic/:chain
```

#### 参数:
```
chain : 1
```

## 使用API

+ 在chrome浏览器中下载插件swagger ui console
+ 访问服务http://40.115.153.174:30335/swagger/swagger.json开始使用跨链浏览器API


