package tcb

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	generatedHash := GenerateHash()
	if len(generatedHash) == 0 {
		t.Error("generatedHash is empty")
	}
	if len(generatedHash) != 11 {
		t.Errorf("generatedHash is not 11 characters, got: %s", generatedHash)
	}
}
