package geoip

import (
	"testing"
)

func TestNewResolver(t *testing.T) {
	resolver, err := NewResolver("./testdata/GeoLite2-Country-Test.mmdb")
	if err != nil {
		t.Fatalf("Failed to create resolver: %v", err)
	}
	if resolver == nil {
		t.Error("Expected resolver to be non-nil")
	}
}
