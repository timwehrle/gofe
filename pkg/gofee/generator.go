package gofee

import (
	"fmt"
)

func Generate(length int, cfg PasswordConfig) (string, string, error) {
	if length <= 0 {
		return "", "", fmt.Errorf("length must be greater than 0")
	}

	minLen := cfg.MinLength
	if cfg.Type == "pin" && cfg.MinPINLength > 0 {
		minLen = cfg.MinPINLength
	}
	if minLen > 0 && length < minLen {
		return "", "", fmt.Errorf("length must be >= %d", minLen)
	}
	if cfg.MaxLength > 0 && length > cfg.MaxLength {
		return "", "", fmt.Errorf("length must be <= %d", cfg.MaxLength)
	}

	charset := BuildCharset(cfg)
	if len(charset) == 0 {
		return "", "", fmt.Errorf("no characters available (check excludes or type)")
	}

	sets := selectedSets(cfg)
	if cfg.RequireClasses && len(sets) > 1 {
		if length < len(sets) {
			return "", "", fmt.Errorf("length must be >= %d to include all selected character classes", len(sets))
		}

		req := make([]byte, 0, len(sets))
		for _, s := range sets {
			ch, err := mapToCharset(1, s)
			if err != nil {
				return "", "", err
			}
			req = append(req, ch[0])
		}

		restStr, err := mapToCharset(length-len(req), charset)
		if err != nil {
			return "", "", err
		}
		buf := append(req, []byte(restStr)...)

		if err := secureShuffle(buf); err != nil {
			return "", "", err
		}
		return string(buf), charset, nil
	}

	pw, err := mapToCharset(length, charset)
	if err != nil {
		return "", "", err
	}

	return pw, charset, nil
}
