# Ontology支持升级

## 升级测试网

### build

进入到项目目录。

编译测试网版本：
```
./build.sh testnet
```

生成build_testnet为测试网执行文件以及配置。

### update

配置升级文件 [config_update_testenet.json](https://github.com/polynetwork/poly-bridge/blob/master/bridge_tools/conf/config_update_testnet.json)

```
cd build_testnet
cd bridge_tools
./bridge_tools --cliconfig config_update_mainnet.json --cmd 6
./bridge_tools --cliconfig config_update_testnet.json --cmd 4
```

更新配置文件 [config_testnet](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_testnet.json)

```
cd build_testnet
cd bridge_server
vi ./config_testnet.json
```
重启bridge_server。

## 升级主网

### build

进入到项目目录。

编译主网版本：
```
./build.sh mainnet
```

生成build_mainnet为主网执行文件以及配置。

### update

配置升级文件 [config_update_mainnet.json](https://github.com/polynetwork/poly-bridge/blob/master/bridge_tools/conf/config_update_mainnet.json)

```
cd build_mainnet
cd bridge_tools
./bridge_tools --cliconfig config_update_mainnet.json --cmd 6
./bridge_tools --cliconfig config_update_mainnet.json --cmd 4
```

更新配置文件 [config_mainnet](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_mainnet.json)

```
cd build_mainnet
cd bridge_server
vi ./config_mainnet.json
``` 

重启bridge_server。