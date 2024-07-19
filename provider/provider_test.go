package provider_test

import (
	"testing"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/provider"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/provider/test"
)

func TestProviderInterface(t *testing.T) {
	var (
		stdName = provider.TypeStd.Prefix() + test.Name
		mvtName = provider.TypeMvt.Prefix() + test.Name
	)
	if _, err := provider.For(stdName, nil, nil); err != nil {
		t.Errorf("retrieve provider err , expected nil got %v", err)
		return
	}
	if test.Count != 1 {
		t.Errorf(" expected count , expected 1 got %v", test.Count)
	}
	provider.Cleanup()
	if test.Count != 0 {
		t.Errorf(" expected count , expected 0 got %v", test.Count)
	}
	if _, err := provider.For(mvtName, nil, nil); err != nil {
		t.Errorf("retrieve provider err , expected nil got %v", err)
		return
	}
	if test.MVTCount != 1 {
		t.Errorf(" expected count , expected 1 got %v", test.MVTCount)
	}
	provider.Cleanup()
	if test.MVTCount != 0 {
		t.Errorf(" expected count , expected 0 got %v", test.MVTCount)
	}
}
