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
	"strconv"
)

// Converts an IsoMessage to its raw representation
func (msg *IsoMessage) EncodeIsoMessage(cfg map[int]IsoFieldConfig) (b []byte) {

	// Header
	header := encodeField(msg.header.Value, cfg[FIELD_HEADER])

	// Message Type Indicator
	mti := encodeField(msg.mti.Value, cfg[FIELD_MTI])

	// Bitmap
	refreshBitmap(msg)
	bitmap := encodeBitmapField(msg.bitmap, cfg[FIELD_BITMAP])

	// Fields
	fields := encodeFields(msg.fields, cfg)

	// Result
	b = append(header, mti...)
	b = append(b, bitmap...)
	b = append(b, fields...)

	return
}

func encodeField(value string, cfg IsoFieldConfig) (b []byte) {

	if cfg.Format == FORMAT_ASCII {

		b = encodeAsciiField(value, cfg)

	} else if cfg.Format == FORMAT_BCD {

		b = encodeBinaryField(value, cfg)

	} else {

		panic("Unknown format")
	}

	return
}

func encodeFields(fields []Field, cfg map[int]IsoFieldConfig) (b []byte) {

	var accumulate []byte

	for f := range fields {
		accumulate = append(accumulate, encodeField(fields[f].Value, cfg[fields[f].ID])...)
	}

	b = accumulate
	return
}

func refreshBitmap(msg *IsoMessage) {

	if msg.maxID > 64 && msg.maxID <= 128 {
		msg.bitmap.Value = NewBitmap(16) // bitmaps 1 and 2
	} else if msg.maxID <= 64 {
		msg.bitmap.Value = NewBitmap(8) // bitmap 1 only
	} else {
		msg.bitmap.Value = NewBitmap(24) // bitmap 1, 2 and 3
	}

	for f := range msg.fields {
		if msg.fields[f].ID >= 1 {
			msg.bitmap.Value.Set(msg.fields[f].ID)
		}
	}
}

func encodeBitmapField(bmp BitmapField, cfg IsoFieldConfig) (b []byte) {

	if cfg.Format == FORMAT_BITMAP_ASCII {

		b = []byte(bmp.Value.ToHexString())

	} else if cfg.Format == FORMAT_BITMAP_BCD {

		data, err := bmp.Value.ToBCD()
		if err != nil {
			panic(err)
		}
		b = data
	}
	return
}

func encodeAsciiField(value string, cfg IsoFieldConfig) (b []byte) {

	// Field length
	l := encodeAsciiLength(cfg.LenFormat, len(value))

	// Field value
	var v []byte

	if cfg.Padding == PADDING_NONE {
		v = []byte(value)
	} else if cfg.Padding == PADDING_LEFT_WITHF || cfg.Padding == PADDING_LEFT_WITHSPACE || cfg.Padding == PADDING_LEFT_WITHZERO {
		v = []byte(padLeft(value, cfg.Padding, cfg.Length))
	} else if cfg.Padding == PADDING_RIGHT_WITHF || cfg.Padding == PADDING_RIGHT_WITHSPACE || cfg.Padding == PADDING_RIGHT_WITHZERO {
		v = []byte(padRight(value, cfg.Padding, cfg.Length))
	}

	b = append(l, v...)
	return
}

func encodeBinaryField(value string, cfg IsoFieldConfig) (b []byte) {

	// TODO

	return
}

func encodeAsciiLength(lenFormat int, lenValue int) (b []byte) {
	if lenFormat == LEN_LLVAR {
		b = []byte(strPadLeft(strconv.Itoa(lenValue), "0", 2))
	} else if lenFormat == LEN_LLLVAR {
		b = []byte(strPadLeft(strconv.Itoa(lenValue), "0", 3))
	}
	return
}

func padLeft(value string, padding int, length int) (s string) {

	var padChar string = " "

	if padding == PADDING_LEFT_WITHF {
		padChar = "F"
	} else if padding == PADDING_LEFT_WITHSPACE {
		padChar = " "
	} else if padding == PADDING_LEFT_WITHZERO {
		padChar = "0"
	}

	s = strPadLeft(value, padChar, length)
	return
}

func padRight(value string, padding int, length int) (s string) {

	var padChar string = " "

	if padding == PADDING_RIGHT_WITHF {
		padChar = "F"
	} else if padding == PADDING_RIGHT_WITHSPACE {
		padChar = " "
	} else if padding == PADDING_RIGHT_WITHZERO {
		padChar = "0"
	}

	s = strPadRight(value, padChar, length)

	return
}
