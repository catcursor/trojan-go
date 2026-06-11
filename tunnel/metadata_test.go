package tunnel

import (
	"bytes"
	"strings"
	"testing"
)

func TestAddressWriteToRejectsInvalidDomainLength(t *testing.T) {
	addr := NewAddressFromHostPort("tcp", strings.Repeat("a", 256), 443)
	if err := addr.WriteTo(bytes.NewBuffer(nil)); err == nil {
		t.Fatal("expected overlong domain to fail")
	}
}

func TestAddressWriteToRejectsInvalidPort(t *testing.T) {
	addr := NewAddressFromHostPort("tcp", "example.com", 70000)
	if err := addr.WriteTo(bytes.NewBuffer(nil)); err == nil {
		t.Fatal("expected invalid port to fail")
	}
}

func TestAddressWriteToRejectsInvalidIP(t *testing.T) {
	addr := &Address{
		AddressType: IPv4,
		Port:        443,
	}
	if err := addr.WriteTo(bytes.NewBuffer(nil)); err == nil {
		t.Fatal("expected invalid IPv4 to fail")
	}
}
