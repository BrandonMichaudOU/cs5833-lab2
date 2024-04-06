package main

import (
	"fmt"
	"strings"

	btcvanity "github.com/MarinX/btc-vanity"
)

func main() {

	// create configuration
	cfg := &btcvanity.Config{
		// buffered channel, more buffer, faster to find matching pattern
		Buffer: 5,
		// if you want to use testnet, set true
		TestNet: false,
	}

	//btc := btcvanity.New(cfg)

	// find a patters eg adddress which starts with "ab"
	pattern := "brand"
	suffix := ""
	var address btcvanity.IWallet
	var err error
	for !strings.EqualFold(suffix, pattern) {
		btc := btcvanity.New(cfg)
		address, err = btc.Find("b")
		if err != nil {
			panic(err)
		}
		pub := address.PublicKey()
		suffix = pub[len(pub)-len(pattern):]
	}

	// print our custom public key
	fmt.Printf("PUBLIC KEY\n%s\n", address.PublicKey())

	// print our private key so it can be imported in most btc wallets
	fmt.Printf("PRIVATE KEY\n%s\n", address.PrivateKey())
}
