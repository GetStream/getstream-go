package getstream

import (
	"errors"
	"strings"
)

type SRTCredentials struct {
	Address string
}

func (c *Call) CreateSRTCredentials(userID string) (*SRTCredentials, error) {
	if c.data == nil {
		return nil, errors.New("call is not initialized, make sure to call .Get or .GetOrCreate first")
	}

	token, err := c.client.client.createToken(userID, nil, nil)
	if err != nil {
		return nil, err
	}

	passphrase := strings.Split(token, ".")[2]

	address := strings.ReplaceAll(c.data.Ingress.Srt.Address, "{passphrase}", passphrase)
	address = strings.ReplaceAll(address, "{token}", token)
	return &SRTCredentials{
		Address: address,
	}, nil
}
