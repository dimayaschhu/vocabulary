package utils

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

var placeholderRegexp = regexp.MustCompile(`<.*>`)

type PlaceholderValidator struct {
}

func NewPlaceholderValidator() *PlaceholderValidator {
	return &PlaceholderValidator{}
}

func (v *PlaceholderValidator) IsPlaceholderValue(expectedValue string) bool {
	return placeholderRegexp.MatchString(expectedValue)
}

func (v *PlaceholderValidator) Validate(placeholder string, actualValue string) error {
	switch placeholder {
	case "<string>":
		if actualValue == "" {
			return errors.New("Empty string")
		}
	case "<_id>":
		if actualValue == "" {
			return errors.New("Empty id")
		}
	case "<date_time>":
		_, err := time.Parse(time.RFC3339Nano, actualValue)
		if err != nil {
			return errors.Wrap(err, "invalid time format")
		}
	case "<null>":
		if actualValue == "" {
			return nil
		}

		if actualValue == "0" {
			return nil
		}

		return errors.Errorf("expected null value, got: %q", actualValue)
	default:
		return errors.Errorf("Placeholder '%s' is not defined", placeholder)
	}

	return nil
}
