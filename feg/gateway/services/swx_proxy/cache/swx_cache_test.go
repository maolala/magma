/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/
package cache_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/services/swx_proxy"
	"magma/feg/gateway/services/swx_proxy/cache"
	"magma/feg/gateway/services/swx_proxy/servicers/test"
	"magma/feg/gateway/services/swx_proxy/test_init"
	orcprotos "magma/orc8r/cloud/go/protos"
)

func TestSwxCacheGC(t *testing.T) {
	err := test_init.InitTestMconfig(t, "127.0.0.1:0", true)
	assert.NoError(t, err)
	interval := time.Millisecond * 10
	ttl := time.Millisecond * 100
	cache, done := cache.NewExt(interval, ttl)
	srv, err := test_init.StartTestServiceWithCache(t, cache)
	if err != nil {
		t.Fatal(err)
	}

	authReq := &protos.AuthenticationRequest{
		UserName:             test.BASE_IMSI,
		SipNumAuthVectors:    uint32(3),
		AuthenticationScheme: protos.AuthenticationScheme_EAP_AKA,
	}

	authRes, err := swx_proxy.Authenticate(authReq)
	if err != nil {
		t.Fatalf("GRPC MAR Error: %v", err)
		return
	}
	assert.Equal(t, 1, len(authRes.GetSipAuthVectors()))
	v := authRes.SipAuthVectors[0]
	assert.Equal(t, protos.AuthenticationScheme_EAP_AKA, v.GetAuthenticationScheme())
	assert.Equal(t, []byte(test.DefaultSIPAuthenticate+strconv.Itoa(int(14))), v.GetRandAutn())
	assert.Equal(t, []byte(test.DefaultSIPAuthorization), v.GetXres())
	assert.Equal(t, []byte(test.DefaultCK), v.GetConfidentialityKey())
	assert.Equal(t, []byte(test.DefaultIK), v.GetIntegrityKey())

	authRes = cache.Get(test.BASE_IMSI)
	assert.Equal(t, 1, len(authRes.GetSipAuthVectors()))
	v = authRes.SipAuthVectors[0]
	assert.Equal(t, protos.AuthenticationScheme_EAP_AKA, v.GetAuthenticationScheme())
	assert.Equal(t, []byte(test.DefaultSIPAuthenticate+strconv.Itoa(int(15))), v.GetRandAutn())
	assert.Equal(t, []byte(test.DefaultSIPAuthorization), v.GetXres())
	assert.Equal(t, []byte(test.DefaultCK), v.GetConfidentialityKey())
	assert.Equal(t, []byte(test.DefaultIK), v.GetIntegrityKey())

	time.Sleep(ttl + interval*2)

	authRes = cache.Get(test.BASE_IMSI)
	assert.Equal(t, (*protos.AuthenticationAnswer)(nil), authRes)

	done <- struct{}{}

	_, err = srv.StopService(context.Background(), &orcprotos.Void{})
	assert.NoError(t, err)
}
