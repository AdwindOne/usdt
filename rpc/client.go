package rpc

import (
	"bytes"
	"fmt"
	"github.com/AdwindOne/usdt/rpc/json"
	"log"
	"net/http"
)

// Implement a simple omni core RPC interface.
// And support http basic auth.
type Client struct {
	connConfig *ConnConfig
}

// NewClient returns a new rpcclient instance.
func NewClient(connConfig *ConnConfig) *Client {
	return &Client{connConfig}
}

// Call the RPC specified by the method variable.
// The args is a variable parameter list.
// The result is saved in result.
func (c *Client) Call(result interface{}, method string, args ...interface{}) error {

	url := fmt.Sprintf("http://%s/", c.connConfig.Host)
	params := make([]interface{}, 0)
	params = append(params, args...)

	message, err := json.EncodeClientRequest(method, params)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.connConfig.User, c.connConfig.Pass)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error in sending request to %s. %s", url, err)
		return err
	}
	defer resp.Body.Close()

	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Printf("Couldn't decode response. %s", err)
		return err
	}
	return nil
}
