package ogenadapter_test

import (
	ogenlib "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
)

func getStringFromOptString(optString ogenlib.OptString) *string {
	if optString.Set {
		return &optString.Value
	}
	return nil
}

func toOptString(value *string) ogenlib.OptString {
	if value != nil {
		return ogenlib.OptString{
			Value: *value,
			Set:   true,
		}
	}
	var v string
	return ogenlib.OptString{
		Value: v,
		Set:   false,
	}
}
