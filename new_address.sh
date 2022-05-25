#!/bin/bash

## create free address script
while true
do
    COINLIST="btc bch bsv dash doge qtum ltc eth trx sol fil ada xrp near dot avax"
    for coin in $COINLIST
    do
        echo "starting create $coin address..."
        ./bin/wallet -coin $coin -num 2000 -loop
    done

    sleep 10
done