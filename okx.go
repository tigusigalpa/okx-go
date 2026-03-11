package okx

import (
	"github.com/tigusigalpa/okx-go/rest/account"
	"github.com/tigusigalpa/okx-go/rest/asset"
	"github.com/tigusigalpa/okx-go/rest/market"
	"github.com/tigusigalpa/okx-go/rest/public"
	"github.com/tigusigalpa/okx-go/rest/support"
	"github.com/tigusigalpa/okx-go/rest/system"
	"github.com/tigusigalpa/okx-go/rest/trade"
	"github.com/tigusigalpa/okx-go/rest/users"
)

type RestClient struct {
	*Client
	Account *account.Client
	Trade   *trade.Client
	Market  *market.Client
	Public  *public.Client
	Asset   *asset.Client
	System  *system.Client
	Support *support.Client
	Users   *users.Client
}

func NewRestClient(apiKey, secretKey, passphrase string, opts ...Option) *RestClient {
	client := NewClient(apiKey, secretKey, passphrase, opts...)

	return &RestClient{
		Client:  client,
		Account: account.NewClient(client.do),
		Trade:   trade.NewClient(client.do),
		Market:  market.NewClient(client.doPublic),
		Public:  public.NewClient(client.doPublic),
		Asset:   asset.NewClient(client.do),
		System:  system.NewClient(client.doPublic),
		Support: support.NewClient(client.doPublic),
		Users:   users.NewClient(client.do),
	}
}
