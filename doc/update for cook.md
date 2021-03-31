# 升级支持cook代币流程

## 升级主网

### build

进入到项目目录。

编译主网版本：
```
./build.sh mainnet
```

生成build_mainnet为主网执行文件以及配置。

### update

配置升级文件 [config_update_token_cook_mainnet.json](https://github.com/polynetwork/poly-bridge/blob/master/bridge_tools/conf/template/config_update_token_cook_mainnet.json)

```
cd build_mainnet
cd bridge_tools
./bridge_tools --cliconfig config_update_add_token_share_mainnet.json --cmd 4
```

更新配置文件 [config_mainnet.json](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_mainnet.json)

```
cd build_mainnet
cd bridge_server
vi ./config_mainnet.json
``` 

重启bridge_server。

不需要更新bridge_server的备份和bridge_http