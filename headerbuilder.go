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

// A HeaderBuilder provides pre-defined HTTP headers.
type HeaderBuilder int

// NewHeader creates a new instance of HeaderBuilder.
func NewHeader() HeaderBuilder {
	return HeaderBuilder(0)
}

// AccessControlAllowCredentials creates a HTTP header to CORS-able API indicate
// that authentication is allowed.
func (HeaderBuilder) AccessControlAllowCredentials() *Header {
	return &Header{
		"Access-Control-Allow-Credentials",
		"", // boolean
	}
}

// AccessControlAllowHeaders creates a HTTP header to CORS-able API indicate
// which HTTP headers are allowed.
func (HeaderBuilder) AccessControlAllowHeaders() *Header {
	return &Header{
		"Access-Control-Allow-Headers",
		"", // comma-separated list
	}
}

// AccessControlAllowMethods creates a HTTP header to CORS-able API indicate
// which HTTP methods are allowed to current resource.
func (HeaderBuilder) AccessControlAllowMethods() *Header {
	return &Header{
		"Access-Control-Allow-Methods",
		"", // comma-separated list of HTTP methods
	}
}

// AccessControlAllowOrigin creates a HTTP header to CORS-able API indicate
// which origin is expected.
func (HeaderBuilder) AccessControlAllowOrigin() *Header {
	return &Header{
		"Access-Control-Allow-Origin",
		"", // http-formatted domain or asterisk to any
	}
}

// AccessControlMaxAge creates a HTTP header to CORS-able API indicate how long
// preflight results should be cached.
func (HeaderBuilder) AccessControlMaxAge() *Header {
	return &Header{
		"Access-Control-Max-Age",
		"", // seconds
	}
}

// AccessControlRequestHeaders creates a HTTP header to CORS-able client
// indicate which headers will be used for request.
func (HeaderBuilder) AccessControlRequestHeaders() *Header {
	return &Header{
		"Access-Control-Request-Headers",
		"", // comma-separated list of HTTP headers
	}
}

// AccessControlRequestMethod creates a HTTP header to CORS-able client indicate
// which HTTP method will be used for request.
func (HeaderBuilder) AccessControlRequestMethod() *Header {
	return &Header{
		"Access-Control-Request-Method",
		"", // HTTP method name
	}
}

// ContentType creates a HTTP header builder to define a content type.
func (HeaderBuilder) ContentType() HeaderContentTypeBuilder {
	return HeaderContentTypeBuilder(0)
}

// Empty creates a empty HTTP header.
func (HeaderBuilder) Empty() *Header {
	return &Header{
		"",
		"",
	}
}

// Location creates a HTTP header to define location of new object.
func (HeaderBuilder) Location() *Header {
	return &Header{
		"Location",
		"", // relative http location
	}
}

// Origin creates a HTTP header to CORS-able client indicate its address.
func (HeaderBuilder) Origin() *Header {
	return &Header{
		"Origin",
		"", // http-formatted domain
	}
}

// A HeaderContentTypeBuilder provides pre-defined Content Types HTTP headers.
type HeaderContentTypeBuilder int

// Empty creates a HTTP header to undefined content type.
func (HeaderContentTypeBuilder) Empty() *Header {
	return &Header{
		"Content-Type",
		"",
	}
}

// HTML creates a HTTP header to define HTML content type.
func (HeaderContentTypeBuilder) HTML() *Header {
	return HeaderContentTypeBuilder(0).
		Empty().
		SetValue("text/html; charset=utf-8")
}

// JSON creates a HTTP header to define JSON content type.
func (HeaderContentTypeBuilder) JSON() *Header {
	return HeaderContentTypeBuilder(0).
		Empty().
		SetValue("application/json; charset=utf-8")
}

// Text creates a HTTP header to define plain text content type.
func (HeaderContentTypeBuilder) Text() *Header {
	return HeaderContentTypeBuilder(0).
		Empty().
		SetValue("text/plain; charset=utf-8")
}

// XML creates a HTTP header to define XML content type.
func (HeaderContentTypeBuilder) XML() *Header {
	return HeaderContentTypeBuilder(0).
		Empty().
		SetValue("application/xml; charset=utf-8")
}
