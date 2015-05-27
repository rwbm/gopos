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

// Field formats
const (
	FORMAT_ASCII = 1 + iota
	FORMAT_BCD
	FORMAT_BITMAP_ASCII
	FORMAT_BITMAP_BCD
)

// Length types
const (
	LEN_FIXED = 1 + iota
	LEN_LLVAR
	LEN_LLLVAR
)

// Special fields IDs
const (
	FIELD_HEADER = -2
	FIELD_MTI    = -1
	FIELD_BITMAP = 1
)

// Padding options
const (
	PADDING_NONE = iota
	PADDING_RIGHT_WITHZERO
	PADDING_RIGHT_WITHF
	PADDING_RIGHT_WITHSPACE
	PADDING_LEFT_WITHZERO
	PADDING_LEFT_WITHF
	PADDING_LEFT_WITHSPACE
)
