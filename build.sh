#!/bin/bash

tag=$1
base=build_${tag}

if [ ! -d "./$base" ]; then
  mkdir -p "./$base"
fi

go build -tags $tag -o bridge_server ./cmd
go build -tags $tag -o bridge_http ./
go build -tags $tag -o tools ./bridge_tools


if [ ! -d "./$base/bridge_server" ]; then
  mkdir -p "./$base/bridge_server"
fi

if [ ! -d "./$base/bridge_http" ]; then
  mkdir -p "./$base/bridge_http"
fi

if [ ! -d "./$base/bridge_tools" ]; then
  mkdir -p "./$base/bridge_tools"
fi


mv bridge_server ./$base/bridge_server/bridge_server
mv bridge_http ./$base/bridge_http/bridge_http
mv tools ./$base/bridge_tools/bridge_tools

if [ "$tag"x = "mainnet"x ]
then
  cp ./conf/config_mainnet.json ./$base/bridge_server/
  cp ./conf/config_mainnet.json ./$base/bridge_http/
else
  cp ./conf/config_testnet.json ./$base/bridge_server/
  cp ./conf/config_testnet.json ./$base/bridge_http/
fi

