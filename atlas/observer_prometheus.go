// +build !noPrometheusObserver

package atlas

// The point of this file is to load and register the prometheus observer backend.
// The prometheus observer can be excluded during the build with the `noPrometheusObserver` build flag
import (
	_ "github.com/sobo-gen3v1/soboweb-tegola-cr/observability/prometheus"
)
