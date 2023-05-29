package forwardnetworks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetVersion function retrieves the version of the Forward Networks API.
func (c *Client) GetVersion() (*Version, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/version", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	version := Version{}
	err = json.Unmarshal(body, &version)
	if err != nil {
		return nil, err
	}

	return &version, nil
}