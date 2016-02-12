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
	"testing"
)

func TestEncodeAsciiLength(t *testing.T) {

	if string(encodeASCIILength(LenLLVAR, 8)) != "08" {
		t.Errorf("Wrong size for LLVAR")
	}

	if string(encodeASCIILength(LenLLVAR, 14)) != "14" {
		t.Errorf("Wrong size for LLVAR")
	}

	if string(encodeASCIILength(LenLLLVAR, 55)) != "055" {
		t.Errorf("Wrong size for LLLVAR")
	}

	if string(encodeASCIILength(LenLLLVAR, 128)) != "128" {
		t.Errorf("Wrong size for LLLVAR")
	}
}

func TestEncodeASCII(t *testing.T) {

	cfg := setupConfigForASCII()
	strResult := "ISO0200723C000000C0900004000000000000001644443333222211110000000000000100001504222209350001231909350422171245454545       99999999858F0F0F0F0F0F0F0F0023Esto es solo una prueba"

	msg := IsoMessage{}
	msg.SetHeader("ISO")
	msg.SetMti("0200")
	msg.AddField(2, "4444333322221111")
	msg.AddField(3, "000000")
	msg.AddField(4, "10000")
	msg.AddField(7, "150422220935")
	msg.AddField(11, "123")
	msg.AddField(12, "190935")
	msg.AddField(13, "0422")
	msg.AddField(14, "1712")
	msg.AddField(41, "45454545")
	msg.AddField(42, "99999999")
	msg.AddField(49, "858")
	msg.AddField(52, "F0F0F0F0F0F0F0F0")
	msg.AddField(70, "Esto es solo una prueba")

	result := msg.EncodeIsoMessage(cfg)

	//	t.Log(hex.Dump(result))

	if string(result) != strResult {
		t.Error("EncodeIsoMessage function returned a wrong message")
	}
}

func setupConfigForASCII() (cfg map[int]IsoFieldConfig) {

	cfg = make(map[int]IsoFieldConfig)

	cfg[FieldHeader] = IsoFieldConfig{ID: FieldHeader, Name: "Header", Length: 3, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[FieldMTI] = IsoFieldConfig{ID: FieldMTI, Name: "MTI", Length: 4, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[FieldBitmap] = IsoFieldConfig{ID: FieldBitmap, Name: "Bitmap", Length: 16, Format: FormatBitmapASCII, LenFormat: LenFixed, Padding: PaddingNone}

	cfg[2] = IsoFieldConfig{ID: 2, Name: "Primary account number", Length: 20, Format: FormatASCII, LenFormat: LenLLVAR, Padding: PaddingNone}
	cfg[3] = IsoFieldConfig{ID: 3, Name: "Processing code", Length: 6, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[4] = IsoFieldConfig{ID: 4, Name: "Amount, transaction", Length: 12, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingLeftWithZero}
	cfg[7] = IsoFieldConfig{ID: 7, Name: "Transmission date & time", Length: 10, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[11] = IsoFieldConfig{ID: 11, Name: "System trace audit number", Length: 6, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingLeftWithZero}
	cfg[12] = IsoFieldConfig{ID: 12, Name: "Time, local transaction", Length: 6, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[13] = IsoFieldConfig{ID: 13, Name: "Date, local transaction", Length: 4, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[14] = IsoFieldConfig{ID: 14, Name: "Date, expiration", Length: 4, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[22] = IsoFieldConfig{ID: 22, Name: "Point of service entry mode", Length: 3, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[24] = IsoFieldConfig{ID: 24, Name: "Network International identifier", Length: 3, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[35] = IsoFieldConfig{ID: 35, Name: "Track 2", Length: 37, Format: FormatASCII, LenFormat: LenLLVAR, Padding: PaddingNone}
	cfg[37] = IsoFieldConfig{ID: 37, Name: "Retrieval reference number", Length: 12, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[39] = IsoFieldConfig{ID: 39, Name: "Response code", Length: 2, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[41] = IsoFieldConfig{ID: 41, Name: "Terminal identification", Length: 8, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[42] = IsoFieldConfig{ID: 42, Name: "Merchant code", Length: 15, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingLeftWithSpace}
	cfg[45] = IsoFieldConfig{ID: 45, Name: "Track 1", Length: 76, Format: FormatASCII, LenFormat: LenLLVAR, Padding: PaddingNone}
	cfg[49] = IsoFieldConfig{ID: 49, Name: "Currency code", Length: 3, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[52] = IsoFieldConfig{ID: 52, Name: "PIN block", Length: 16, Format: FormatASCII, LenFormat: LenFixed, Padding: PaddingNone}
	cfg[70] = IsoFieldConfig{ID: 70, Name: "ISO Reserved", Length: 100, Format: FormatASCII, LenFormat: LenLLLVAR, Padding: PaddingNone}

	return
}
