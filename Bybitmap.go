package ByGolang

import "math"

type BitMap []uint64

func NewBitMap(length int) BitMap {
	bitLen := (length + 63) / 64
	return make([]uint64, bitLen)
}

func (this BitMap) Len() int {
	return len(this) * 64
}

func (this BitMap) SetBit(index uint64) {
	this[index/64] |= 1 << (index % 64)
}

func (this BitMap) SetBitEx(index, value uint64) {
	this[index] |= value
}

func (this BitMap) ClearBit(index uint64) {
	this[index/64] &= math.MaxUint64 - (1 << (index % 64))
}

func (this BitMap) ToNumbers() []int {
	res := make([]int, 0, this.Len())
	for i := 0; i < len(this); i++ {
		if this[i] == 0 {
			continue
		}
		for index := uint(0); index < 64; index++ {
			if this[i]&(1<<index) > 0 {
				res = append(res, 64*i+int(index))
			}
		}
	}
	return res
}

func (this BitMap) IsZero() bool {
	for i := 0; i < len(this); i++ {
		if this[i] != 0 {
			return false
		}
	}
	return true
}

func (this BitMap) Clone() BitMap {
	res := NewBitMap(this.Len())
	for i := 0; i < len(this); i++ {
		res[i] = this[i]
	}
	return res
}

func And(bitMapArray ...BitMap) BitMap {
	res := NewBitMap(bitMapArray[0].Len())
	var bits uint64
	for i := 0; i < len(bitMapArray[0]); i++ {
		bits = math.MaxUint64
		for j := 0; j < len(bitMapArray); j++ {
			if bits == 0 {
				break
			}
			bits &= bitMapArray[j][i]
		}
		res[i] = bits
	}
	return res
}

func Or(bitMapArray ...BitMap) BitMap {
	res := NewBitMap(bitMapArray[0].Len())
	var bits uint64
	for i := 0; i < len(bitMapArray[0]); i++ {
		bits = 0
		for j := 0; j < len(bitMapArray); j++ {
			bits |= bitMapArray[j][i]
		}
		res[i] = bits
	}
	return res
}

func Not(bitMap BitMap) BitMap {
	res := NewBitMap(bitMap.Len())
	for i := 0; i < len(bitMap); i++ {
		res[i] = ^bitMap[i]
	}
	return res
}
