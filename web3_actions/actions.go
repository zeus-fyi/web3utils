package web3_actions

import (
	"github.com/zeus-fyi/web3utils/accounts"
	web3_client "github.com/zeus-fyi/web3utils/client"
)

type Web3Actions struct {
	web3_client.Client
	*accounts.Account
	NodeURL string
	Network string
}

func (w *Web3Actions) Dial() {
	c, err := web3_client.Dial(w.NodeURL)
	if err != nil {
		panic(err)
	}
	w.Client = c
}

func NewWeb3ActionsClient(nodeUrl string) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
	}
}

func NewWeb3ActionsClientWithAccount(nodeUrl string, account *accounts.Account) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
		Account: account,
	}
}
