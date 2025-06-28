package common

import (
	"github.com/google/uuid"
)

type UUIDGenerator func() (uuid.UUID, error)

func NewDeterministicUUID(input string) uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(input))
}
