package labuladong

import (
	"fmt"
	"math/bits"
)

func evenOddBit1(n int) []int {
	res := []int{}
	var binN string
	odd, even := 0, 0
	if n > 0 {
		binN = fmt.Sprintf("%b", n)
	} else {
		binN = fmt.Sprintf("%b", uint64(n))
	}

	length := len(binN)
	for i := length - 1; i >= 0; i-- {
		if binN[i] == '1' {
			if (length-1-i)%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	res = append(res, even, odd)
	return res
}
func evenOddBit2(n int) []int {
	res := []int{0, 0}
	for i := 0; n > 0; n >>= 1 {
		res[i] += n & 1
		i ^= 1
	}
	return res
}

func evenOddBit3(n int) []int {

	mask1 := 0x55555555
	mask2 := 0xaaaaaaaa

	return []int{bits.OnesCount(uint(n & mask1)), bits.OnesCount(uint(n & mask2))}
}
