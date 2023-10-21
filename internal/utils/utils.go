package utils

import (
	"errors"
	"regexp"
)

var ErrWrUrl = errors.New("sent string is not an URL")

const regex = "^(ht|f)tp(s?)://[0-9a-zA-Z]([-.\\w]*[0-9a-zA-Z])*(:(0-9)*)*(/?)([a-zA-Z0-9\\-.?,'/\\\\+&amp;%$#_]*)?$"

func ValidateUrl(str string) error {
	ok, err := regexp.MatchString(regex, str)
	if err != nil {
		return err
	}
	if !ok {
		return ErrWrUrl
	}
	return nil
}
