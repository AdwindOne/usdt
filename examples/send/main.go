package main

import (
	"fmt"
	"log"
	"usdt/omni"
	"usdt/rpcclient"
)

var (
	connCfg = &rpcclient.ConnConfig{
		Host: "localhost:19031",
		User: "admin3",
		Pass: "123",
	}
)

func main() {
	omniClient := omni.NewClient(connCfg)

	b, r := omniClient.GetBalance("mveUkR2wkxL1fVPaD7APMXwbDxbE57yDWC", 3)
	log.Printf("%s, %s\n", b, r)

	h, err := omniClient.Send("mveUkR2wkxL1fVPaD7APMXwbDxbE57yDWC", "mpF14fMrBJ3kLAePfHMC3Nppi2wdTZiTiq", 3, "1")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%v\n", h)

	tx := omniClient.ListTransactions()
	log.Printf("%v\n", tx)
}
