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
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Foo struct {
	Number    int
	BigNumber int64
	Fraction  float32
	Text      string
	Boolean   bool
}

func (f Foo) IsEqual(other Foo) bool {
	return f.Number == other.Number &&
		f.BigNumber == other.BigNumber &&
		f.Fraction == other.Fraction &&
		f.Text == other.Text &&
		f.Boolean == other.Boolean
}

type UnbufferedResponse struct {
	Code      int
	HeaderMap http.Header
	Body      io.Writer
	Flushed   bool

	wroteHeader bool
}

func NewUnbufferedResponse(w io.Writer) *UnbufferedResponse {
	return &UnbufferedResponse{
		HeaderMap: make(http.Header),
		Body:      w,
		Code:      200,
	}
}

func (ur *UnbufferedResponse) Header() http.Header {
	m := ur.HeaderMap
	if m == nil {
		m = make(http.Header)
		ur.HeaderMap = m
	}
	return m
}

func (ur *UnbufferedResponse) Write(buf []byte) (int, error) {
	if !ur.wroteHeader {
		ur.WriteHeader(200)
	}
	if ur.Body != nil {
		return ur.Body.Write(buf)
	}
	return 0, nil
}

func (ur *UnbufferedResponse) WriteHeader(code int) {
	if !ur.wroteHeader {
		ur.Code = code
	}
	ur.wroteHeader = true
}

func TestJSONRead(t *testing.T) {
	foo := Foo{
		40,
		math.MaxInt64,
		93.09476283495,
		"Lorem ipsum",
		true,
	}

	r, w := io.Pipe()
	go func() {
		defer w.Close()

		err := json.NewEncoder(w).Encode(&foo)
		if err != nil {
			t.Errorf("Error encoding object: %v", err)
		}
	}()

	resp := httptest.NewRecorder()

	var fooCopy Foo
	JSONRead(r, &fooCopy, resp)
	if !foo.IsEqual(fooCopy) {
		t.Errorf(
			"Encoded object is not equal to decoded one."+
				"\n\nOriginal: %#v\nDecoded: %#v", foo, fooCopy)
	}

	if resp.Body.Len() > 0 {
		t.Errorf("Body is not empty: %s", resp.Body.String())
	}
}

func TestJSONWrite(t *testing.T) {
	foo := Foo{
		40,
		math.MaxInt64,
		93.09476283495,
		"Lorem ipsum",
		true,
	}

	r, w := io.Pipe()
	defer r.Close()

	go func() {
		defer w.Close()
		err := JSONWrite(NewUnbufferedResponse(w), http.StatusOK, foo)
		if err != nil {
			t.Errorf("Error on JSONWrite: %v", err)
		}
	}()

	var fooCopy Foo
	if err := json.NewDecoder(r).Decode(&fooCopy); err != nil {
		t.Errorf("Error on decoding object: %v", err)
	}

	if !foo.IsEqual(fooCopy) {
		t.Errorf(
			"Decoded object is not equal to encoded one."+
				"\n\nOriginal: %#v\nDecoded: %#v", foo, fooCopy)
	}
}
