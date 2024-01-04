package ratelimiter

import (
	"time"
)

type TokenBucket struct {
	cap    int64
	rate   int64
	tokens int64
	last   time.Time
}

func NewTokenBucket(cap, rate int64) *TokenBucket {
	return &TokenBucket{
		cap:    cap,
		rate:   rate,
		tokens: cap,
		last:   time.Now(),
	}
}

func (t *TokenBucket) TryAcquire() bool {
	now := time.Now()
	newTokens := now.Sub(t.last).Milliseconds() * t.rate
	t.tokens = min(t.tokens+newTokens, t.cap)
	t.last = now
	if t.tokens > 0 {
		t.tokens--
		return true
	} else {
		return false
	}
}
