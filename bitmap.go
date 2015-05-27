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

// Thanks:
// - Egon Elbre (egonelbre@gmail.com) who provided part of this implementation and
// - Dan Kortschak (dan.kortschak@adelaide.edu.au) who fixed the IsSet() function
//
package gopos

import (
	"encoding/hex"
	"strings"
)

type Bitmap struct {
	data   []byte
	maxpos int
}

// Creates a new Bitmap of the specified size (in bytes)
func NewBitmap(size int) (b *Bitmap) {
	b = &Bitmap{}
	b.data = make([]byte, size, size)
	b.maxpos = size * 8
	return

}

// Checks if a bit is turned on; otherwise returns false
func (this *Bitmap) IsSet(i int) bool {
	i -= 1
	return this.data[i/8]&(1<<uint(7-i%8)) != 0
}

// Turns on the indicated bit
func (this *Bitmap) Set(i int) {
	i -= 1
	this.data[i/8] |= 1 << uint(7-i%8)
}

// Turns off the indicated bit
func (this *Bitmap) UnSet(i int) {
	i -= 1
	this.data[i/8] &^= 1 << uint(7-i%8)
}

// Turns off all the bits
func (this Bitmap) ClearAll() {
	for i := 1; i <= this.maxpos; i++ {
		this.UnSet(i)
	}
}

// Turns on all the bits in the list
func (this *Bitmap) Sets(xs ...int) {
	for _, x := range xs {
		this.Set(x)
	}
}

func (this *Bitmap) BitsOn() (bits []int) {
	for i := 2; i <= this.maxpos; i++ {
		if this.IsSet(i) {
			bits = append(bits, i)
		}
	}
	return
}

// Converts the bitmap to it's Hex representation
func (this Bitmap) ToHexString() string {
	return strings.ToUpper(hex.EncodeToString(this.data[:]))
}

// Converts the bitmap to BCD
func (this Bitmap) ToBCD() (b []byte, err error) {
	b, err = bcd(hex.EncodeToString(this.data[:]))
	return
}

// Converts the bitmap to it's Binary representation
func (this Bitmap) ToBinaryString() string {
	var result string

	for i := 1; i <= this.maxpos; i++ {
		if this.IsSet(i) {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result
}

// Converts from Hex string to bitmap
func (this *Bitmap) FromHex(hexStr string) (result error) {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		result = err
	} else {
		this.data = b
		this.maxpos = len(hexStr) * 4
	}
	return
}

// Converts from BCD data to Bitmap
func (this Bitmap) FromBCD(bcd []byte) (result error) {
	// TODO
	return
}
