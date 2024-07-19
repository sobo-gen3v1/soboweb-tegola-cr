// +build !cgo

package gpkg

import "github.com/sobo-gen3v1/soboweb-tegola-cr/provider"

func NewTileProvider(config map[string]interface{}) (provider.Tiler, error) {
	return nil, provider.ErrUnsupported
}

func Cleanup() {}
