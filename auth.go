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
	"encoding/base64"
	"net/http"
)

// A Authenticable defines rules for a type that offers HTTP authentication.
type Authenticable interface {
	TryAuthentication(r *http.Request, user, secret string) bool
}

// A Authenticator defines rules for a type that handles HTTP authentication.
type Authenticator interface {
	AuthHandler(http.Handler) http.Handler
}

// WwwAuthenticate creates a HTTP header to require client authentication.
func (HeaderBuilder) WwwAuthenticate() *Header {
	return &Header{
		"WWW-Authenticate",
		"", // type and params
	}
}

// Authorization creates a HTTP header to request authentication.
func (HeaderBuilder) Authorization(user, secret string) *Header {
	var value string

	if len(user)+len(secret) > 0 {
		credentials := []byte(user + ":" + secret)
		value = basicPrefix + base64.StdEncoding.EncodeToString(credentials)
	}

	return &Header{
		authHeaderName,
		value, // credentials
	}
}
