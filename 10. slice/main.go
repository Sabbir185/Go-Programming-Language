package main

import "fmt"

// variadic function
func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

func main() {
	arr := [6]string{"This", "is", "a", "go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s) // [is a go]

	// slice from a slice or slice to slice
	s2 := s[1:2]
	fmt.Println(s2)     // [a]
	fmt.Println(len(s)) // 1
	fmt.Println(cap(s)) // 4 {"a", "go", "interview", "question"}

	sl := []int{1, 2, 3, 4, 5} // slice literal --> array er size define na korle slice hoye jay
	fmt.Println("slice: ", sl, "len: ", len(sl), "cap: ", cap(sl))
	// length and capacity same for slice literal

	msl := make([]int, 3)
	fmt.Println(msl)
	fmt.Println(len(msl))
	fmt.Println(cap(msl))
	msl[0] = 10
	msl[1] = 20
	msl[2] = 30
	fmt.Println(msl)

	msl2 := make([]int, 3, 5) // length 3, capacity 5
	msl2[0] = 10
	msl2[3] = 20 // invalid, because len is 3, out of range, array ar bayre
	fmt.Println(msl2)

	var ss []int // emply slice or nil slice
	fmt.Println(ss)
	// pointer = nil, length = 0, capacity = 0

	ss = append(ss, 3) // add value into ss nil slice
	fmt.Println(ss)    // it will create a new array with length 1 and capacity 1 and return array

	ss = append(ss, 4, 5, 6)

	// -------------------------
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y := x
	x = append(x, 4)
	y = append(y, 5)
	x[0] = 10
	fmt.Println(x)
	fmt.Println(y)

	// -------------------------
	aa := []int{1, 2, 3, 4, 5}
	aa = append(aa, 6)
	aa = append(aa, 7)

	bb := aa[4:]

	yy := changeSlice(bb)

	fmt.Println(aa) // [1 2 3 4 10 6 7]
	fmt.Println(yy) // [10 6 7 11]

	fmt.Println(aa[0:8]) // [1 2 3 4 10 6 7 11] bxz of capacity
}

/*

1. Slice from an existing array
2. Slice from a slice
3. Slice literal
4. using make function
5. make function with length and capacity
6. empty slice
7. slice underlying array
   * len & cap same hole, then new array create & capacity increase kore
   * 1st 1024 er jonno double size memory allocate kore (array capacity)
   * 1024 er beshi hole 25% increment kore

*/

// slice -> part of array
// Pizza is the best example of slice

// slice holds --> pointer, length, capacity
// pointer --> value of starting index of array --> address of array
// length --> number of elements in slice
// capacity --> starting index of array to last index of array
// so, pointer = 1, length = 3, capacity = 5
