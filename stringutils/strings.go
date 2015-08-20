// Copyright (c) 2015 elarasu. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");

// stringutils - utilities on string
package stringutils

import (
	"encoding/json"
	"strings"
)

// do some simple checks before going
// to invoke json Unmarshal
func IsJson(s string) bool {
	b, _ := _isJson(s)
	return b
}

// for now, consider only object type
// @todo: array types to be handled
func IsJsonRpc(s string) bool {
	b, js := _isJson(s)
	if b {
		switch js.(type) {
		case map[string]interface{}:
			obj := js.(map[string]interface{})
			if _, ok := obj["jsonrpc"]; ok {
				return true
			}
		}
	}
	return false
}

// --- internal methods --

// do some simple checks before going
// to invoke json Unmarshal
func _isJson(s string) (bool, interface{}) {
	if s == "" {
		return false, nil
	}
	tmpStr := strings.TrimSpace(s)
	if tmpStr == "" {
		return false, nil
	}
	// check if first char is '{' || '['
	firstChar := tmpStr[0]
	if firstChar != '{' && firstChar != '[' {
		return false, nil
	}

	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil, js
}
