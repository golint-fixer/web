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

/*
Package web provides operations to help HTTP server implementation.

Chain

A Chain provides a function to chain HTTP handlers, also know as middlewares,
before a specified HTTP handler. A Chain is basically a slice of middlewares.

Header

A Header provides functions to help handle HTTP headers, both for reading from
client request and write to server response.

HeaderBuilder

A HeaderBuilder provides some pre-defined HTTP headers.

JSON

Provides a JSONRead and JSONWrite functions for easiest JSON communication, and
a JSONError struct which defines a format for JSON errors as defined by best
practices.

SessionStore

A SessionStore provides session tokens to uniquely identify an user session and
links it to specified data. Each token expires automatically if it is not used
after defined time.
*/
package web
