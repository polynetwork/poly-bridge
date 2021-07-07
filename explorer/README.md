
# cross chain explorer


## API

### 1. getexplorerinfo
查询信息

Get
```
http://{{host}}/api/v1/getexplorerinfo
```


#### 参数:
无

```json

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
    "chain": 2,
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


