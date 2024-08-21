package util

import "errors"

func IsCommonError(err error, _errors []error) bool {
	for _, commonErr := range _errors {
		if errors.Is(err, commonErr) {
			return true
		}
	}
	return false
}
