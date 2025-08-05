package helper

import "errors"

// error enum
var ErrInvalidResponseFormat = errors.New("invalid response format from main-service")

func SafeString(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

func ExtractRole(roles interface{}) string {
	rolesSlice, ok := roles.([]interface{})
	if !ok || len(rolesSlice) == 0 {
		return ""
	}
	firstRole, ok := rolesSlice[0].(map[string]interface{})
	if !ok {
		return ""
	}
	return SafeString(firstRole["role_name"])
}
