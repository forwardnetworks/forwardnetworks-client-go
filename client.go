package forwardnetworks

import (
    "fmt"
    "io/ioutil"
	"crypto/tls"
	"net/http"
)

type Client struct {
	Username string
	Password string
	HostURL  string
	Insecure bool
	HTTPClient *http.Client
}

func NewClient(host, username, password *string, insecure bool) (*Client, error) {
    c := Client{
        HTTPClient: &http.Client{
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
            },
        },
    }

    if host != nil {
        c.HostURL = *host
    }
	
    if username == nil || password == nil {
		return &c, nil
	}

    if username != nil {
        c.Username = *username
    }

    if password != nil {
        c.Password = *password
    }

    return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
    if c.Username != "" && c.Password != "" {
        req.SetBasicAuth(c.Username, c.Password)
    }

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
    }

    return body, err
}