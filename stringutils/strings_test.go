// Copyright (c) 2015 elarasu. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");

package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsJson(t *testing.T) {
	a := [...]string{"Not a json obj",
		"",
		"      \t\r\n   \t",
		` { "here"`,
		" [ some where ]]",
		" {{}",
	}
	assert := assert.New(t)
	for i := 0; i < len(a); i++ {
		assert.False(IsJson(a[i]), a[i])
	}
	b := [...]string{"  {}",
		`{"a":["x","y"]}`,
		`   ["one"]`,
	}
	for i := 0; i < len(b); i++ {
		assert.True(IsJson(b[i]), b[i])
	}
}

func BenchmarkIsJson(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
