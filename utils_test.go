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

func TestBcd(t *testing.T) {

	strOk := "80FFA256AE10"
	strBinOk, err1 := bcd(strOk)

	if err1 != nil {
		t.Error("bcd() failed: ", err1)
	}
	strOkRecover := strings.ToUpper(hex.EncodeToString(strBinOk))

	if strOk != strOkRecover {
		t.Error("Original and converted does not match")
	}

	strNotEven := "80FFA256AE1"
	_, err2 := bcd(strNotEven)
	if err2 == nil {
		t.Error("bcd() should have failed: ")
	}

}
