# 升级支持原生代币作为手续费

## 升级测试网

### build

进入到项目目录。

编译测试网版本：
```
./build.sh testnet
```

生成build_testnet为测试网执行文件以及配置。

### update

更新配置文件 [config_testnet](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_testnet.json)

```
cd build_testnet
cd bridge_server
vi ./config_testnet.json
```
重启bridge_server。
重启bridge_http。

## 升级主网主机

### build

进入到项目目录。

编译主网版本：
```
./build.sh mainnet
```

生成build_mainnet为主网执行文件以及配置。

### update

更新配置文件 [config_mainnet.json](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_mainnet.json)

```
cd build_mainnet
cd bridge_server
vi ./config_mainnet.json
``` 

重启bridge_server。
重启bridge_http。

## 升级主网备机

### build

进入到项目目录。

编译主网版本：
```
./build.sh mainnet
```

生成build_mainnet为主网执行文件以及配置。

### update

更新配置文件 [config_mainnet_backup.json](https://github.com/polynetwork/poly-bridge/blob/master/conf/config_mainnet_backup.json)

```
cd build_mainnet
cd bridge_server
vi ./config_mainnet_backup.json
``` 

重启bridge_server。
重启bridge_http。


