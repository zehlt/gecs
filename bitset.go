package gecs

import (
	"fmt"
	"strings"
)

type bitset struct {
	bytes []uint8
}

func newBitset() *bitset {
	return &bitset{
		bytes: make([]uint8, 0),
	}
}

func (b *bitset) Get(n int) bool {
	// TODO: negative numbers
	if n >= len(b.bytes)*8 {
		return false
	}

	byteIndex := n / 8
	bitPosition := n % 8

	return (b.bytes[byteIndex] & (1 << bitPosition)) != 0
}

func (b *bitset) Set(n int, val bool) {
	byteIndex := n / 8
	bitPosition := n % 8

	if byteIndex+1 > len(b.bytes) {
		diff := byteIndex + 1 - len(b.bytes)
		for i := 0; i < diff; i++ {
			b.bytes = append(b.bytes, 0)
		}
	}

	if val {
		b.bytes[byteIndex] |= (1 << bitPosition)
	} else {
		b.bytes[byteIndex] &= (0b11111111 - (1 << bitPosition))
	}
}

func (b *bitset) Include(other *bitset) bool {
	if other.Len() > b.Len() {
		return false
	}

	for i := 0; i < other.Len(); i++ {
		if b.bytes[i] != other.bytes[i] {
			return false
		}
	}

	return true
}

func (b *bitset) Contains(matcher *bitset) bool {
	// TODO: the other may be longer with only 00000
	if matcher.Len() > b.Len() {
		return false
	}

	for i := 0; i < matcher.Len(); i++ {
		match_byte := matcher.bytes[i]

		if match_byte == 0 {
			continue
		}

		if (b.bytes[i] & match_byte) != match_byte {
			return false
		}
	}

	return true
}

func (b *bitset) Equal(other *bitset) bool {
	if b.Len() != other.Len() {
		return false
	}

	for i := 0; i < b.Len(); i++ {
		if b.bytes[i] != other.bytes[i] {
			return false
		}
	}

	return true
}

func (b *bitset) Len() int {
	return len(b.bytes)
}

func (b *bitset) String() string {
	var sb strings.Builder

	for i := len(b.bytes) - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%08b ", b.bytes[i]))
	}

	return sb.String()
}
