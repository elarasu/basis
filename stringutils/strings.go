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
	tmpStr := strings.TrimSpace(s)
	if s == "" || tmpStr == "" {
		return false
	}
	// check if first char is '{' || '['
	firstChar := tmpStr[0]
	if firstChar != '{' && firstChar != '[' {
		return false
	}
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func IsJsonRpc(s string) bool {
	return true
}
