/*
 * Copyright 2016 Fabrício Godoy
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
	"testing"
)

func TestHeader(t *testing.T) {
	testingValues := []struct {
		name  string
		value string
	}{
		{"Testing-Name", "TestingValue"},
		{"Another-Name", "AnotherValue"},
	}
	h1 := NewHeader().Empty()

	if h1.Name != "" {
		t.Errorf("The header name should be empty but got '%s'", h1.Name)
	}
	if h1.Value != "" {
		t.Errorf("The header value should be empty but got '%s'", h1.Value)
	}

	h2 := h1.Clone()
	h1.SetName(testingValues[0].name)

	if h1.Name != testingValues[0].name {
		t.Errorf("The header name should be '%s' but got '%s'",
			testingValues[0].name, h1.Name)
	}
	if h2.Name != "" {
		t.Errorf("The cloned header name should be empty but got '%s'",
			h2.Name)
	}

	h1.SetValue(testingValues[0].value)

	if h1.Value != testingValues[0].value {
		t.Errorf("The header value should be '%s' but got '%s'",
			testingValues[0].value, h1.Value)
	}
	if h2.Value != "" {
		t.Errorf("The cloned header value should be empty but got '%s'",
			h2.Value)
	}

	httpHeader := make(map[string][]string)
	h1.Write(http.Header(httpHeader))

	if len(httpHeader) != 1 {
		t.Errorf("More than one key was written: %v", httpHeader)
	}

	val, ok := httpHeader[h1.Name]
	if !ok {
		t.Fatalf("Header was not written: %v", httpHeader)
	}

	if len(val) != 1 {
		t.Fatalf("Unexpected header value: %v", val)
	}
	if val[0] != h1.Value {
		t.Errorf("Header value written does not match: %s", val[0])
	}

	http.
		Header(httpHeader).
		Add(testingValues[1].name, testingValues[1].value)
	h2.
		SetName(testingValues[1].name).
		Read(http.Header(httpHeader))

	if h2.Value != testingValues[1].value {
		t.Errorf("Unexpected value read: %s", h2.Value)
	}

	if len(httpHeader) != 2 {
		t.Errorf("Unexpected headers length: %d", len(httpHeader))
	}
	val, ok = httpHeader[h2.Name]
	if !ok {
		t.Fatalf("Header should not be modified: %v", httpHeader)
	}
	if len(val) != 1 {
		t.Fatalf("Unexpected header value: %v", val)
	}
	if val[0] != h2.Value {
		t.Errorf("Headers value modified: %s", val[0])
	}
}
