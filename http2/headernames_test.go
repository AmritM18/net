// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"slices"
	"sort"
	"testing"
)

func TestHeaderNamesAddAndContains(t *testing.T) {
	h := &HeaderNames{}
	h.Add("Authorization", "X-Api-Key")

	if !h.contains("authorization") {
		t.Fatal("contains(authorization) = false, want true")
	}
	if !h.contains("x-api-key") {
		t.Fatal("contains(x-api-key) = false, want true")
	}
	if h.contains("x-normal") {
		t.Fatal("contains(x-normal) = true, want false")
	}
}

func TestHeaderNamesNilContains(t *testing.T) {
	var h *HeaderNames
	if h.contains("authorization") {
		t.Fatal("nil HeaderNames contains(authorization) = true, want false")
	}
}

func TestHeaderNamesDelete(t *testing.T) {
	h := &HeaderNames{}
	h.Add("Authorization", "X-Api-Key", "X-Normal")
	h.Delete("X-Api-Key")

	if !h.contains("authorization") {
		t.Fatal("contains(authorization) = false after deleting x-api-key")
	}
	if h.contains("x-api-key") {
		t.Fatal("contains(x-api-key) = true after delete")
	}
	if !h.contains("x-normal") {
		t.Fatal("contains(x-normal) = false after deleting x-api-key")
	}
}

func TestHeaderNamesAll(t *testing.T) {
	h := &HeaderNames{}
	h.Add("Authorization", "X-Api-Key")

	var got []string
	for name := range h.All() {
		got = append(got, name)
	}
	sort.Strings(got)

	want := []string{"authorization", "x-api-key"}
	if !slices.Equal(got, want) {
		t.Fatalf("All() = %v, want %v", got, want)
	}
}
