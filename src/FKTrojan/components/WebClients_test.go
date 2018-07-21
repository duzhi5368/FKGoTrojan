package components

import (
	"crypto/tls"
	"net/http"
	"testing"
)

func TestDefaultClient(t *testing.T) {
	client := &http.Client{}

	if client == http.DefaultClient {
		t.Errorf("should be unequal")
	}
}

func TestDefaultClientEqual(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
	}
	client := &http.Client{Transport: tr}

	if client == http.DefaultClient {
		t.Error("client and http.Default shoudl be unequal")
	}
}
