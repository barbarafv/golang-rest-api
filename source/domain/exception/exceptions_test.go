package exception

import (
	"testing"
)

func TestAbc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	panic(NewNotFoundException("Field XPTO was not found"))
}
