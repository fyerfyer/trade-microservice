package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
	"trade-microservice.fyerfyer.net/internal/e"
)

func GetServerTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, _ := tls.LoadX509KeyPair("cert/server.cert.pem", "cert/server-key.pem")

	// handle err
	certPool := x509.NewCertPool()
	caCert, _ := ioutil.ReadFile("cert/ca-key.pem")

	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, e.FAILED_TLS_CERT
	}

	return credentials.NewTLS(
		&tls.Config{
			ClientAuth:   tls.RequireAnyClientCert,
			Certificates: []tls.Certificate{serverCert},
			ClientCAs:    certPool,
		}), nil
}
