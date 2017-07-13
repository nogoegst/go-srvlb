// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package srv

import (
	"net"
	"strconv"
	"strings"
	"time"
)

// NewGoResolver is a resolver that uses net-package LookupSRV.
// It doesn't support TTL expiration and as such returns a dummy TTL.
// It's also very inefficient for large targets.
func NewGoResolver(dummyTtl time.Duration) Resolver {
	return &golangResolver{ttl: dummyTtl}
}

type golangResolver struct {
	ttl time.Duration
}

func (r *golangResolver) Lookup(domainName string) ([]*Target, error) {
	if !strings.HasPrefix(domainName, "_") { // non-SRV
		target := &Target{
			Ttl:      r.ttl,
			DialAddr: domainName,
		}
		return []*Target{target}, nil
	}
	_, srvs, err := net.LookupSRV("", "", domainName)
	if err != nil {
		return nil, err
	}
	ret := []*Target{}
	for _, s := range srvs {
		target := &Target{
			Ttl:      r.ttl,
			DialAddr: net.JoinHostPort(s.Target, strconv.FormatUint(uint64(s.Port), 10)),
		}
		ret = append(ret, target)
	}
	return ret, nil
}
