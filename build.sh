#!/bin/bash

tag=$1

if [ ! -d "./build" ]; then
  mkdir -p "./build"
fi

cd ./tools
go build -tags $tag -o bridge_tools .

if [ ! -d "./build/bridge_tools" ]; then
  mkdir -p "./build/bridge_tools"
fi
mv bridge_tools ./../build/bridge_tools
cp conf/conf_depoly_mainnet.json ./../build/bridge_tools
cp conf/conf_depoly_testnet.json ./../build/bridge_tools
cp conf/conf_dump.json ./../build/bridge_tools

cd ./../cmd
go build -tags $tag -o bridge_server main.go

if [ ! -d "./build/bridge_server" ]; then
  mkdir -p "./build/bridge_server"
fi
mv bridge_server ./../build/bridge_server
cp ./../conf/config_mainnet.json ./../build/bridge_server
cp ./../conf/config_testnet.json ./../build/bridge_server

cd ./../
go build -tags $tag -o bridge_http main.go

if [ ! -d "./build/bridge_http" ]; then
  mkdir -p "./build/bridge_http"
fi
mv bridge_http ./build/bridge_http
if [ ! -d "./build/bridge_http/conf" ]; then
  mkdir -p "./build/bridge_http/conf"
fi
cp ./conf/app.conf ./build/bridge_http/conf

cd ./crosschainlisten

cd ethereumlisten/cmd
go build -tags $tag -o ethereum_listen main.go

if [ ! -d "./build/ethereum_listen" ]; then
  mkdir -p "./build/ethereum_listen"
fi
mv ethereum_listen ./../../../build/ethereum_listen
cp ./../../../conf/config_mainnet.json ./../../../build/ethereum_listen
cp ./../../../conf/config_testnet.json ./../../../build/ethereum_listen

cd ./../../neolisten/cmd
go build -tags $tag -o neo_listen main.go
if [ ! -d "./build/neo_listen" ]; then
  mkdir -p "./build/neo_listen"
fi
mv neo_listen ./../../../build/neo_listen
cp ./../../../conf/config_mainnet.json ./../../../build/neo_listen
cp ./../../../conf/config_testnet.json ./../../../build/neo_listen

cd ./../../polylisten/cmd
go build -tags $tag -o poly_listen main.go
if [ ! -d "./build/poly_listen" ]; then
  mkdir -p "./build/poly_listen"
fi
mv poly_listen ./../../../build/poly_listen
cp ./../../../conf/config_mainnet.json ./../../../build/poly_listen
cp ./../../../conf/config_testnet.json ./../../../build/poly_listen

cd ./../..
cd ./..

cd ./crosschaineffect/cmd
go build -tags $tag -o crosschain_effect main.go
if [ ! -d "./build/crosschain_effect" ]; then
  mkdir -p "./build/crosschain_effect"
fi
mv crosschain_effect ./../../build/crosschain_effect
cp ./../../conf/config_mainnet.json ./../../build/crosschain_effect
cp ./../../conf/config_testnet.json ./../../build/crosschain_effect

cd ./../..

cd ./coinpricelisten/cmd
go build -tags $tag -o coinprice_listen main.go
if [ ! -d "./build/coinprice_listen" ]; then
  mkdir -p "./build/coinprice_listen"
fi
mv coinprice_listen ./../../build/coinprice_listen
cp ./../../conf/config_mainnet.json ./../../build/coinprice_listen
cp ./../../conf/config_testnet.json ./../../build/coinprice_listen


cd ./../..

cd ./chainfeelisten/cmd
go build -tags $tag -o chainfee_listen main.go
if [ ! -d "./build/chainfee_listen" ]; then
  mkdir -p "./build/chainfee_listen"
fi
mv chainfee_listen ./../../build/chainfee_listen
cp ./../../conf/config_mainnet.json ./../../build/chainfee_listen
cp ./../../conf/config_testnet.json ./../../build/chainfee_listen

cd ./../..


