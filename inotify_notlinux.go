// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !linux

package inotify // import "github.com/sigma/go-inotify"

import (
	"errors"
)

// NewWatcher creates and returns a new inotify instance using inotify_init(2)
func NewWatcher() (*Watcher, error) {
	return nil, errors.New("inotify is only supported on linux")
}
