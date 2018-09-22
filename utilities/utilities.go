package utilities

import b64 "encoding/base64"

func EncodeBase64String(someString string) string {
	stringEncoding := b64.StdEncoding.EncodeToString([]byte(someString))
	return stringEncoding
}
