/*
 * Copyright (C) 2016 Fabrício Godoy <skarllot@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
 */

package web

import (
	"crypto/rand"

	"gopkg.in/raiqub/crypt.v0"
	"gopkg.in/raiqub/data.v0"
)

// A SessionStoreBuilder provides methods to build a new SessionStore.
type SessionStoreBuilder interface {
	// Build creates and returns a new SessionStore.
	Build() *SessionStore

	// Salter sets a custom salter to generate random tokens.
	Salter(*crypt.Salter) SessionStoreBuilder

	// SalterFast sets a new salter backed by system random source to generate
	// new tokens. Optionally can be specified a initial salt to first token.
	SalterFast([]byte) SessionStoreBuilder

	// SalterSecure sets a new salter backed by an aggreagation of system and
	// SSTDEG random source to generate new tokens. Optionally can be specified
	// a initial salt to first token.
	//
	// The token creation will take at least 200 nanoseconds, but could normally
	// take one millisecond. This set is built with security over performance.
	SalterSecure([]byte) SessionStoreBuilder

	// Store sets a custom Store to store sessions.
	Store(data.Store) SessionStoreBuilder
}

type ssb struct {
	store  data.Store
	salter *crypt.Salter
}

// NewSessionStore creates a new builder for SessionStore.
func NewSessionStore() SessionStoreBuilder {
	return &ssb{}
}

func (b *ssb) Build() *SessionStore {
	return &SessionStore{
		salter: b.salter,
		cache:  b.store,
	}
}

func (b *ssb) Salter(salter *crypt.Salter) SessionStoreBuilder {
	b.salter = salter
	return b
}

func (b *ssb) SalterFast(initialSalt []byte) SessionStoreBuilder {
	b.salter = crypt.NewSalter(rand.Reader, initialSalt)
	return b
}

func (b *ssb) SalterSecure(initialSalt []byte) SessionStoreBuilder {
	source := crypt.NewRandomAggr().SecureSet()
	b.salter = crypt.NewSalter(source, initialSalt)
	return b
}

func (b *ssb) Store(store data.Store) SessionStoreBuilder {
	b.store = store
	return b
}
