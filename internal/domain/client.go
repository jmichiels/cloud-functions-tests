package domain

import "github.com/pkg/errors"

const (
	clientLegalAge = 16
)

// Client represents a client of the bank.
type Client struct {
	Id        UniqueId
	FirstName string
	LastName  string
	Age       uint
}

// Validate ensures the Client is valid.
func (client *Client) Validate() error {
	if err := client.Id.Validate(); err != nil {
		return errors.Wrap(err, "invalid id")
	}
	if client.FirstName == "" {
		return errors.New("empty first name")
	}
	if client.LastName == "" {
		return errors.New("empty last name")
	}
	if client.Age < clientLegalAge {
		return errors.New("under legal Age")
	}
	return nil
}
