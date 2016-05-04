/*
 * Copyright 2015 Fabrício Godoy
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
 */

//go:generate ffjson $GOFILE

package web

import (
	"bytes"
)

// A JSONError represents an error returned by JSON-based API.
type JSONError struct {
	// HTTP status code.
	Status int `json:"status,omitempty"`
	// Error code.
	Code int `json:"code,omitempty"`
	// Error type.
	Type string `json:"type,omitempty"`
	// A message with error details.
	Message string `json:"message,omitempty"`
	// A URL for reference.
	MoreInfo string `json:"moreInfo,omitempty"`
}

// Error returns string representation of current instance error.
func (e *JSONError) Error() string {
	var buf bytes.Buffer
	if len(e.Type) > 0 {
		buf.WriteString(e.Type)
		buf.WriteString(": ")
	}
	buf.WriteString(e.Message)
	if len(e.MoreInfo) > 0 {
		buf.WriteString(" (")
		buf.WriteString(e.MoreInfo)
		buf.WriteString(")")
	}

	return buf.String()
}

// String returns string representation of current instance.
func (e *JSONError) String() string {
	return e.Error()
}
