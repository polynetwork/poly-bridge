# Deploy And Update

## build

进入到项目目录。

编译测试网版本：
```
./build.sh testnet
```

生成build_testnet为测试网执行文件以及配置。

编译主网版本：
```
./build.sh mainnet
```

生成build_mainnet为主网执行文件以及配置。

## deploy

部署测试网：
```
cd build_testnet
cd bridge_tools
./bridge_tools --cliconfig config_deploy_testnet.json --cmd 1
```

部署主网网：
```
cd build_mainnet
cd bridge_tools
./bridge_tools --cliconfig config_deploy_mainnet.json --cmd 1
```

## 运行

运行测试网：
```
cd build_testnet
cd bridge_server
./bridge_server --cliconfig config_testnet.json
cd ./../bridge_http
./bridge_http
```

运行主网：
```
cd build_mainnet
cd bridge_server
./bridge_server --cliconfig config_mainnet.json
cd ./../bridge_http
./bridge_http
```

## 升级

升级测试网

配置升级文件 ./build_testnet/bridge_tools/config_update_testnet.json
```
cd build_testnet
cd bridge_tools
./bridge_tools --cliconfig config_update_testnet.json --cmd 4
```

更新配置文件 ./build_testnet/bridge_server/config_testnet.json

重启bridge_server。

升级主网

配置升级文件 ./build_mainnet/bridge_tools/config_update_mainnet.json
```
cd build_mainnet
cd bridge_tools
./bridge_tools --cliconfig config_update_mainnet.json --cmd 4
```

更新配置文件 ./build_mainnet/bridge_server/config_mainnet.json

重启bridge_server。


