package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

func Hash[T any](data T) string {

	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(data)

	h := sha256.New()
	h.Write(b.Bytes())

	return fmt.Sprintf("%x", h.Sum(nil))
}

func Clamp(num, min, max int) int {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}
