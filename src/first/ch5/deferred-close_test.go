package ch5

import (
	"testing"
)

func TestCloseExperiment(t *testing.T) {
	n, e := CloseExperiment("/tmp/from", "/tmp/to")
	if e != nil {
		t.Errorf("could not copy cleanly, wrote %d bytes; the error: %v", n, e)
	}
}

