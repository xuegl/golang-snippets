package golang_snippets

import (
	"testing"
)

func TestTOTP_At(t1 *testing.T) {

}

func TestTOTP_Now(t1 *testing.T) {
	totp := NewDefaultTOTP("golang")
	t1.Log(totp.Now())
}
