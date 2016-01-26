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

var stacker []int

func TestChainOrder(t *testing.T) {
	stacker = make([]int, 0)

	chain := NewChain()
	chain = append(chain, FooHandler(1).Middleware)
	chain = append(chain, FooHandler(2).Middleware)
	chain = append(chain, FooHandler(3).Middleware)
	chain = append(chain, FooHandler(4).Middleware)
	chain = append(chain, FooHandler(5).Middleware)
	handler := chain.Get(http.HandlerFunc(FooHandler(6).EndPoint))
	handler.ServeHTTP(nil, nil)

	if len(stacker) != 6 {
		t.Errorf("Not all handlers was called: %d instead of %d",
			len(stacker), 6)
	}

	counter := 1
	for _, v := range stacker {
		if v != counter {
			t.Errorf("Chain not called in order: got %d instead of %d",
				v, counter)
		}

		counter++
	}
}

type FooHandler int

func (h FooHandler) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		stacker = append(stacker, int(h))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h FooHandler) EndPoint(w http.ResponseWriter, r *http.Request) {
	stacker = append(stacker, int(h))
}
