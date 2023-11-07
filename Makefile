build:
	go build

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

.PHONY: build init-submodules update-submodules pull-all
