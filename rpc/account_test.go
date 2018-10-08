package rpc

import (
	"fmt"
	"testing"
)

var (
	//client = NewClient("https://s.altnet.rippletest.net:51234", "https://testnet.data.api.ripple.com")
	client = NewClient("http://47.75.70.201:9003", "http://47.75.70.201:9003")
	//client = NewClient("https://data.ripple.com")
)

func TestGetAccountBalance(t *testing.T) {
	address := "rh4WZwXaDhamjM7hw8gArB9Jgs6fkxUGnw"
	res, err := client.GetAccountBalances(address, map[string]string{})
	if err != nil {
		t.Error("get err: ", err)
	}
	for _, v := range res.Balances {
		fmt.Printf("balance: %+v\n", v)
	}
}
