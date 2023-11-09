#!/usr/bin/env bash

docker run --name dev-geth -it -p 8545:8545 -p 8546:8546 -p 30303:30303 -p 30303:30303/udp -v ~/.ethereum:/root/.ethereum ethereum/client-go:stable --datadir /root/.ethereum --dev --http --http.api eth,web3,net --http.corsdomain "http://remix.ethereum.org"