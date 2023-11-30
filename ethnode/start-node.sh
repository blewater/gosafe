#!/usr/bin/env bash

docker run -it -p 8545:8545 -p 8546:8546 -p 30303:30303 -p 30303:30303/udp --name dev-geth -v ~/.ethereum:/root/.ethereum ethereum/client-go:stable \
  --http.vhosts '*,localhost,host.docker.internal' \
  --http \
  --http.api eth,net,web3,debug \
  --http.corsdomain '*' \
  --http.addr "0.0.0.0" \
  --nodiscover --maxpeers 0 --mine \
  --networkid 1337 \
  --dev \
  --allow-insecure-unlock \
  --rpc.allow-unprotected-txs \
  --miner.gaslimit 12000000
