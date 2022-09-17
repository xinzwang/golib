package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func SHA256(s string) string {
	seed := strconv.FormatInt(time.Now().UnixNano(), 10)
	var data []byte = []byte(s + seed)
	out := sha256.Sum256(data)
	return hex.EncodeToString(out[:])
}
