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
	"testing"
)

func TestBitmap(t *testing.T) {

	b := NewBitmap(16)
	b.Sets(2, 3, 4, 7, 11, 12, 13, 14, 41, 42, 49, 70, 100)

	bStrHex := b.ToHexString()
	bStrBin1 := b.ToBinaryString()

	b2 := NewBitmap(16)
	b2.FromHex(bStrHex)
	bStrBin2 := b2.ToBinaryString()

	if bStrBin1 != bStrBin2 {
		t.Error("Failed to convert from Hex String to Bitmap")
	}
}

func TestFromHex(t *testing.T) {

	str := "FF00FFZZ44"
	b := NewBitmap(16)

	err := b.FromHex(str)
	if err == nil {
		t.Error("FromHex() should give an error decoding ", str)
	}
}

func TestToBCD(t *testing.T) {

	b := NewBitmap(16)
	b.Sets(2, 3, 4, 7, 11, 12, 13, 14, 41, 42, 49, 70, 100)
	data, err := b.ToBCD()

	if err == nil {
		bStrHex := b.ToHexString()
		bStrHex2 := strings.ToUpper(hex.EncodeToString(data))

		if bStrHex != bStrHex2 {
			t.Error("Original and converted are not equal")
		}
	} else {
		t.Error("ToBCD() function failed:", err)
	}
}
