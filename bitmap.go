package bitmap

import "fmt"

type BitMap struct {
	bits []byte
	size int
	max  int
}

// NewBitMap .
func NewBitMap(max int) *BitMap {
	return &BitMap{
		max: max,
	}
}

// Add 添加一个数据到bits中
func (b *BitMap) Add(num int) {
	if num > b.max || num < 0 {
		return
	}

	index := num / 8
	pos := num % 8

	if index+1 > b.size { // 超容量了，扩充bits容量
		b.bits = append(b.bits, make([]byte, index+1-b.size)...)
	}

	b.bits[index] |= 1 << pos

	b.size = len(b.bits)
}

// IsExist 是否存在bits中
func (b *BitMap) IsExist(num int) bool {
	if num > b.max || num < 0 {
		return false
	}

	// 数据保存在bits数组哪个下标
	index := num / 8
	// 数据在这个bits[index]的哪个位置
	pos := num % 8

	// 左位移后的数字
	n := uint8(1 << pos)

	return b.bits[index]&n == n
}

// Size bits 长度
func (b *BitMap) Size() int {
	return b.size
}

func (b *BitMap) GetBits() []byte {
	return b.bits
}

// ToString 二进制方式显示输出字符串
func (b *BitMap) ToString() string {
	return fmt.Sprintf("%b", b.bits)
}

// GetByteToInts 将指定的byte转化为[]int
func (b *BitMap) GetByteToInts(index int) []int {
	var intArr []int
	if b.size > index {
		byteNum := b.bits[index]

		for i := 0; i < 7; i++ {
			if (byteNum & (1 << i)) > 0 {
				intArr = append(intArr, index*8+i+1)
			}
		}
	}

	return intArr
}
