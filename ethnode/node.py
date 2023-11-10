#!/usr/bin/env python3
from web3 import Web3, HTTPProvider

# Initialize Web3 with the Geth node HTTP provider
w3 = Web3(HTTPProvider('http://127.0.0.1:8545'))

# Ensure connection is established
if not w3.is_connected():
    print("Failed to connect to Geth node")
    exit()

# Set default account
w3.eth.defaultAccount = w3.eth.accounts[0]

# Print default account
print('Default account:', w3.eth.defaultAccount)

# Get the balance of the default account
balance = w3.eth.get_balance(w3.eth.defaultAccount)

# Get the balance of the default account in ether
balance_ether = w3.from_wei(balance, 'ether')
print('Balance in ether:', balance_ether)

# Reading the encrypted key from the keystore file
keystore_path = '/Users/mar/.ethereum/keystore/UTC--2023-11-11T19-57-37.248094917Z' \
                '--fdd4ed9cd4b234c80446d68a755520287a630c60'
with open(keystore_path, 'r') as keyfile:
    encrypted_key = keyfile.read()

# Assuming the account is unlocked or there's no passphrase in dev mode
try:
    private_key = w3.eth.account.decrypt(encrypted_key, '')
    print('Private key:', private_key.hex())
except ValueError as e:
    print("Error decrypting key:", e)
