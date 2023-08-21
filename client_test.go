package hummingbot_go

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

var (
	client *Client
)

func TestMain(m *testing.M) {
	caFile := os.Getenv("CA")
	certFile := os.Getenv("CERT")
	keyFile := os.Getenv("KEY")
	password := os.Getenv("PASS")
	//log.Printf("caFile = %s, certFile = %s, keyFile = %s, password = %s", caFile, certFile, keyFile, password)

	cfg := Config{
		CaFile:   caFile,
		CertFile: certFile,
		KeyFile:  keyFile,
		Password: password,
		Host:     "localhost",
		Port:     "15888",
	}
	client = New(cfg)
	if err := client.Init(); err != nil {
		log.Fatalf("init error: %s", err)
	}
	os.Exit(m.Run())
}

func TestClient_Ping(t *testing.T) {
	//require.NoError(t, err)
}

func TestClient_Price(t *testing.T) {
	param := PriceRequest{
		Chain:     "polygon",
		Network:   "mainnet",
		Connector: "uniswap",
		Base:      "WMATIC",
		Quote:     "USDC",
		Amount:    "0.1",
		Side:      "BUY",
	}
	resp, err := client.Price(param)
	require.NoError(t, err)
	data, err := json.Marshal(resp)
	require.NoError(t, err)
	t.Logf("received: %s", string(data))
}
