//go:build go1.9
// +build go1.9

package executor

import "sync"

type Map struct {
	sync.Map
}

func NewMap() *Map {
	return &Map{}
}
