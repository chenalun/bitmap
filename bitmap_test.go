package bitmap_test

import (
	"bitmap"
	"fmt"
	"math/rand"
	"testing"
	"time"
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

// 20 亿个数字在 4G 内存中如何去重排序
func TestBigMap(t *testing.T) {
	s := time.Now()

	var total = 2e9 // 20亿
	total1 := int(total)

	b := bitmap.NewBitMap(total1)

	for i := 0; i < total1; i++ {
		b.Add(rand.Intn(total1)) // 从1~20亿 取一个随机数
	}

	fmt.Println(fmt.Sprintf("99 IsExist %v, 999 IsExist %v,9999 IsExist %v,99999 IsExist %v,999999 IsExist %v", b.IsExist(99), b.IsExist(999), b.IsExist(9999), b.IsExist(99999), b.IsExist(999999)))
	fmt.Println(time.Since(s), b.Size())

	fmt.Println(b.GetBits()[b.Size()-2], fmt.Sprintf("%b", b.GetBits()[b.Size()-2]), b.GetByteToInts(b.Size()-2))
	fmt.Println(b.GetBits()[b.Size()-1], fmt.Sprintf("%b", b.GetBits()[b.Size()-1]), b.GetByteToInts(b.Size()-1))

}
