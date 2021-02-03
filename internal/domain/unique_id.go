package domain

import (
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

// UniqueId represents a universally unique Id.
//
// The zero value of this type is invalid. It should be created using GenerateRandomUniqueId or ParseUniqueIdFromString.
type UniqueId struct {
	id xid.ID
}

// GenerateRandomUniqueId returns a random unique id.
func GenerateRandomUniqueId() UniqueId {
	return UniqueId{
		id: xid.New(),
	}
}

// ParseUniqueIdFromString parses a unique id from its string representation.
func ParseUniqueIdFromString(str string) (UniqueId, error) {
	uid, err := xid.FromString(str)
	if err != nil {
		return UniqueId{}, errors.Wrap(err, "cannot parse unique id from string")
	}
	return UniqueId{id: uid}, nil
}

// Validate ensures the UniqueId is valid.
func (uid UniqueId) Validate() error {
	if uid.id.IsNil() {
		return errors.New("unique id cannot be nil")
	}
	return nil
}

// String implements io.Stringer
func (uid UniqueId) String() string {
	return uid.id.String()
}
