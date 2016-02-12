/*
 * Copyright 2015 Robert Barreiro (rbarreiro@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gopos

import (
	"encoding/hex"
	"strings"
)

// Bitmap Represents a ISO8583 Bitmap
type Bitmap struct {
	data   []byte
	maxpos int
}

// NewBitmap creates a new Bitmap of the specified size (in bytes)
func NewBitmap(size int) (b *Bitmap) {
	b = &Bitmap{}
	b.data = make([]byte, size, size)
	b.maxpos = size * 8
	return

}

// IsSet checks if a bit is turned on; otherwise returns false
func (bmp *Bitmap) IsSet(i int) bool {
	i--
	return bmp.data[i/8]&(1<<uint(7-i%8)) != 0
}

// Set turns on the indicated bit
func (bmp *Bitmap) Set(i int) {
	i--
	bmp.data[i/8] |= 1 << uint(7-i%8)
}

// UnSet turns off the indicated bit
func (bmp *Bitmap) UnSet(i int) {
	i--
	bmp.data[i/8] &^= 1 << uint(7-i%8)
}

// ClearAll turns off all the bits
func (bmp *Bitmap) ClearAll() {
	for i := 1; i <= bmp.maxpos; i++ {
		bmp.UnSet(i)
	}
}

// Sets turns on all the bits in the list
func (bmp *Bitmap) Sets(xs ...int) {
	for _, x := range xs {
		bmp.Set(x)
	}
}

// BitsOn returns a list of bits that are currently on
func (bmp *Bitmap) BitsOn() (bits []int) {
	for i := 2; i <= bmp.maxpos; i++ {
		if bmp.IsSet(i) {
			bits = append(bits, i)
		}
	}
	return
}

// ToHexString converts the bitmap to it's Hex representation
func (bmp *Bitmap) ToHexString() string {
	return strings.ToUpper(hex.EncodeToString(bmp.data[:]))
}

// ToBCD converts the bitmap to BCD
func (bmp *Bitmap) ToBCD() (b []byte, err error) {
	b, err = bcd(hex.EncodeToString(bmp.data[:]))
	return
}

// ToBinaryString converts the bitmap to it's Binary representation
func (bmp *Bitmap) ToBinaryString() string {
	var result string

	for i := 1; i <= bmp.maxpos; i++ {
		if bmp.IsSet(i) {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result
}

// FromHex converts from Hex string to bitmap
func (bmp *Bitmap) FromHex(hexStr string) (result error) {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		result = err
	} else {
		bmp.data = b
		bmp.maxpos = len(hexStr) * 4
	}
	return
}

// FromBCD converts from BCD data to Bitmap
func (bmp *Bitmap) FromBCD(bcd []byte) (result error) {
	// TODO
	return
}
