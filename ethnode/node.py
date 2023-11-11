#!/usr/bin/env python3
from web3 import Web3, HTTPProvider

ETHER_TRANSFER = 10000


def disp_ether_balance(address_str: str) -> float:
    # Get the balance of the default account
    bal_in_wei = w3.eth.get_balance(Web3.to_checksum_address(address_str))
    # Get the balance of the default account in ether
    balance_ether = w3.from_wei(bal_in_wei, 'ether')
    print('balance for', address_str, ' in ether:', balance_ether)
    return bal_in_wei


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

coinbase = w3.eth.defaultAccount
balance = disp_ether_balance(coinbase)

# Specify the recipient address and amount to transfer
recipient_address = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
amount_to_send = w3.to_wei(ETHER_TRANSFER, 'ether')

# Check if the balance is sufficient
if balance < amount_to_send:
    print("Insufficient balance to send the transaction")
    exit()

# Create and send the transaction
tx_hash = w3.eth.send_transaction({
    'to': recipient_address,
    'from': w3.eth.defaultAccount,
    'value': amount_to_send
})

# Wait for the transaction to be mined
tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)

print('Transfer successful with hash:', tx_receipt.transactionHash.hex())

# Check the balance of the recipient
disp_ether_balance(w3.eth.defaultAccount)
disp_ether_balance(recipient_address)
