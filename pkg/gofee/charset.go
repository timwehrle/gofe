package gofee

import "strings"

// Charset constants for lowercase letters, uppercase letters, digits, and symbols.
const (
	Lowers  = "abcdefghijklmnopqrstuvwxyz"
	Uppers  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits  = "0123456789"
	Symbols = "!@#$%^&*()-_=+[]{}|;:,.<>?/~"
)

type PasswordConfig struct {
	IncludeLowers  bool
	IncludeUppers  bool
	IncludeDigits  bool
	IncludeSymbols bool
	Type           string

	RequireClasses bool
	MinPINLength   int
	MinLength      int
	MaxLength      int
}

func BuildCharset(cfg PasswordConfig) string {
	switch cfg.Type {
	case "pin":
		return Digits
	case "memorable":
		return Lowers + Uppers
	}

	if cfg.IncludeLowers && cfg.IncludeUppers && cfg.IncludeDigits && cfg.IncludeSymbols {
		return Lowers + Uppers + Digits + Symbols
	}

	var b strings.Builder
	if cfg.IncludeLowers {
		b.WriteString(Lowers)
	}

	if cfg.IncludeUppers {
		b.WriteString(Uppers)
	}

	if cfg.IncludeDigits {
		b.WriteString(Digits)
	}

	if cfg.IncludeSymbols {
		b.WriteString(Symbols)
	}

	return b.String()
}

func selectedSets(cfg PasswordConfig) []string {
	switch cfg.Type {
	case "pin":
		return []string{Digits}
	case "memorable":
		return []string{Lowers, Uppers}
	default:
		sets := make([]string, 0, 4)
		if cfg.IncludeLowers {
			sets = append(sets, Lowers)
		}
		if cfg.IncludeUppers {
			sets = append(sets, Uppers)
		}
		if cfg.IncludeDigits {
			sets = append(sets, Digits)
		}
		if cfg.IncludeSymbols {
			sets = append(sets, Symbols)
		}
		return sets
	}
}
