package pem_test

import (
	"fmt"
	"github.com/snow/golang-snippets/crypto/x509"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadCertificate(t *testing.T) {
	certificate, err := x509.LoadCertificateWithPath(filepath.Join("testdata", "test_cert.pem"))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Issuer: %#v\n", certificate.Issuer)
	fmt.Printf("Subject: %#v\n", certificate.Subject)
	fmt.Printf("NotBefore: %#v\n", certificate.NotBefore.Format("2006-01-02 15:04:05"))
	fmt.Printf("NotAfter: %#v\n", certificate.NotAfter.Format("2006-01-02 15:04:05"))
	assert.True(t, x509.IsCertificateValid(certificate, time.Now()))
}

func TestLoadPublicKeyWithPath(t *testing.T) {
	publicKey, err := x509.LoadPublicKeyWithPath(filepath.Join("testdata", "test_pub.pem"))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("public key size: %d\n", publicKey.Size())
}

func TestLoadPrivateKey(t *testing.T) {
	privateKey, err := x509.LoadPrivateKeyWithPath(filepath.Join("testdata", "test_priv.pem"))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("private key: %p\n", privateKey)
}
