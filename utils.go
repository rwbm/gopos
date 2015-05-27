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
)

// Left padding
func strPadLeft(str, pad string, length int) string {
	if len(str) < length {
		for {
			str = pad + str
			if len(str) >= length {
				return str[0:length]
			}
		}
	} else {
		return str[0:length]
	}
}

// Right padding
func strPadRight(str, pad string, length int) string {
	if len(str) < length {
		for {
			str += pad
			if len(str) >= length {
				return str[0:length]
			}
		}
	} else {
		return str[0:length]
	}
}

// BCD format conversion
func bcd(data string) (b []byte, err error) {
	if len(data)%2 != 0 {
		errors.New("Length has to be even")
	}
	b, err = hex.DecodeString(data)

	return
}
