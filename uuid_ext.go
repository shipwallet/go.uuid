package uuid

import "encoding/hex"

// Returns canonical string representation of UUID without dashes:
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func (u UUID) StringNoDash() string {
	buf := make([]byte, 32)

	hex.Encode(buf[0:32], u[0:16])

	return string(buf)
}
