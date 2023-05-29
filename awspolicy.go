package fwdnet

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetVersion function retrieves the version of the Forward Networks API.
func (c *Client) GetAwsPolicy() (*AwsPolicy, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/cloud/aws-policy", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	awsPolicy := AwsPolicy{}
	err = json.Unmarshal(body, &awsPolicy)
	if err != nil {
		return nil, err
	}

	return &awsPolicy, nil
}