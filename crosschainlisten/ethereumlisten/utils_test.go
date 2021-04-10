package ethereumlisten

import "testing"

func TestSimple(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				break
			}
			t.Log(i, j)
		}
	}
}
