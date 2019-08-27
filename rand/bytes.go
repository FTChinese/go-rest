package rand

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

// Bytes generates cypto-strong random bytes.
func Bytes(len int) ([]byte, error) {
	b := make([]byte, len)

	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// Hex generates a random hexadecimal number of 2*len chars
func Hex(len int) (string, error) {

	b, err := Bytes(len)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b[:]), nil
}

// Base64 returns a base64 url encoded, unpadded string of random bytes.
// len is the length of bytes.
// It should be the multiple of 3 to get a string without padding:
// len * 8 / 6.
// 3 bytes -- 4 chars
// 6 bytes -- 8 chars
// 9 bytes -- 12 chars
// Wechat OAuth code has 32 characters, which is 24 bytes long;
// Wechat Access Token has 110 chars, which is 82.5 bytes?
// Wechat Refresh Token has 110 chars.
// OpenID has 28 chars, which is 21 bytes
// UnionID has 28 chars.
func Base64(len int) (string, error) {
	b, err := Bytes(len)

	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
