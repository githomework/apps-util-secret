package secret

import (
	"fmt"
	"testing"
)

func TestRandom16(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println("Random16 test:", Random16())
	}
}
