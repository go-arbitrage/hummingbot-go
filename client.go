package hummingbot_go

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	CaFile   string `json:"caFile"`
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
	Password string `json:"password"`

	Host string `json:"host"`
	Port string `json:"port"`
}

type Client struct {
	cfg Config

	client *http.Client
}

func New(cfg Config) *Client {
	return &Client{cfg: cfg}
}

func (c *Client) Init() error {
	// Load CA cert
	caCert, err := os.ReadFile(c.cfg.CaFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Load client cert
	cert, err := tls.LoadX509KeyPair(c.cfg.CertFile, c.cfg.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	c.client = &http.Client{Transport: transport}

	return nil
}

func (c *Client) Ping() (string, error) {
	url := c.formatURL("/")
	// Do GET something
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Dump response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *Client) Price(req PriceRequest) (*PriceResponse, error) {
	url := c.formatURL("/amm/price")
	var response PriceResponse
	if err := c.doRequest(url, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) formatURL(path string) string {
	return fmt.Sprintf("https://%s:%s%s", c.cfg.Host, c.cfg.Port, path)
}

func (c *Client) doRequest(url string, request, response interface{}) error {
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	resp, err := c.client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//log.Printf("received data: %s", string(data))
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}
	return nil
}
