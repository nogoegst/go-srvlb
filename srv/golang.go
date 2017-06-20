// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package srv

import (
	"errors"
	"net"
	"strconv"
	"sync"
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
	_, srvs, err := net.LookupSRV("", "", domainName)
	if err != nil {
		return nil, err
	}
	ret := []*Target{}
	addrch := make(chan string)
	var wg sync.WaitGroup
	for _, s := range srvs {
		wg.Add(1)
		go func(s *net.SRV) {
			addrs, err := net.LookupHost(s.Target)
			if err != nil {
				return
			}
			if len(addrs) == 0 {
				return
			}
			addrch <- net.JoinHostPort(addrs[0], strconv.FormatUint(uint64(s.Port), 10))
			wg.Done()
		}(s)
	}
	go func() {
		wg.Wait()
		close(addrch)
	}()
	for da := range addrch {
		ret = append(ret, &Target{Ttl: r.ttl, DialAddr: da})
	}
	if len(ret) == 0 {
		return nil, errors.New("failed resolving hostnames for SRV entries")
	}
	return ret, nil
}
