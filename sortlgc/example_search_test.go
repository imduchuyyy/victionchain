// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sortlgc_test

import (
	"fmt"

	"github.com/tomochain/tomochain/sortlgc"
)

// This example demonstrates searching a list sorted in ascending order.
func ExampleSearch() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sortlgc.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	// Output:
	// found 6 at index 2 in [1 3 6 10 15 21 28 36 45 55]
}

// This example demonstrates searching a list sorted in descending order.
// The approach is the same as searching a list in ascending order,
// but with the condition inverted.
func ExampleSearch_descendingOrder() {
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x := 6

	i := sortlgc.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	// Output:
	// found 6 at index 7 in [55 45 36 28 21 15 10 6 3 1]
}

// This example demonstrates searching for float64 in a list sorted in ascending order.
func ExampleSearchFloat64s() {
	a := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}

	x := 2.0
	i := sortlgc.SearchFloat64s(a, x)
	fmt.Printf("found %g at index %d in %v\n", x, i, a)

	x = 0.5
	i = sortlgc.SearchFloat64s(a, x)
	fmt.Printf("%g not found, can be inserted at index %d in %v\n", x, i, a)
	// Output:
	// found 2 at index 1 in [1 2 3.3 4.6 6.1 7.2 8]
	// 0.5 not found, can be inserted at index 0 in [1 2 3.3 4.6 6.1 7.2 8]
}

// This example demonstrates searching for int in a list sorted in ascending order.
func ExampleSearchInts() {
	a := []int{1, 2, 3, 4, 6, 7, 8}

	x := 2
	i := sortlgc.SearchInts(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i, a)

	x = 5
	i = sortlgc.SearchInts(a, x)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)
	// Output:
	// found 2 at index 1 in [1 2 3 4 6 7 8]
	// 5 not found, can be inserted at index 4 in [1 2 3 4 6 7 8]
}
