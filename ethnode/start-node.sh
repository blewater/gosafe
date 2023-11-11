#!/usr/bin/env bash

docker run -it -p 8545:8545 -p 8546:8546 -p 30303:30303 -p 30303:30303/udp --name dev-geth -v ~/.ethereum:/root/.ethereum ethereum/client-go:stable --datadir /root/.ethereum --dev --http --http.addr 0.0.0.0 --http.api eth,web3,net --http.corsdomain "http://remix.ethereum.org"