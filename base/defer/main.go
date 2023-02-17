package main

import "fmt"

func deferFunc(i int) (t int) {
	fmt.Println("t = ", t)
	return 2
}

func returnByDeffer() (t int) {
	defer func() {
		t = t * 10
	}()
	return 2
}

func deferCall() {
	defer func() { fmt.Println("defer func 1") }()
	defer func() { fmt.Println("defer func 2") }()

	panic("panic")
	// is never run
	//defer func() { fmt.Println("defer func last") }()
}

func main1() {
	//deferFunc(10)
	//fmt.Println(returnByDeffer())
	fmt.Println("Main step exit.")
	deferCall()

}

func main2() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("in recover.")
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("main panic")
}

func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t

}

func deferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func deferFunc4() (t int) {
	defer func(i int) {
		fmt.Println("in defer")
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func main() {
	fmt.Println(deferFunc1(1))
	fmt.Println(deferFunc2(1))
	fmt.Println(deferFunc3(1))
	deferFunc4()

}
