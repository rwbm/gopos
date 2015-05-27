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
	"errors"
	"fmt"
	"strconv"
)

// Extracts all the data from a byte array to an internal IsoMessage
func DecodeIsoMessage(data []byte, cfg map[int]IsoFieldConfig) (msg IsoMessage, err error) {

	msg = IsoMessage{}
	var offset int

	// Header
	var header Field
	header, offset = decodeField(data, offset, cfg[FIELD_HEADER])
	msg.header = header

	// MTI
	var mti Field
	mti, offset = decodeField(data, offset, cfg[FIELD_MTI])
	msg.mti = mti

	// Bitmap(s)
	var bmp BitmapField
	bmp, offset, err = decodeBitmap(data, offset, cfg[FIELD_BITMAP])
	if err != nil {
		return
	}
	msg.bitmap = bmp

	// If bitmap was 'correctly' decoded, we analyze it to extract the present fields
	bitsOn := msg.bitmap.Value.BitsOn()
	for i := 0; i < len(bitsOn); i++ {
		var f Field
		f, offset = decodeField(data, offset, cfg[bitsOn[i]])
		msg.AddField(f.ID, f.Value)
	}

	return
}

func decodeField(data []byte, offset int, cfg IsoFieldConfig) (f Field, o int) {

	if cfg.Format == FORMAT_ASCII {

		value, newOffset := decodeAsciiField(data, offset, cfg)
		f = Field{
			ID:    cfg.ID,
			Value: value,
		}
		o = newOffset

	} else if cfg.Format == FORMAT_BCD {

		// TODO

	} else {
		panic("Unknown format")
	}

	return
}

func decodeAsciiField(data []byte, offset int, cfg IsoFieldConfig) (s string, o int) {

	l, newOffset := decodeAsciiLength(data, offset, cfg)

	if len(data) >= newOffset+l {
		s = string(data[newOffset : newOffset+l])
		o = newOffset + l
	} else {
		panic(fmt.Sprintf("Not enough data: ID=%d, Offset=%d", cfg.ID, offset))
	}

	return
}

func decodeBitmap(data []byte, offset int, cfg IsoFieldConfig) (bmp BitmapField, o int, err error) {

	if cfg.Format == FORMAT_BITMAP_ASCII {

		value, newOffset := decodeAsciiBitmap(data, offset, cfg)

		bmp = BitmapField{}
		bmp.Value = &Bitmap{}
		errTemp := bmp.Value.FromHex(value)
		if errTemp == nil {
			o = newOffset
		} else {
			err = errors.New(fmt.Sprintf("Could not decide the Bitmap: ", value))
		}

	} else if cfg.Format == FORMAT_BITMAP_BCD {

		// TODO

	} else {
		panic("Unknown bitmap format")
	}

	return
}

func decodeAsciiBitmap(data []byte, offset int, cfg IsoFieldConfig) (s string, o int) {

	bitmapLen := cfg.Length * 2

	ss := string(data[offset : offset+bitmapLen])
	dd, _ := hex.DecodeString(ss)

	// TODO: Capturar el error anterior

	s = string(hex.EncodeToString(dd))
	o = offset + bitmapLen

	return
}

func decodeAsciiLength(data []byte, offset int, cfg IsoFieldConfig) (l int, o int) {

	if cfg.LenFormat == LEN_FIXED {
		l = cfg.Length
		o = offset
	} else if cfg.LenFormat == LEN_LLVAR || cfg.LenFormat == LEN_LLLVAR {

		var lenSize int
		if cfg.LenFormat == LEN_LLVAR {
			lenSize = 2
		} else {
			lenSize = 3
		}

		d := data[offset : offset+lenSize]
		strd := string(d)

		if len(strd) > 0 {
			var err error
			l, err = strconv.Atoi(strd)
			if err != nil {
				panic(fmt.Sprintf("Could not decode length for field %d. Offset=%d", cfg.ID, offset))
			}
			o = offset + lenSize
		} else {
			panic(fmt.Sprintf("Could not decode length for field %d. Offset=%d", cfg.ID, offset))
		}
	}

	return
}
