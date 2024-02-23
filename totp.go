package golang_snippets

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/snow/golang-snippets/util"
	"hash"
	"math"
	"time"
)

type TOTPAlgorithmConfig struct {
	ReturnDigitalSize int   // the returned digit value length
	T0                int64 // the To
	TimeStep          int64 // the X
}

type TOTP struct {
	secret string
	cfg    TOTPAlgorithmConfig
	hash   hash.Hash
}

func NewDefaultTOTP(secret string) *TOTP {
	h := hmac.New(sha256.New, []byte(secret))
	return &TOTP{
		secret: secret,
		cfg: TOTPAlgorithmConfig{
			ReturnDigitalSize: 6,
			T0:                0,
			TimeStep:          30,
		},
		hash: h,
	}
}

func (t *TOTP) Now() string {
	return t.At(time.Now().Unix())
}

func (t *TOTP) At(ts int64) string {
	t.hash.Reset()
	t.hash.Write(util.Int64toB(t.steps(ts)))
	sum := t.hash.Sum(nil)
	offset := sum[len(sum)-1] & 0x0F // offset is the lower 4 bits of the last byte
	code := ((int(sum[offset]) & 0x7F) << 24) |
		((int(sum[offset+1]) & 0xFF) << 16) |
		((int(sum[offset+2]) & 0xFF) << 8) |
		(int(sum[offset+3]) & 0xFF)
	code = code % int(math.Pow10(t.cfg.ReturnDigitalSize))
	format := fmt.Sprintf("%%0%dd", t.cfg.ReturnDigitalSize)
	return fmt.Sprintf(format, code)
}

func (t *TOTP) steps(ts int64) int64 {
	return (ts - t.cfg.T0) / t.cfg.TimeStep
}
