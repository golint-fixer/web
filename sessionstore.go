/*
 * Copyright (C) 2015 Fabrício Godoy <skarllot@gmail.com>
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
	"gopkg.in/raiqub/crypt.v0"
	"gopkg.in/raiqub/data.v0"
	"gopkg.in/raiqub/dot.v1"
)

// A SessionStore provides a temporary token to uniquely identify an user
// session.
type SessionStore struct {
	cache  data.Store
	salter *crypt.Salter
}

// Count gets the number of tokens stored by current instance.
//
// Errors:
// NotSupportedError when current method is not supported by store.
func (s *SessionStore) Count() (int, error) {
	return s.cache.Count()
}

// Get gets the value stored by specified token.
//
// Errors:
// InvalidTokenError when requested token could not be found.
func (s *SessionStore) Get(token string, ref interface{}) error {
	err := s.cache.Get(token, ref)
	if _, ok := err.(dot.InvalidKeyError); ok {
		return InvalidTokenError(token)
	}

	return err
}

// Add creates a new unique token and stores it into current SessionCache
// instance.
func (s *SessionStore) Add() string {
	strSum, err := s.salter.Token(0)
	if err != nil {
		panic("Could not generate a new token")
	}

	err = s.cache.Add(strSum, nil)
	if err != nil {
		panic("Something is seriously wrong, a duplicated token was generated")
	}

	return strSum
}

// Delete deletes specified token from current SessionCache instance.
//
// Errors:
// InvalidTokenError when requested token could not be found.
func (s *SessionStore) Delete(token string) error {
	err := s.cache.Delete(token)
	if err != nil {
		return InvalidTokenError(token)
	}
	return nil
}

// Set store a value to specified token.
//
// Errors:
// InvalidTokenError when requested token could not be found.
func (s *SessionStore) Set(token string, value interface{}) error {
	err := s.cache.Set(token, value)
	if err != nil {
		return InvalidTokenError(token)
	}
	return nil
}
