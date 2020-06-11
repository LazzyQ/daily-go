package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"net/url"
	"strings"
	"time"
)

type proxymw struct {
	ctx       context.Context
	next      StringService
	uppercase endpoint.Endpoint
}

func (mw proxymw) Count(s string) int {
	return mw.next.Count(s)
}

func (mw proxymw) Uppercase(s string) (string, error) {
	response, err := mw.uppercase(mw.ctx, uppercaseRequest{S: s})
	if err != nil {
		return "", err
	}
	resp := response.(uppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}



func proxyingMiddleware(ctx context.Context, instances string, logger log.Logger) ServiceMiddleware {
	if instances == "" {
		logger.Log("proxy_to", "none")
		return func(next StringService) StringService {
			return next
		}
	}

	// 一些客户堵的参数
	var (
		qps         = 100
		maxAttempts = 3
		maxTime     = 250 * time.Millisecond
	)

	// 为每个instance构造出endpoint
	// 在真的服务中，通常使用服务发现系统来实现
	var (
		instanceList = split(instances)
		endpointer   sd.FixedEndpointer
	)

	logger.Log("proxy_to", fmt.Sprint(instanceList))
	for _, instance := range instanceList {
		var e endpoint.Endpoint
		e = makeUppercaseProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	return func(next StringService) StringService {
		return proxymw{ctx, next, retry}
	}
}

func split(s string) []string {
	a := strings.Split(s, ",")
	for i, ii := range a {
		a[i] = strings.TrimSpace(ii)
	}
	return a
}

func makeUppercaseProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/uppercase"
	}
	return httptransport.NewClient(
		"Get",
		u,
		encodeRequest,
		decodeUppercaseResponse,
	).Endpoint()
}

