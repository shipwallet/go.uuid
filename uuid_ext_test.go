package uuid_test

import (
	"strings"
	"testing"

	uuid "github.com/shipwallet/go.uuid"
)

func TestStringNoDash(t *testing.T) {

	u := uuid.NewV4().StringNoDash()
	if strings.ContainsAny(u, "-") {
		t.Errorf("Generated string contains '-'")
	}

	if len(u) != 32 {
		t.Errorf("Generated string is incorrect length")
	}
}
