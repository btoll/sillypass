package sillypass

import (
	"crypto/rand"
	"errors"
	"strings"
)

const BUCKET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_.,<>?:!@#$%&*()-+[]^=/'\"\\`"
const BUCKETLEN = len(BUCKET)

func Generate(n int) (string, error) {
	if n < 3 {
		return "", errors.New("Cowardly refusing to generate a password less than three characters.")
	}

	b := make([]byte, n)

	// The slice will now contain `n` random bytes.
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	s := make([]string, n)

	for i, j := range b {
		// Are these conversions considered bad practice?
		s[i] = string(BUCKET[int(j)%BUCKETLEN])
	}

	return strings.Join(s, ""), nil
}
