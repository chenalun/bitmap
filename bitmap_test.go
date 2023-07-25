package bitmap_test

import (
	"bitmap"
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	b := bitmap.NewBitMap(1000)
	for i := 0; i < 1000; i++ {
		if i < 8 {
			b.Add(i)
		}

		if i >= 990 {
			b.Add(i)
		}
	}

	fmt.Println(b.GetBits())
	fmt.Println(b.ToString())
	fmt.Println(fmt.Sprintf("size=%d", b.Size()))
	fmt.Println(fmt.Sprintf("IsExist %v,%v", b.IsExist(990), b.IsExist(991)))
}
