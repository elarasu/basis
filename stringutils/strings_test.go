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
		r, _ := IsJson(a[i])
		assert.False(r, a[i])
	}
	b := [...]string{"  {}",
		`{"a":["x","y"]}`,
		`   ["one"]`,
	}
	for i := 0; i < len(b); i++ {
		r, _ := IsJson(b[i])
		assert.True(r, b[i])
	}
}

func BenchmarkIsJson(b *testing.B) {
	a := [...]string{"  {}",
		`{"a":["x","y"]}`,
		`   ["one"]`,
		"",
		"      \t\r\n   \t",
		` { "here"`,
	}
	for i := 0; i < b.N; i++ {
		IsJson(a[i%len(a)])
	}
}

func TestIsJsonRpc(t *testing.T) {
	a := [...]string{"{}",
		` { "here":12 }`,
		` [{ "jsonrpc":1 }]`,
	}
	assert := assert.New(t)
	for i := 0; i < len(a); i++ {
		r, _ := IsJsonRpc(a[i])
		assert.False(r, a[i])
	}
	b := [...]string{
		`{"jsonrpc":["x","y"]}`,
		`{"one":1, "two":[1,2,3],"jsonrpc":"hello"}`,
	}
	for i := 0; i < len(b); i++ {
		r, _ := IsJsonRpc(b[i])
		assert.True(r, b[i])
	}
}

func BenchmarkIsJsonRpc(b *testing.B) {
	a := [...]string{"  {}",
		` { "here":12 }`,
		` [{ "jsonrpc":1 }]`,
		`{"jsonrpc":["x","y"]}`,
		`{"one":1, "two":[1,2,3],"jsonrpc":"hello"}`,
	}
	for i := 0; i < b.N; i++ {
		IsJsonRpc(a[i%len(a)])
	}
}
