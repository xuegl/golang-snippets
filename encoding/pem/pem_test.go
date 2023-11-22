package pem

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadCertificate(t *testing.T) {
	certificate, err := LoadCertificateWithPath("./resources/test_cert.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Issuer: %#v\n", certificate.Issuer)
	fmt.Printf("Subject: %#v\n", certificate.Subject)
	fmt.Printf("NotBefore: %#v\n", certificate.NotBefore.Format("2006-01-02 15:04:05"))
	fmt.Printf("NotAfter: %#v\n", certificate.NotAfter.Format("2006-01-02 15:04:05"))
	assert.True(t, IsCertificateValid(certificate, time.Now()))
}

func TestLoadPublicKeyWithPath(t *testing.T) {
	publicKey, err := LoadPublicKeyWithPath("./resources/test_pub.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("public key size: %d\n", publicKey.Size())
}

func TestLoadPrivateKey(t *testing.T) {
	privateKey, err := LoadPrivateKeyWithPath("./resources/test_priv.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("private key: %p\n", privateKey)
}
