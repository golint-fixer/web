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
	"net/http/httptest"
	"strings"
	"testing"
)

type FooAuthenticator int

func (a *FooAuthenticator) TryAuthentication(
	r *http.Request,
	user, secret string,
) bool {
	if user == "user" && secret == "secret" {
		*a = FooAuthenticator(int(*a) + 1)
		return true
	}
	return false
}

func (a *FooAuthenticator) EndPoint(w http.ResponseWriter, r *http.Request) {
	*a = FooAuthenticator(int(*a) * 10)
}

func TestBasicAuthenticator(t *testing.T) {
	testValues := []struct {
		user   string
		secret string
	}{
		{"user", "user"},
		{"user", "secret"},
		{"user", "123"},
		{"user", ""},
		{"secret", "user"},
		{"secret", "secret"},
		{"secret", "123"},
		{"secret", ""},
		{"123", "user"},
		{"123", "secret"},
		{"123", "123"},
		{"123", ""},
		{"", "user"},
		{"", "secret"},
		{"", "123"},
		{"", ""},
	}
	bodyUnauthorized := http.StatusText(http.StatusUnauthorized)

	for idx, testVal := range testValues {
		foo := FooAuthenticator(1)
		basicauth := BasicAuthenticator{&foo}

		chain := NewChain()
		chain = append(chain, basicauth.AuthHandler)
		server := chain.Get(http.HandlerFunc(foo.EndPoint))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "http://localhost", nil)
		NewHeader().
			Authorization(testVal.user, testVal.secret).
			Write(req.Header)

		server.ServeHTTP(w, req)

		switch int(foo) {
		case 1:
			if idx == 1 {
				t.Error("Failed authentication: neither the middleware and endpoint was called")
			}
		case 2:
			t.Error("Failed authentication: endpoint was not called")
		case 10:
			t.Error("Failed authentication: enpoint was called without authentication")
		case 20:
			if idx != 1 {
				t.Errorf("Should not authenticate")
			}
		default:
			t.Errorf("Failed authentication: unexpected value %d", int(foo))
		}

		body := strings.TrimSpace(w.Body.String())
		if idx == 1 && len(body) > 0 {
			t.Errorf("Body is not empty: %s", w.Body.String())
		}
		if idx != 1 && body != bodyUnauthorized {
			t.Errorf("Body should be '%s' but got '%s'",
				bodyUnauthorized, body)
		}
	}
}
