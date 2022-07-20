package tools

import (
	"testing"
)

func TestRandomHash(t *testing.T) {
	v := "https://google.com"

	hash, err := RandomHash(v)

	if err != nil {
		t.Errorf("TestRandomHash: %s", err)
	} else if len(hash) != 6 {
		t.Errorf("TestRandomHash: %s", err)
	}
}
