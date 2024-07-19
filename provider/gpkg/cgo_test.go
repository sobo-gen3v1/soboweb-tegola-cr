//go:build cgo
// +build cgo

package gpkg

import (
	"testing"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/dict"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/provider"
)

// This is a test to just see that the init function is doing something.
func TestNewProviderStartup(t *testing.T) {
	_, err := NewTileProvider(dict.Dict{}, nil)
	if err == provider.ErrUnsupported {
		t.Fatalf("supported, expected any but unsupported got %v", err)
	}
}
