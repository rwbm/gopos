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

func TestDecodeASCII(t *testing.T) {

	cfg := setupConfigForASCII()
	dataISO := []byte("ISO0200723C000000C09000040000000000000016444433332222111100000000000001000015042222093500012319093504221712454545       99999999858F0F0F0F0F0F0F0F0023Esto es solo una prueba")

	msg, _ := DecodeIsoMessage(dataISO, cfg)

	t.Log(msg.header.Value)
	t.Log(msg.mti.Value)
	t.Log(msg.bitmap.Value.ToHexString())

	msgRaw := msg.EncodeIsoMessage(cfg)

	if string(dataISO) != string(msgRaw) {
		t.Error("Decoded ISO and original are not equal")
	}

}
