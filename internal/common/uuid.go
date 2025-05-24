package common

import "github.com/google/uuid"

func NewDeterministicUUID(seed string) uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(seed))
}
