package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 20

	fmt.Println("Age =  ", age)

	show := func() {
		money := money + a + p
		fmt.Println("Money = ", money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

// another closure example: Counter
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	call()

	fmt.Println("====Counter Example====")
	c := counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	d := counter()
	fmt.Println(d()) // 1
	fmt.Println(d()) // 2
}

func init() {
	fmt.Println("====Bank=====")
}


/*

Closure হচ্ছে এমন একটি ফাংশন, যা তার বাইরের (parent) ফাংশনের ভ্যারিয়েবলগুলোকে মনে রাখতে এবং ব্যবহার করতে পারে—even যদি parent ফাংশনটি শেষও হয়ে যায়। অর্থাৎ, Closure তার আশেপাশের স্কোপের ভ্যারিয়েবলগুলোকে "বন্ধ" (close) করে রাখে।

Closure holo ek dhoroner function, ja nijer surrounding scope-er variable gulo ke access korte pare—even function ta tar nijer scope-er baire theke call korleo. Go te closure mane holo, ekta function onno ekta function-er vitore declare kora hoy, ebong oi vitorer variable gulo ke use korte pare.

func outer() func() {
    money := 100
    show := func() {
        money := money + 10
        fmt.Println(money)
    }
    return show
}

So again, Closure holo ekta function, ja tar parent function-er variable gulo ke mone rakhte pare ebong use korte pare, jodio parent function ta already execute hoye geche. Ei rokom function ke closure bole.

*/