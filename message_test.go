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

func TestMessages(t *testing.T) {

	msg := IsoMessage{}

	// Test MTI
	msg.SetMTI("0800")
	if msg.mti.Value != "0800" {
		t.Errorf("Wrong MTI value: %s", msg.mti.Value)
	}

	// Test header
	msg.SetHeader("ISO")
	if msg.header.Value != "ISO" {
		t.Errorf("Wrong Header value: %s", msg.header.Value)
	}

	// Test some fields
	msg.AddField(3, "000000")
	procCode, err := msg.GetValue(3)
	if err != nil || procCode != "000000" {
		t.Errorf("Wrong value for Field 3: %s", procCode)
	}

	msg.AddField(7, "0212213400")
	transDateTime, err := msg.GetValue(7)
	if err != nil || transDateTime != "0212213400" {
		t.Errorf("Wrong value for Field 7: %s", transDateTime)
	}

	msg.AddField(11, "123456")
	stanStr, err := msg.GetValue(11)
	if err != nil || stanStr != "123456" {
		t.Errorf("Wrong value for Field 11: %s", stanStr)
	}

	// Remove field
	msg.RemoveField(7)
	if msg.HasField(7) {
		t.Error("Field 11 wasn't removed")
	}

	// Get value as int
	stanInt, err := msg.GetInt(11)
	if err != nil || stanInt != 123456 {
		t.Errorf("Wrong value for Field 11: %d", stanInt)
	}

	xmlDump := msg.DumpXML()
	if xmlDump != "<iso mti=\"0800\"><field id=\"Header\">ISO</field><field id=\"11\">123456</field></iso>" {
		t.Errorf("XML Dump is wrong: %s", xmlDump)
	}

}
