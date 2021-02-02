#!/bin/bash

tag=$1

cd ./tools
go build -tags $tag -o bridge_tools .
cp bridge_tools ./../

cd ./../cmd
go build -tags $tag -o bridge_server main.go
cp bridge_server ./../

cd ./../
go build -tags $tag -o bridge_http main.go

cd ./crosschainlisten

cd ethereumlisten/cmd
go build -tags $tag -o ethereum_listen main.go
cp ethereum_listen ./../../../

cd ./../../neolisten/cmd
go build -tags $tag -o neo_listen main.go
cp neo_listen ./../../../

cd ./../../polylisten/cmd
go build -tags $tag -o poly_listen main.go
cp poly_listen ./../../../

cd ./../..
cd ./..

cd ./crosschaineffect/cmd
go build -tags $tag -o crosschain_effect main.go
cp crosschain_effect ./../../

cd ./../..

cd ./coinpricelisten/cmd
go build -tags $tag -o coinprice_listen main.go
cp coinprice_listen ./../../

cd ./../..

cd ./chainfeelisten/cmd
go build -tags $tag -o chainfee_listen main.go
cp chainfee_listen ./../../

cd ./../..


