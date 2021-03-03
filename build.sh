#!/bin/bash

tag=$1
base=build_${tag}

if [ ! -d "./$base" ]; then
  mkdir -p "./$base"
fi

if [ ! -d "./$base/bridge_tools" ]; then
  mkdir -p "./$base/bridge_tools"
fi

if [ ! -d "./$base/bridge_server" ]; then
  mkdir -p "./$base/bridge_server"
fi

if [ ! -d "./$base/bridge_http" ]; then
  mkdir -p "./$base/bridge_http"
fi

if [ ! -d "./$base/bridge_http/conf" ]; then
  mkdir -p "./$base/bridge_http/conf"
fi

if [ ! -d "./$base/ethereum_listen" ]; then
  mkdir -p "./$base/ethereum_listen"
fi

if [ ! -d "./$base/neo_listen" ]; then
  mkdir -p "./$base/neo_listen"
fi

if [ ! -d "./$base/poly_listen" ]; then
  mkdir -p "./$base/poly_listen"
fi

if [ ! -d "./$base/ontology_listen" ]; then
  mkdir -p "./$base/ontology_listen"
fi

if [ ! -d "./$base/crosschain_effect" ]; then
  mkdir -p "./$base/crosschain_effect"
fi

if [ ! -d "./$base/coinprice_listen" ]; then
  mkdir -p "./$base/coinprice_listen"
fi

if [ ! -d "./$base/chainfee_listen" ]; then
  mkdir -p "./$base/chainfee_listen"
fi

cd ./bridge_tools
go build -tags $tag -o bridge_tools .

mv bridge_tools ./../$base/bridge_tools
if [ "$tag"x = "mainnet"x ]
then
  cp ./conf/config_deploy_mainnet.json ./../$base/bridge_tools
  cp ./conf/config_update_mainnet.json ./../$base/bridge_tools
else
  cp ./conf/config_deploy_testnet.json ./../$base/bridge_tools
  cp ./conf/config_update_testnet.json ./../$base/bridge_tools
fi
cp ./conf/config_transactions.json ./../$base/bridge_tools
cp -rf ./conf/template ./../$base/bridge_tools

cd ./../cmd
go build -tags $tag -o bridge_server main.go

mv bridge_server ./../$base/bridge_server
if [ "$tag"x = "mainnet"x ]
then
  cp ./../conf/config_mainnet.json ./../$base/bridge_server
else
  cp ./../conf/config_testnet.json ./../$base/bridge_server
fi

cd ./../
go build -tags $tag -o bridge_http main.go

mv bridge_http ./$base/bridge_http
cp ./conf/app.conf ./$base/bridge_http/conf

cd ./crosschainlisten

cd ethereumlisten/cmd
go build -tags $tag -o ethereum_listen main.go

mv ethereum_listen ./../../../$base/ethereum_listen
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../../conf/config_mainnet.json ./../../../$base/ethereum_listen
else
  cp ./../../../conf/config_testnet.json ./../../../$base/ethereum_listen
fi

cd ./../../neolisten/cmd
go build -tags $tag -o neo_listen main.go

mv neo_listen ./../../../$base/neo_listen
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../../conf/config_mainnet.json ./../../../$base/neo_listen
else
  cp ./../../../conf/config_testnet.json ./../../../$base/neo_listen
fi

cd ./../../polylisten/cmd
go build -tags $tag -o poly_listen main.go

mv poly_listen ./../../../$base/poly_listen
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../../conf/config_mainnet.json ./../../../$base/poly_listen
else
 cp ./../../../conf/config_testnet.json ./../../../$base/poly_listen
fi

cd ./../../ontologylisten/cmd
go build -tags $tag -o ontology_listen main.go

mv ontology_listen ./../../../$base/ontology_listen
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../../conf/config_mainnet.json ./../../../$base/ontology_listen
else
 cp ./../../../conf/config_testnet.json ./../../../$base/ontology_listen
fi

cd ./../..
cd ./..

cd ./crosschaineffect/cmd
go build -tags $tag -o crosschain_effect main.go

mv crosschain_effect ./../../$base/crosschain_effect
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../conf/config_mainnet.json ./../../$base/crosschain_effect
else
  cp ./../../conf/config_testnet.json ./../../$base/crosschain_effect
fi

cd ./../..

cd ./coinpricelisten/cmd
go build -tags $tag -o coinprice_listen main.go

mv coinprice_listen ./../../$base/coinprice_listen
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../conf/config_mainnet.json ./../../$base/coinprice_listen
else
  cp ./../../conf/config_testnet.json ./../../$base/coinprice_listen
fi

cd ./../..

cd ./chainfeelisten/cmd
go build -tags $tag -o chainfee_listen main.go

mv chainfee_listen ./../../$base/chainfee_listen
if [ "$tag"x = "mainnet"x ]
then
  cp ./../../conf/config_mainnet.json ./../../$base/chainfee_listen
else
  cp ./../../conf/config_testnet.json ./../../$base/chainfee_listen
fi

cd ./../..


