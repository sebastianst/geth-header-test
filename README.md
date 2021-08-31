# Geth Header.Hash Mismatch

Run like
```sh
# use default example block 5410648 for Goerli and default Metamask Goerli endpoint
go run main.go

# or specify block number, expected hash and node URL
# last pre-London block https://goerli.etherscan.io/block/5062604
go run main.go -n 5062604 -h 0x7411a2c91948f961c9e1447c4d223c8b27fdf02e147a0bbaaa8402eb74d2d0a6 -u $ETH_GOERLI_NODE_URL
# some very early block https://goerli.etherscan.io/block/420
go run main.go -n 420 -h 0x863e87aeee886a822999f0e110e45a976bbdf2ed869519ee914e07ca7006b9c7
# also fails on Rinkeby https://rinkeby.etherscan.io/block/9210679
go run main.go -n 9210679 -h 0xe4de20e03c8bcb780a2144d898365c315b4d060d9ab46415bab6d585f4dc4696 -u $ETH_RINKEBY_NODE_URL

# Everything is fine on mainnet https://etherscan.io/block/13132500
go run main.go -n 13132500 -h 0x7816d95f8ddb050425528d700abf7b7cab00742b08645d70b4a3057ea4ff804e -u $ETH_MAINNET_NODE_URL
```
