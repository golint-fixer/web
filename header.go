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

package web

import (
	"net/http"
)

// A Header represents a key-value pair in a HTTP header.
type Header struct {
	// HTTP header field name.
	Name string
	// HTTP header field value.
	Value string
}

// Clone make a copy of current instance.
func (s Header) Clone() *Header {
	return &s
}

// Read gets HTTP header value, as defined by current instance, from Request
// Header and sets to current instance.
func (s *Header) Read(h http.Header) *Header {
	s.Value = h.Get(s.Name)
	return s
}

// SetName sets header name of current instance.
func (s *Header) SetName(name string) *Header {
	s.Name = name
	return s
}

// SetValue sets header value of current instance.
func (s *Header) SetValue(value string) *Header {
	s.Value = value
	return s
}

// Write sets HTTP header, as defined by current instance, to ResponseWriter
// Header.
func (s *Header) Write(h http.Header) *Header {
	h.Set(s.Name, s.Value)
	return s
}
