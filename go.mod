module github.com/nbit99/hx_sdk

go 1.12

require (
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/nbit99/hcd v1.0.1
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/stretchr/testify v1.7.1
	github.com/nbit99/keccak v1.0.0
	github.com/pquerna/ffjson v0.0.0-20181028064349-e517b90714f7
	//github.com/denkhaus/bitshares v0.6.1-0.20190502142618-5ae8c00cb394
)

replace (
	//github.com/nbit99/hcd => ../hcd
)
