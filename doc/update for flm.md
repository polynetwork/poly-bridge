# 升级支持FLM和swth代币流程

## 升级主网

### build

进入到项目目录。

编译主网版本：
```
./build.sh mainnet
```

生成build_mainnet为主网执行文件以及配置。

### update

配置升级文件 [config_update_add_token_flm_mainnet.json](https://github.com/polynetwork/poly-bridge/blob/master/bridge_tools/conf/template/config_update_add_token_flm_mainnet.json)

```
cd build_mainnet
cd bridge_tools
./bridge_tools --cliconfig config_update_add_token_flm_mainnet.json --cmd 4
```

更新配置文件 [config_mainnet.json](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_mainnet.json)

重启bridge_server。

```
cd build_mainnet
cd bridge_server
``` 

重启bridge_http。
```
cd build_mainnet
cd bridge_http
``` 

增加告警：log的error为"is not listening"