package types

import (
	"regexp"
)

const (
	TypeNumber     byte = 1
	TypeString     byte = 2
	TypeIdentifier byte = 3
)

const (
	RegTokenNumber     = `[0-9]+(\.[0-9]+)?`
	RegTokenString     = `"([^"]|\\"|\n|\\)*"`
	RegTokenIdentifier = `([a-z_A-Z][a-z_A-Z\d]*|==|<|<=|>|>=|&&|\\|[!"#$%&'()*+,-./:;<=>?@[\]^_{|}~])`
	RegTokenComment    = `//.*`
	RegToken           = "\\s*((" + RegTokenComment + ")|(" + RegTokenIdentifier +
		")|(" + RegTokenIdentifier + ")|(" + RegTokenNumber + "))"
)

var (
	RegNumber,
	RegString,
	RegIdentifier,
	RegComment *regexp.Regexp
)

func init() {
	RegNumber = regexp.MustCompile(RegTokenNumber)
	RegString = regexp.MustCompile(RegTokenString)
	RegIdentifier = regexp.MustCompile(RegTokenIdentifier)
	RegComment = regexp.MustCompile(RegTokenComment)
}
