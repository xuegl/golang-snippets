# Generate private key
openssl genrsa -out test_priv.pem 2048

# Generate public key
openssl rsa -in test_priv.pem -pubout -out test_pub.pem

# Generate certificate signing request
openssl req -new -key test_priv.pem -out test_csr.pem

# Generate certificate
openssl x509 -req -days 3600 -in test_csr.pem -signkey test_priv.pem -out test_cert.pem
