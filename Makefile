build:
	go build -o build

run:
	go run main.go

clean:
	go clean

# Initialize submodules
init-submodules:
	git submodule init
	git submodule update

# Update submodules to the commit specified in the superproject
update-submodules:
	git submodule update --remote

# Pull latest changes for the repo and submodules
pull-all:
	git pull --recurse-submodules
	git submodule update --init --recursive

init-contracts:
	cd contracts && npm i

test-contracts:
	cd contracts && npm run build && npm run test

build-contracts:
	cd contracts && npx hardhat compile

deploy-contracts:
	cd contracts && npx hardhat --network anvil deploy
# "SimulateTxAccessor" at 0x59AD6735bCd8152B84860Cb256dD9e96b85F69Da
# "GnosisSafeProxyFactory" at 0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2
# "DefaultCallbackHandler" at 0x1AC114C2099aFAf5261731655Dc6c306bFcd4Dbd
# "CompatibilityFallbackHandler" at 0xf48f2B2d2a534e402487b3ee7C18c33Aec0Fe5e4
# "CreateCall" at 0x7cbB62EaA69F79e6873cD1ecB2392971036cFAa4
# "MultiSend" at 0xA238CBeb142c10Ef7Ad8442C6D1f9E89e07e7761
# "MultiSendCallOnly" at 0x40A2aCCbd92BCA938b02010E17A5b8929b49130D
# "SignMessageLib" at 0xA65387F16B013cf2Af4605Ad8aA5ec25a2cbA3a2
# "GnosisSafeL2" at 0x3E5c63644E683549055b9Be8653de26E0B4CD36E
# "GnosisSafe" at 0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552

.PHONY: build init-submodules update-submodules pull-all
