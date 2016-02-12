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
	FormatASCII = 1 + iota
	FormatBCD
	FormatBitmapASCII
	FormatBitmapBCD
)

// Length types
const (
	LenFixed = 1 + iota
	LenLLVAR
	LenLLLVAR
)

// Special fields IDs
const (
	FieldHeader = -2
	FieldMTI    = -1
	FieldBitmap = 1
)

// Padding options
const (
	PaddingNone = iota
	PaddingRightWithZero
	PaddingRightWithF
	PaddingRightWithSpace
	PaddingLeftWithZero
	PaddingLeftWithF
	PaddingLeftWithSpace
)
