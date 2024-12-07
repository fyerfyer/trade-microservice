package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
	"trade-microservice.fyerfyer.net/internal/e"
)

func GetClientTLSCredentials() (credentials.TransportCredentials, error) {
	clientCert, _ := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")

	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("cert/ca.cert.pem")
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, e.FAILED_TLS_CERT
	}

	return credentials.NewTLS(&tls.Config{
		ServerName:   "*.microservices.dev",
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}), nil
}
