package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"topic-service/pkg/consul"

	"github.com/hashicorp/consul/api"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GatewayClient struct {
	ServiceName      string
	Token            string
	HTTPClient       HTTPClient
	ServiceDiscovery consul.ServiceDiscovery
}

func NewGatewayClient(serviceName, token string, consulClient *api.Client, httpClient HTTPClient) (*GatewayClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	sd, err := consul.NewServiceDiscovery(consulClient, serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to init service discovery: %v", err)
	}

	return &GatewayClient{
		ServiceName:      serviceName,
		Token:            token,
		HTTPClient:       httpClient,
		ServiceDiscovery: sd,
	}, nil
}

// Call gọi API tới service khác thông qua Consul discovery
func (c *GatewayClient) Call(method, path string, body interface{}) ([]byte, error) {
	service, err := c.ServiceDiscovery.DiscoverService()
	if err != nil {
		return nil, fmt.Errorf("service discovery failed: %v", err)
	}

	var reqBody io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body failed: %v", err)
		}
		reqBody = bytes.NewReader(jsonBytes)
	}

	url := fmt.Sprintf("http://%s:%d%s", service.ServiceAddress, service.ServicePort, path)

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http call failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("http error: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
