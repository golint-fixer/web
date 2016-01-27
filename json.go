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
	"encoding/json"
	"io"
	"net/http"
)

const (
	// StatusUnprocessableEntity defines WebDAV status; RFC 4918
	StatusUnprocessableEntity = 422
)

// JSONWrite sets response content type to JSON, sets HTTP status and serializes
// defined content to JSON format.
func JSONWrite(w http.ResponseWriter, status int, content interface{}) error {
	NewHeader().ContentType().JSON().Write(w.Header())
	w.WriteHeader(status)
	if content != nil {
		err := json.NewEncoder(w).Encode(content)
		if err != nil {
			return err
		}
	}

	return nil
}

// JSONRead tries to read client sent content using JSON decoding and
// writes it to object pointed to by obj.
//
// Returns true whether no error occurred; otherwise, false.
//
// Body is automatically closed when true is returned.
func JSONRead(
	body io.ReadCloser,
	maxlen int64,
	obj interface{},
	w http.ResponseWriter,
) bool {
	if err := json.
		NewDecoder(io.LimitReader(body, maxlen)).
		Decode(obj); err != nil {

		jerr := NewJSONError().
			FromError(err).
			Status(StatusUnprocessableEntity).
			Build()
		JSONWrite(w, jerr.Status, jerr)
		return false
	}

	if err := body.Close(); err != nil {
		jerr := NewJSONError().
			FromError(err).
			Build()
		JSONWrite(w, jerr.Status, jerr)
		return false
	}

	return true
}
