package tcb

import (
	"testing"
)

func TestSolve(t *testing.T) {
	text, err := Solve("data:image/png;base64,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")
	if err != nil {
		t.Error(err)
	}
	if text != "mz517p" {
		t.Errorf("text is not mz517p, got: %s", text)
	}
}