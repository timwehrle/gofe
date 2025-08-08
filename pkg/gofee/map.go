package gofee

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func mapToCharset(length int, charset string) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}
	l := int64(len(charset))
	if l == 0 {
		return "", fmt.Errorf("charset is empty")
	}

	out := make([]byte, length)
	for i := range out {
		n, err := rand.Int(rand.Reader, big.NewInt(l))
		if err != nil {
			return "", fmt.Errorf("rng failure: %w", err)
		}
		out[i] = charset[n.Int64()]
	}
	return string(out), nil
}

func secureShuffle(buf []byte) error {
	// Fisher-Yates shuffle using crypto/rand
	for i := len(buf) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return fmt.Errorf("rng failure during shuffle: %w", err)
		}
		j := int(n.Int64())
		buf[i], buf[j] = buf[j], buf[i]
	}
	return nil
}
