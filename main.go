package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	ksfilefullpath := ""
	kstype := ""
	kspassword := ""

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flags.StringVar(&ksfilefullpath, "keystore-src", "", "`path` to keystore to be processed")
	flags.StringVar(&kstype, "keystore-type", "", "keystore type: `pem`,`keystore`")
	flags.StringVar(&kspassword, "keystore-password", "", "keystore password")

	if err := flags.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse flags:", err)
		os.Exit(1)
	}

	keyjson, err := ioutil.ReadFile(ksfilefullpath)
	if err != nil {
		fmt.Println(err)
	}

	// Decrypt with the correct password
	key, err := keystore.DecryptKey(keyjson, kspassword)
	if err != nil {
		fmt.Println("test : json key failed to decrypt: %v", err)
	}

	fmt.Println("Private Key= " + hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)))
}
