package node

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"time"
)

const DefaultCAKey = `
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDPP8vbX1mNvkIg
qdSSl3aO6UYL6PcT3spSK45qUIf3B2qCkmL2cGotxzA8JIVcV7DoI0EgEz4JbaNP
dRXCDON1cnumC27yBaPQ1Do2pIuJdXZZulmi9+tWHfHtnhuM+Ajei5aPJ6FR+YrU
Y8iV3gF9RJC0+VQnnf5dPfYZGLNTL5vQJwnziOnVe/JkAxqNidyUVE00lLD1Ze1O
UxS75nUzgrtzhk/ASqAScKBFzRMJpePXzS+juYIclCgOI1IJEBPB8f6RbE44r+d4
+Sra3iY9IVpfzEBLjIgshIg5Zifuw2B3L0PqAY0WjZJnQVHfH2GNSNgRXdw+qDcG
Y0TcMPKdAgMBAAECggEAa6lLgEl3HyAQACHZUNGoADOEdNlvyP26gpcn42i0SQqs
NOpQyI67Sc6o6wVZ1g+j0ePGiCAW4RT4emVriSPi4Xc4bpiP6OAvKmOlXg96gUzo
z1H0EKnTsifaLsMssr2C9gDzlKhUsF3+1biEUf5DLcz5k1nWcsIrikqO1pizR2mL
nMuhW+kRwR6bY75pGlQ8JmqQ3BTS6YBMJ+BkZ8XOBcj9c8mvxTL4tVk4Wnc9dcd+
dlkP1PkCM79WIF6pswW5ZDwGlU6pqk08qPzENjei/De6IBg8Lt4uJve79agV6Uxz
Y4xfRCWMunlj4VJwUcAzGhyjcBB1ZJi9JDBvnFmaDQKBgQDrMoBJ+vDcS0+2GQtH
k0GZ6wjhFttjqIVWFLmWgtvNO1K2+y6zwQjpFCAuAAUXXgfbVhIfFvp0HUxujjkE
3zeCYOlS8o9ESyWoPdg8woWjkZcya/4mE2jlj1+GI4SfHMska0EHUd4aaDk1en2a
vYKH27zD+yV+4IdyRsD0Bi93EwKBgQDhlHnODvf5QDVro+b3dsthkaP6F/s5oH/o
eLD1jEIGAN9Yz10njyLPfzaeiNPm6LSsKVXlgCjxQIuHTU2n0BsFke4F1ib0jOK4
o3NN4rfVRb5O9vP50THv4+VhB9806FwiB371JeOhmitCSwiYp0wZmV3WURF4e4Fr
Cr3YlC+1jwKBgQC03BTCzvFAtbkKMp/13krn7VDaphT2wbQmybEdCGu1mhS1GNqE
57/OW+eS9/jySyCHjdxJhAX8HDuWGE/Ia03oOFWzr0p0HcVLZqNNtdfGPEKkR18c
MHjNbj7qi42EPUQJMWDEHDRK4jJ76UGFKI2jo1m46vueYVJGkhn2jHsbeQKBgDim
/Ew20Col6QSmfhwKFpvjYsYtfaeEWns8zFxupCozz+PS+Dc2KGzqKwJ3pJgqOy29
l9fybtXf+uq5DFan2hF1C80lclUaiNoMGqol1TtXr6rPNIi59AumNXY/7tuvu2vE
bCsPH/L28ARPKdKEuYT4Umu/ol6aze7fHLymwrCbAoGAbm4vMLIy8wHYrAj5/een
YthUYbQ+XWOrrI1dJoGiwCIXTL4iBSO1ZjkR6M7yg9Zxv0BuqA+2tZqImLXsqIq/
g7Jtrn7LfogS3KOiVa3QAHoXegMrRf4S3DyP/GW8E27MMvqBEkdCDX53h07wvPng
0NRspbyMaX43/tp8ZWkaZf0=
-----END PRIVATE KEY-----
`

