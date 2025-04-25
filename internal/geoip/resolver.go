package geoip

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

type Resolver struct {
	db *geoip2.Reader
}

func NewResolver(dbPath string) (*Resolver, error) {
	db, err := geoip2.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &Resolver{db: db}, nil
}

func (r *Resolver) Country(ipStr string) (string, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", fmt.Errorf("invalid IP format")
	}
	record, err := r.db.Country(ip)
	if err != nil {
		return "", err
	}
	return record.Country.IsoCode, nil
}
