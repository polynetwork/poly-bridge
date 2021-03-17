# explorer migrate

## 测试网迁移

目前服务所在的机器：A1-JP-CROSSCHAIN-PROD-NODE8
服务目录：/data1/cross_chain_explorer
数据库：127.0.0.1：3306 cross_chain_explorer

+ 将数据库中数据迁移到新的数据库cross_chain_explorer
+ 将目录下的bsc、cmd、config、crosschaineffect、heco拷贝到新的服务器
+ 更新bsc、heco、crosschaineffect以及config下的配置，确保6380、30334这两个配置的端口没有被占用
+ 启动cmd/crosschain-explorer、bsc/bsc_chain_explorer、heco/heco_chain_explorer、crosschaineffect/effect_explorer

## 主网迁移

目前服务所在的机器：A1-CrossChain-EXPLORER-TEST-NODE1
程序目录：crosschain_explorer为主、crosschain_explorer_r1和crosschain_explorer_r2为从
数据库：127.0.0.1:3306 cross_chain_explorer4

1. 迁移主服务器

+ 将数据库中数据迁移到新的数据库cross_chain_explorer
+ 将目录下的bsc、cmd、config、crosschain_effect、heco拷贝到新的服务器
+ 更新bsc、heco、crosschain_effect以及config下的配置，确保数据库访问正确，确保6380、30324（可以改为30334）这两个配置的端口没有被占用
+ 启动cmd/crosschain-explorer、bsc/bsc_chain_explorer、heco/heco_chain_explorer、crosschain_effect/effect_explorer

2. 迁移从服务器

+ 将目录下的cmd、config拷贝到新的服务器
+ 更新config下的配置，确保数据库访问正确，确保6380、30335这两个配置的端口没有被占用
+ 启动cmd/crosschain-explorer-r1


