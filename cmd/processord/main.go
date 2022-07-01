package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/VigorousDeveloper/poc-human/cmd"
	diverclient "github.com/VigorousDeveloper/poc-human/processor/humanclient"
	"github.com/VigorousDeveloper/poc-human/processor/humanclient/cosmos"
	"github.com/VigorousDeveloper/poc-human/processor/observer"
)

func initPrefix() {
	cosmosSDKConfg := cosmos.GetConfig()
	cosmosSDKConfg.SetBech32PrefixForAccount(cmd.Bech32PrefixAccAddr, cmd.Bech32PrefixAccPub)
	cosmosSDKConfg.Seal()
}

func main() {
	// Init prefix
	initPrefix()

	// args := os.Args
	// if len(args) < 1 {
	// 	return
	// }

	signer := "validator" //args[2]
	password := "password"

	kb, _, err := diverclient.GetKeyringKeybase("", signer, password)
	if err != nil {
		fmt.Println("fail to get keyring keybase")
		return
	}

	info, err := kb.Key(signer)
	pubKey := info.GetPubKey().Address().String()
	addr := info.GetAddress().String()

	k := diverclient.NewKeysWithKeybase(kb, signer, password)

	// setup TSS signing
	priKey, err := k.GetPrivateKey()
	if err != nil {
		fmt.Println("fail to get private key")
		return
	}

	fmt.Println(priKey)

	// cfg := &diverclient.BridgeConfig{
	// 	ChainId:         "test",
	// 	ChainHost:       "127.0.0.1:1317",
	// 	ChainRPC:        "127.0.0.1:26657",
	// 	ChainHomeFolder: "~/.diversifi/",
	// }

	cfg := &diverclient.BridgeConfig{
		ChainId:         "test",
		ChainHost:       "127.0.0.1:1317",
		ChainRPC:        "127.0.0.1:26657",
		ChainHomeFolder: "~/.poc-human/",
	}

	HumanChainBridge, err := diverclient.NewHumanChainBridge(k, cfg, signer, pubKey, addr)

	obs_storage := ""
	obs, err := observer.NewObserver(HumanChainBridge, obs_storage)
	if err != nil {
		fmt.Println("fail to create observer")
		return
	}

	if err = obs.Start(); err != nil {
		fmt.Println("fail to start observer")
		return
	}

	// tss, err := tss.NewTssSigner(HumanChainBridge)
	// if err != nil {
	// 	fmt.Println("fail to create tss")
	// 	return
	// }

	// if err = tss.Start(); err != nil {
	// 	fmt.Println("fail to start tss")
	// 	return
	// }

	// wait....
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("stop signal received")

	// stop observer
	if err := obs.Stop(); err != nil {
		fmt.Println("fail to stop observer")
	}

	// // stop tss
	// if err := tss.Stop(); err != nil {
	// 	fmt.Println("fail to stop tss")
	// }
}
