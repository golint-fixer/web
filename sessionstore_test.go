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
	"testing"
	"time"

	"gopkg.in/raiqub/data.v0/memstore"
)

const TokenSalt = "CvoTVwDw685Ve0qjGn//zmHGKvoCcslYNQT4AQ9FygSk9t6NuzBHuohyO" +
	"Hhqb/1omn6c"

func TestSessionLifetime(t *testing.T) {
	store := memstore.New(time.Millisecond*10, false)
	ts := NewSessionStore().
		SalterFast([]byte(TokenSalt)).
		Store(store).
		Build()

	t1, err := ts.Add(nil)
	if err != nil {
		t.Errorf("The session t1 could not be generated: %v", err)
	}
	t2, err := ts.Add(nil)
	if err != nil {
		t.Errorf("The session t2 could not be generated: %v", err)
	}

	if err := ts.Get(t1, nil); err != nil {
		t.Error("The session t1 was not stored")
	}
	if err := ts.Get(t2, nil); err != nil {
		t.Error("The session t2 was not stored")
	}

	time.Sleep(time.Millisecond * 20)

	if err := ts.Get(t1, nil); err == nil {
		t.Error("The session t1 was not expired")
	}
	if err := ts.Get(t2, nil); err == nil {
		t.Error("The session t2 was not expired")
	}

	if err := ts.Delete(t1); err == nil {
		t.Error("The expired session t1 should not be removable")
	}
	if err := ts.Set(t2, nil); err == nil {
		t.Error("The expired session t2 should not be changeable")
	}
}

func TestSessionHandling(t *testing.T) {
	testValues := []struct {
		ref   string
		token string
		value int
	}{
		{"t1", "", 3},
		{"t2", "", 6},
		{"t3", "", 83679},
		{"t4", "", 2748},
		{"t5", "", 54},
		{"t6", "", 6},
		{"t7", "", 2},
		{"t8", "", 8},
		{"t9", "", 7},
		{"t10", "", 8},
	}
	rmTestIndex := 6
	changeValues := map[int]int{
		2: 5062,
		9: 4099,
	}

	store := memstore.New(time.Millisecond*100, false)
	ts := NewSessionStore().
		SalterFast([]byte(TokenSalt)).
		Store(store).
		Build()
	if _, err := ts.Count(); err != nil {
		t.Fatal("The Count() method should be supported by MemStore")
		return
	}

	if count, _ := ts.Count(); count != 0 {
		t.Errorf(
			"The session cache should be empty, but it has %d items",
			count)
	}

	lastCount, _ := ts.Count()
	for i := range testValues {
		var err error
		item := &testValues[i]
		item.token, err = ts.Add(nil)
		if err != nil {
			t.Errorf("A new session could not be generated: %v", err)
		}

		if count, _ := ts.Count(); count != lastCount+1 {
			t.Errorf(
				"The new session '%s' was not stored into session cache",
				item.token)
		}
		lastCount, _ = ts.Count()

		err = ts.Set(item.token, item.value)
		if err != nil {
			t.Errorf("The session %s could not be set", item.ref)
		}
	}

	if count, _ := ts.Count(); count != len(testValues) {
		t.Errorf("The session count do not match (%d!=%d)",
			count, len(testValues))
	}

	for _, i := range testValues {
		var v int
		if err := ts.Get(i.token, &v); err != nil {
			t.Errorf("The session %s could not be read", i.ref)
		}
		if v != i.value {
			t.Errorf("The session %s was stored incorrectly", i.ref)
		}
	}

	rmTestKey := testValues[rmTestIndex]
	if err := ts.Delete(rmTestKey.token); err != nil {
		t.Errorf("The session %s could not be removed", rmTestKey.ref)
	}
	if err := ts.Get(rmTestKey.token, nil); err == nil {
		t.Errorf("The removed session %s should not be retrieved", rmTestKey.ref)
	}
	if count, _ := ts.Count(); count == len(testValues) {
		t.Error("The session count should not match")
	}

	for k, v := range changeValues {
		item := testValues[k]
		err := ts.Set(item.token, v)
		if err != nil {
			t.Errorf("The session %s could not be changed", item.ref)
		}
	}
	for k, v := range changeValues {
		item := testValues[k]
		var v2 int
		if err := ts.Get(item.token, &v2); err != nil {
			t.Errorf("The session %s could not be read", item.ref)
		}
		if v2 != v {
			t.Errorf("The session %s was not changed", item.ref)
		}
	}
}

func BenchmarkSessionCreation(b *testing.B) {
	store := memstore.New(time.Millisecond, false)
	ts := NewSessionStore().
		SalterSecure([]byte(TokenSalt)).
		Store(store).
		Build()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ts.Add(nil)
	}
}

func BenchmarkSessionCreationFast(b *testing.B) {
	store := memstore.New(time.Millisecond, false)
	ts := NewSessionStore().
		SalterFast([]byte(TokenSalt)).
		Store(store).
		Build()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ts.Add(nil)
	}
}
