package forwardnetworks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetVersion function retrieves the version of the Forward Networks API.
func (c *Client) GetAllNetworks() (*[]Network, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	networks := []Network{}
	err = json.Unmarshal(body, &networks)
	if err != nil {
		return nil, err
	}

	return &networks, nil
}

func (c *Client) GetNetwork(networkID string) (*Network, error) {
	networks, err := c.GetAllNetworks()
	if err != nil {
		return nil, err
	}

	for _, network := range *networks {
		if network.ID == networkID {
			return &network, nil
		}
	}

	return nil, fmt.Errorf("network with ID %s not found", networkID)
}

func (c *Client) CreateNetwork(networkName string) (*Network, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/networks?name=%s", c.HostURL, networkName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	networks := Network{}
	err = json.Unmarshal(body, &networks)
	if err != nil {
		return nil, err
	}

	return &networks, nil
}

func (c *Client) UpdateNetwork(networkID string, networkItems []Network) (*Network, error) {
	rb, err := json.Marshal(networkItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/networks/%s", c.HostURL, networkID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	network := Network{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

// DeleteNetwork - Deletes an network
func (c *Client) DeleteNetwork(networkID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/networks/%s", c.HostURL, networkID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}