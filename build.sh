#!/bin/bash

tag=$1

cd ./tools
go build -tags $tag -o bridge_tools .
mv bridge_tools ./../build

cd ./../cmd
go build -tags $tag -o bridge_server main.go
mv bridge_server ./../build

cd ./../
go build -tags $tag -o bridge_http main.go
mv bridge_http ./build

cd ./crosschainlisten

cd ethereumlisten/cmd
go build -tags $tag -o ethereum_listen main.go
mv ethereum_listen ./../../../build

cd ./../../neolisten/cmd
go build -tags $tag -o neo_listen main.go
mv neo_listen ./../../../build

cd ./../../polylisten/cmd
go build -tags $tag -o poly_listen main.go
mv poly_listen ./../../../build

cd ./../..
cd ./..

cd ./crosschaineffect/cmd
go build -tags $tag -o crosschain_effect main.go
mv crosschain_effect ./../../build

cd ./../..

cd ./coinpricelisten/cmd
go build -tags $tag -o coinprice_listen main.go
mv coinprice_listen ./../../build

cd ./../..

cd ./chainfeelisten/cmd
go build -tags $tag -o chainfee_listen main.go
mv chainfee_listen ./../../build

cd ./../..


