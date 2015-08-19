// Copyright (c) 2015 elarasu. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");

package version

import "fmt"

func ExampleString() {
	var v *Version = &Version{0, 0, 1}
	var v1 *Version = &Version{3, 34, 1}
	fmt.Println(v.String())
	fmt.Println(v1.String())
	// Output:
	// 0.0.1
	// 3.34.1
}

func ExampleAppString() {
	var v *Version = &Version{1, 0, 1}
	fmt.Println(v.String())
	// Output:
	// 1.0.1
}
