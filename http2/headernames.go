// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"iter"
	"strings"
)

// HeaderNames is an unordered set of header names.
// Its entries are automatically canonicalized to lowercase.
type HeaderNames struct {
	m map[string]struct{}
}

func (h *HeaderNames) Add(names ...string) {
	if h.m == nil {
		h.m = make(map[string]struct{})
	}
	for _, n := range names {
		h.m[strings.ToLower(n)] = struct{}{}
	}
}

func (h *HeaderNames) Delete(names ...string) {
	for _, n := range names {
		delete(h.m, strings.ToLower(n))
	}
}

func (h *HeaderNames) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		for n := range h.m {
			if !yield(n) {
				return
			}
		}
	}
}

func (h *HeaderNames) contains(name string) bool {
	if h == nil {
		return false
	}
	_, ok := h.m[name]
	return ok
}
