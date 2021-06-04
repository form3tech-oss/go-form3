package form3

import (
	"crypto/tls"
	"net/http"
)

type MTLSOption func(transport *MTLSTransportBuilder)

type MTLSTransportBuilder struct {
	certificate        tls.Certificate
	insecureSkipVerify bool
}

func WithCertificate(c tls.Certificate) MTLSOption {
	return func(t *MTLSTransportBuilder) {
		t.certificate = c
	}
}

func WithIgnoreSkipVerify(skip bool) MTLSOption {
	return func(t *MTLSTransportBuilder) {
		t.insecureSkipVerify = skip
	}
}

func NewMTLSTransport(opts ...MTLSOption) http.RoundTripper {
	tr := &MTLSTransportBuilder{}

	for _, opt := range opts {
		opt(tr)
	}

	return &http.Transport{TLSClientConfig: &tls.Config{
		Certificates:       []tls.Certificate{tr.certificate},
		InsecureSkipVerify: tr.insecureSkipVerify,
	}}
}
