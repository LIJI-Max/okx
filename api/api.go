package api

import (
	"context"
	"github.com/LIJI-Max/okx"
	"github.com/LIJI-Max/okx/api/rest"
	"github.com/LIJI-Max/okx/api/ws"
)

// Client is the main api wrapper of okx
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination) (*Client, error) {
	restURL := okx.RestURL
	wsPubURL := okx.PublicWsURL
	wsPriURL := okx.PrivateWsURL
	switch destination {
	case okx.AwsServer:
		restURL = okx.AwsRestURL
		wsPubURL = okx.AwsPublicWsURL
		wsPriURL = okx.AwsPrivateWsURL
	case okx.DemoServer:
		restURL = okx.DemoRestURL
		wsPubURL = okx.DemoPublicWsURL
		wsPriURL = okx.DemoPrivateWsURL
	case okx.OmegaServer:
		restURL = okx.OmegaRestURL
		wsPubURL = okx.OmegaPublicWsURL
		wsPriURL = okx.OmegaPrivateWsURL
	case okx.ColoServer:
		restURL = okx.ColoRestURL
		wsPubURL = okx.ColoPublicWsURL
		wsPriURL = okx.ColoPrivateWsURL
	case okx.BusinessServer:
		restURL = okx.RestURL
		wsPubURL = okx.BusinessWsURL
		wsPriURL = okx.BusinessWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okx.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}

func NewClientWithUrl(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination, restURL, wsPriURL, wsPubURL okx.BaseURL) (*Client, error) {
	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okx.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}

// NewClient returns a pointer to a fresh Client
func NewClientWithIP(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination, ip string) (*Client, error) {
	restURL := okx.RestURL
	wsPubURL := okx.PublicWsURL
	wsPriURL := okx.PrivateWsURL
	switch destination {
	case okx.AwsServer:
		restURL = okx.AwsRestURL
		wsPubURL = okx.AwsPublicWsURL
		wsPriURL = okx.AwsPrivateWsURL
	case okx.DemoServer:
		restURL = okx.DemoRestURL
		wsPubURL = okx.DemoPublicWsURL
		wsPriURL = okx.DemoPrivateWsURL
	case okx.OmegaServer:
		restURL = okx.OmegaRestURL
		wsPubURL = okx.OmegaPublicWsURL
		wsPriURL = okx.OmegaPrivateWsURL
	case okx.ColoServer:
		restURL = okx.ColoRestURL
		wsPubURL = okx.ColoPublicWsURL
		wsPriURL = okx.ColoPrivateWsURL
	case okx.BusinessServer:
		restURL = okx.RestURL
		wsPubURL = okx.BusinessWsURL
		wsPriURL = okx.BusinessWsURL
	}

	r := rest.NewClientWithIP(apiKey, secretKey, passphrase, restURL, destination, ip)
	c := ws.NewClientWithIP(ctx, apiKey, secretKey, passphrase, map[bool]okx.BaseURL{true: wsPriURL, false: wsPubURL}, ip)

	return &Client{r, c, ctx}, nil
}

func NewClientWithIPWithUrl(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination, ip string, restURL, wsPriURL, wsPubURL okx.BaseURL) (*Client, error) {
	r := rest.NewClientWithIP(apiKey, secretKey, passphrase, restURL, destination, ip)
	c := ws.NewClientWithIP(ctx, apiKey, secretKey, passphrase, map[bool]okx.BaseURL{true: wsPriURL, false: wsPubURL}, ip)

	return &Client{r, c, ctx}, nil
}

// NewClient returns a pointer to a fresh Client
func NewClientWithSourceAndTargetIP(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination, sourceIp, targetIp string) (*Client, error) {
	restURL := okx.RestURL
	wsPubURL := okx.PublicWsURL
	wsPriURL := okx.PrivateWsURL
	switch destination {
	case okx.AwsServer:
		restURL = okx.AwsRestURL
		wsPubURL = okx.AwsPublicWsURL
		wsPriURL = okx.AwsPrivateWsURL
	case okx.DemoServer:
		restURL = okx.DemoRestURL
		wsPubURL = okx.DemoPublicWsURL
		wsPriURL = okx.DemoPrivateWsURL
	case okx.OmegaServer:
		restURL = okx.OmegaRestURL
		wsPubURL = okx.OmegaPublicWsURL
		wsPriURL = okx.OmegaPrivateWsURL
	case okx.ColoServer:
		restURL = okx.ColoRestURL
		wsPubURL = okx.ColoPublicWsURL
		wsPriURL = okx.ColoPrivateWsURL
	case okx.BusinessServer:
		restURL = okx.RestURL
		wsPubURL = okx.BusinessWsURL
		wsPriURL = okx.BusinessWsURL
	}

	r := rest.NewClientWithIP(apiKey, secretKey, passphrase, restURL, destination, sourceIp)
	c := ws.NewClientWithSourceAndTargetIP(ctx, apiKey, secretKey, passphrase, map[bool]okx.BaseURL{true: wsPriURL, false: wsPubURL}, sourceIp, targetIp)

	return &Client{r, c, ctx}, nil
}

func NewClientWithSourceAndTargetIPWithUrl(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination, sourceIp, targetIp string, restURL, wsPriURL, wsPubURL okx.BaseURL) (*Client, error) {
	r := rest.NewClientWithIP(apiKey, secretKey, passphrase, restURL, destination, sourceIp)
	c := ws.NewClientWithSourceAndTargetIP(ctx, apiKey, secretKey, passphrase, map[bool]okx.BaseURL{true: wsPriURL, false: wsPubURL}, sourceIp, targetIp)

	return &Client{r, c, ctx}, nil
}
