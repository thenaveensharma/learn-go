package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("go" + "lang")
	fmt.Println(1+1, "=", "1"+"1")
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	var g string
	fmt.Println(g)

	f := "apple"
	fmt.Println(f)

	//constants

	// const s string = "constant"

	// fmt.Println(s)
	const n = 500000000
	const h = 3e20 / n
	fmt.Println(int64(h))
	fmt.Println(math.Sin(n))

	//for loop

	var i = 0

	for i <= 5 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 10; j += 1 {
		fmt.Println(j)
	}

	for k := range 13 {
		fmt.Println("range", k)
	}
	for {
		fmt.Println("loop")
		break
	}

	//if-else

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%4 == 0 || 7%2 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digit")
	}

	//switch statement
	z := 2
	switch z {
	case 1:
		fmt.Println("num is 1")

	case 2:
		fmt.Println("num is 2")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	//arrays

	var y [5]int
	y[4] = 100
	fmt.Println("set:", y)
	fmt.Println("get:", y[4])
	var x [5]string
	fmt.Println("emp:", x)
	var w [5]bool
	fmt.Println("emp:", w)

	v := [5]int{1, 2, 3, 4, 5}
	fmt.Println(v)
	t := [...]int{1, 2, 3, 4, 5}
	fmt.Println(t)
	s := [...]int{1, 5: 50}
	fmt.Println(s)

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	//Slices

	var m []string
	fmt.Println("unint:", m, m == nil, len(m) == 0)

	m = make([]string, 10, 11)
	fmt.Println("emp:", m, "len:", len(m), "cap:", cap(m))

	m = append(m, "a", "b", "c", "d")
	fmt.Println(m)
	fmt.Println("emp:", m, "len:", len(m), "cap:", cap(m))

	o := make([]string, len(m))
	copy(o, m)
	fmt.Println("emp:", o, "len:", len(o), "cap:", cap(o))

	//Maps

	var _map = make(map[string]int)
	_map["1"] = 1
	_map["2"] = 2
	_, prss := _map["2"]
	fmt.Println(prss)
	_, prs := _map["3"]
	fmt.Println(prs)

	fmt.Println(_map)

	//Functions
	fmt.Println(plus(2, 3))
	fmt.Println(plusPlus(2, 3, 5))

	// Multiple return value

	fmt.Println(returnMulValues())

	// Variadic Functions

	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func returnMulValues() (int, string) {

	return 1, "naveen"
}

func sum(nums ...int) {

	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
