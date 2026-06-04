//go:build go1.18
// +build go1.18

package roaring

import (
	"bytes"
	"testing"
)

func FuzzSerializationBufferNative(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		newrb := NewBitmap()
		_, _ = newrb.FromBuffer(data)
	})
}

func FuzzSerializationStreamNative(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		newrb := NewBitmap()
		_, _ = newrb.ReadFrom(bytes.NewReader(data))
	})
}
