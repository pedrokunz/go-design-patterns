package common_test

import (
	"github.com/google/uuid"
	"github.com/pedrokunz/go-design-patterns/internal/common"
	"testing"
)

func TestNewDeterministicUUID(t *testing.T) {
	// Arrange
	seed := "example"
	expectedUUID := uuid.MustParse("7cb48787-6d91-5b9f-bc60-f30298ea5736")

	// Act
	deterministicUUID := common.NewDeterministicUUID(seed)

	// Assert
	if deterministicUUID != expectedUUID {
		t.Errorf("expected %s, got %s", expectedUUID, deterministicUUID)
	}
}
