#!/bin/bash

rm -rf bin/*

for app in "tool,wallet-tool" "wallet" "sign,sign-srv" "encrypt,encrypt-srv"
do
    arr=(`echo $app|tr ',' ' '`)
    dir=${arr[0]}
    if [ ${#arr[@]} == 1 ] 
    then
        appname=${arr[0]}
    else 
        appname=${arr[1]}
    fi
    echo "go build $dir/main.go"
    go build -o bin/$appname $dir/main.go
    chmod +x bin/$appname
done

echo "build finish."
