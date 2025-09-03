package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Do not work with this architecture")
	}

	t.Fail()
}
