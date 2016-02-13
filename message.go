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
	"errors"
	"fmt"
	"strconv"
)

// Field represents a normal type message's field
type Field struct {
	ID    int
	Value string
}

// BitmapField represents a bitmap type message's field
type BitmapField struct {
	Value *Bitmap
}

// IsoMessage represents a ISO8583 message
type IsoMessage struct {
	header, mti Field
	bitmap      BitmapField
	fields      []Field
	maxID       int
}

// AddField adds a new field to a message
func (msg *IsoMessage) AddField(id int, value string) {

	if msg.HasField(id) {
		msg.RemoveField(id)
	}

	msg.fields = append(msg.fields, Field{ID: id, Value: value})
	if msg.maxID < id {
		msg.maxID = id
	}
}

// HasField checks if a field exists
func (msg *IsoMessage) HasField(id int) (b bool) {
	for i := 0; i < len(msg.fields); i++ {
		if msg.fields[i].ID == id {
			b = true
			return
		}
	}
	return
}

// RemoveField removes a field form the collection
func (msg *IsoMessage) RemoveField(id int) (b bool) {
	for i, f := range msg.fields {
		if f.ID == id {
			msg.fields = append(msg.fields[:i], msg.fields[i+1:]...)
			return
		}
	}
	return
}

// GetField returns a field from the collection
func (msg *IsoMessage) GetField(id int) (f Field, err error) {
	for i := 0; i < len(msg.fields); i++ {
		if msg.fields[i].ID == id {
			f = msg.fields[i]
			break
		}
	}

	if f.ID == 0 {
		err = errors.New("Field not found")
	}

	return
}

// GetValue returns the field's value
func (msg *IsoMessage) GetValue(id int) (value string, err error) {
	f, err := msg.GetField(id)
	if err == nil {
		value = f.Value
	}

	return
}

// GetInt returns the field's value as integer
func (msg *IsoMessage) GetInt(id int) (value int, err error) {
	f, err := msg.GetField(id)
	if err == nil {
		value, err = strconv.Atoi(f.Value)
	}
	return
}

// SetHeader sets the message's header
func (msg *IsoMessage) SetHeader(value string) {
	msg.header = Field{ID: FieldHeader, Value: value}
}

// SetMTI sets the message type
func (msg *IsoMessage) SetMTI(value string) {
	msg.mti = Field{ID: FieldMTI, Value: value}
}

// DumpXMLWithFormat returns a dump of the message's content to a formatted XML
func (msg *IsoMessage) DumpXMLWithFormat() (dump string) {

	dump += fmt.Sprintf("<iso mti=\"%s\">\n", msg.mti.Value)
	dump += fmt.Sprintf("\t<field id=\"%s\">%s</field>\n", "Header", msg.header.Value)

	if msg.bitmap.Value != nil {
		dump += fmt.Sprintf("\t<field id=\"%s\">%s</field>\n", "Bitmaps", msg.bitmap.Value.ToHexString())
	}

	for i := 1; i < len(msg.fields); i++ {
		dump += fmt.Sprintf("\t<field id=\"%d\">%s</field>\n", msg.fields[i].ID, msg.fields[i].Value)
	}

	dump += "</iso>"

	return
}

// DumpXML dumps the content to a one-lined XML
func (msg *IsoMessage) DumpXML() (dump string) {

	dump += fmt.Sprintf("<iso mti=\"%s\">", msg.mti.Value)
	dump += fmt.Sprintf("<field id=\"%s\">%s</field>", "Header", msg.header.Value)

	if msg.bitmap.Value != nil {
		dump += fmt.Sprintf("<field id=\"%s\">%s</field>", "Bitmaps", msg.bitmap.Value.ToHexString())
	}

	for i := 1; i < len(msg.fields); i++ {
		dump += fmt.Sprintf("<field id=\"%d\">%s</field>", msg.fields[i].ID, msg.fields[i].Value)
	}

	dump += "</iso>"

	return
}
