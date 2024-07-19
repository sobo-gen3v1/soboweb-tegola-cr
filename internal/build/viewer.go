//go:build !noViewer
// +build !noViewer

package build

import (
	"github.com/sobo-gen3v1/soboweb-tegola-cr/ui"
)

func ViewerVersion() string {
	version := ui.Version()
	if version == "" {
		return uiVersionDefaultText
	}

	return version
}
