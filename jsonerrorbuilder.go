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
	"reflect"
)

// A JSONErrorBuilder provides methods to construct a new JSONError.
type JSONErrorBuilder interface {
	// Build creates and returns a new JSONError.
	Build() JSONError

	// CustomError sets current instance from specified error.
	CustomError(code int, errorType, msg string) JSONErrorBuilder

	// FromError sets current instance based on native error.
	FromError(e error) JSONErrorBuilder

	// Message sets a message for current instance.
	Message(string) JSONErrorBuilder

	// Status sets a HTTP status for current instance.
	Status(int) JSONErrorBuilder

	// URL sets a reference URL for current instance.
	URL(string) JSONErrorBuilder
}

type jsonErrorBuilder struct {
	instance JSONError
}

// NewJSONError creates a new instance of JSONErrorBuilder and sets the default
// HTTP status to Internal Server Error (500).
func NewJSONError() JSONErrorBuilder {
	return &jsonErrorBuilder{JSONError{
		Status: http.StatusInternalServerError,
	}}
}

func (b *jsonErrorBuilder) Build() JSONError {
	return b.instance
}

func (b *jsonErrorBuilder) CustomError(
	code int,
	errorType,
	msg string,
) JSONErrorBuilder {
	b.instance.Code = code
	b.instance.Type = errorType
	b.instance.Message = msg

	return b
}

func (b *jsonErrorBuilder) FromError(e error) JSONErrorBuilder {
	errType := reflect.TypeOf(e)
	var typeName string
	if errType.Kind() == reflect.Ptr {
		typeName = errType.Elem().Name()
	} else {
		typeName = errType.Name()
	}

	b.instance.Type = typeName
	b.instance.Message = e.Error()
	return b
}

func (b *jsonErrorBuilder) Message(msg string) JSONErrorBuilder {
	b.instance.Message = msg
	return b
}

func (b *jsonErrorBuilder) Status(status int) JSONErrorBuilder {
	b.instance.Status = status
	return b
}

func (b *jsonErrorBuilder) URL(url string) JSONErrorBuilder {
	b.instance.MoreInfo = url
	return b
}

var _ JSONErrorBuilder = (*jsonErrorBuilder)(nil)
