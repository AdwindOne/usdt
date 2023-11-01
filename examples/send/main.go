package main

import (
	"fmt"
	"log"
	"usdt"
	"usdt/rpc"
)

var (
	connCfg = &rpc.ConnConfig{
		Host: "localhost:19031",
		User: "admin3",
		Pass: "123",
	}
)

func main() {
	omni := usdtapi.NewOmniClient(connCfg)

	b, r := omni.GetBalance("mveUkR2wkxL1fVPaD7APMXwbDxbE57yDWC", 3)
	log.Printf("%s, %s\n", b, r)

	h, err := omni.Send("mveUkR2wkxL1fVPaD7APMXwbDxbE57yDWC", "mpF14fMrBJ3kLAePfHMC3Nppi2wdTZiTiq", 3, "1")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%v\n", h)

	tx := omni.ListTransactions()
	log.Printf("%v\n", tx)
}
