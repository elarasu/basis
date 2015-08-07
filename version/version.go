/*
 * Copyright(c) 2015 elarasu. All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 */
package version

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type Version struct {
	Major int64
	Minor int64
	Patch int64
}

const SEPARATOR string = "."
const BASE10 int = 10

// returns human friendly version major.minor.patch
func (v Version) String() string {
	return strings.Join(([]string{
		strconv.FormatInt(v.Major, BASE10),
		strconv.FormatInt(v.Minor, BASE10),
		strconv.FormatInt(v.Patch, BASE10)}),
		SEPARATOR)
}

// returns command line friendly version with
// appname version and go version
func (v Version) AppString(appName string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", appName, v.String(), runtime.Version())
}
