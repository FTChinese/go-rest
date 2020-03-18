package rand

import "testing"

func TestHex(t *testing.T) {
	hex, err := Hex(32)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Hexdeciaml: %s", hex)
}

func TestBase64(t *testing.T) {
	s, err := Base64(12)
	if err != nil {
		t.Error(err)
	}

	t.Logf("16 chars of base64: %s", s)
}