const DefaultCACertificate = `
-----BEGIN CERTIFICATE-----
MIIDKTCCAhGgAwIBAgIUXIpxI5MoZQ65/vhc7DK/d5ymoMUwDQYJKoZIhvcNAQEL
BQAwRDENMAsGA1UECgwEQ2hpYTEQMA4GA1UEAwwHQ2hpYSBDQTEhMB8GA1UECwwY
T3JnYW5pYyBGYXJtaW5nIERpdmlzaW9uMB4XDTIxMDEyMzA4NTEwNloXDTMxMDEy
MTA4NTEwNlowRDENMAsGA1UECgwEQ2hpYTEQMA4GA1UEAwwHQ2hpYSBDQTEhMB8G
A1UECwwYT3JnYW5pYyBGYXJtaW5nIERpdmlzaW9uMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAzz/L219Zjb5CIKnUkpd2julGC+j3E97KUiuOalCH9wdq
gpJi9nBqLccwPCSFXFew6CNBIBM+CW2jT3UVwgzjdXJ7pgtu8gWj0NQ6NqSLiXV2
WbpZovfrVh3x7Z4bjPgI3ouWjyehUfmK1GPIld4BfUSQtPlUJ53+XT32GRizUy+b
0CcJ84jp1XvyZAMajYnclFRNNJSw9WXtTlMUu+Z1M4K7c4ZPwEqgEnCgRc0TCaXj
180vo7mCHJQoDiNSCRATwfH+kWxOOK/nePkq2t4mPSFaX8xAS4yILISIOWYn7sNg
dy9D6gGNFo2SZ0FR3x9hjUjYEV3cPqg3BmNE3DDynQIDAQABoxMwETAPBgNVHRMB
Af8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAEugnFQjzHhS0eeCqUwOHmP3ww
/rXPkKF+bJ6uiQgXZl+B5W3m3zaKimJeyatmuN+5ST1gUET+boMhbA/7grXAsRsk
SFTHG0T9CWfPiuimVmGCzoxLGpWDMJcHZncpQZ72dcy3h7mjWS+U59uyRVHeiprE
hvSyoNSYmfvh7vplRKS1wYeA119LL5fRXvOQNW6pSsts17auu38HWQGagSIAd1UP
5zEvDS1HgvaU1E09hlHzlpdSdNkAx7si0DMzxKHUg9oXeRZedt6kcfyEmryd52Mj
1r1R9mf4iMIUv1zc2sHVc1omxnCw9+7U4GMWLtL5OgyJyfNyoxk3tC+D3KNU
-----END CERTIFICATE-----
`

var (
	DefaultTLSConfig = &tls.Config{
		ServerName:   "chia.net",
		Certificates: []tls.Certificate{},
		RootCAs:      &x509.CertPool{},
		MinVersion:   tls.VersionTLS12,
	}
)

func GenerateCASignedCert(caCertPem []byte, caKeyPem []byte) (rCert []byte, rKey []byte, err error) {

	var ca tls.Certificate
	var caCert *x509.Certificate

	if ca, err = tls.X509KeyPair(caCertPem, caKeyPem); err != nil {
		return nil, nil, err
	}

	var certDERBlock *pem.Block

	if certDERBlock, caCertPem = pem.Decode(caCertPem); certDERBlock == nil {
		return nil, nil, errors.New("failed to decode CA certificate")
	}

	if caCert, err = x509.ParseCertificate(certDERBlock.Bytes); err != nil {
		return nil, nil, err
	}

	var certPrivKey *rsa.PrivateKey

	if certPrivKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		return nil, nil, err
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:       []string{"Chia"},
			OrganizationalUnit: []string{"Organic Farming Division"},
		},
		DNSNames:     []string{"chia.net"},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now().Add(-24 * time.Hour),
		NotAfter:     time.Now().AddDate(80, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, &certPrivKey.PublicKey, ca.PrivateKey)

	if err != nil {
		return nil, nil, err
	}

	certPEM := new(bytes.Buffer)

	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM := new(bytes.Buffer)

	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})

	return certPEM.Bytes(), certPrivKeyPEM.Bytes(), nil
}
