package ratelimiter

import (
	"sync/atomic"
	"time"
)

type RateLimiterSimpleWindow struct {
	qps    int64
	window time.Duration
	reqCnt atomic.Int64
	start  time.Time
}

func NewRateLimiterSimpleWindow(qps int64, window time.Duration) *RateLimiterSimpleWindow {
	return &RateLimiterSimpleWindow{
		qps:    qps,
		window: window,
		start:  time.Now(),
	}
}

func (r *RateLimiterSimpleWindow) TryAcquire() bool {
	if time.Now().Sub(r.start) > r.window {
		r.reqCnt.Store(0)
		r.start = time.Now()
	}
	return r.reqCnt.Load() <= r.qps
}
