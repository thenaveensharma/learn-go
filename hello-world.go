// package main

// import (
// 	"fmt"
// 	"math"
// 	"unicode/utf8"
// )

// func main() {

// 	fmt.Println("go" + "lang")
// 	fmt.Println(1+1, "=", "1"+"1")
// 	fmt.Println("7.0/3.0 =", 7.0/3.0)
// 	fmt.Println(true && false)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)
// 	fmt.Println(!false)

// 	var a = "initial"
// 	fmt.Println(a)

// 	var b, c int = 1, 2
// 	fmt.Println(b, c)
// 	var d = true
// 	fmt.Println(d)
// 	var e int
// 	fmt.Println(e)
// 	var g string
// 	fmt.Println(g)

// 	f := "apple"
// 	fmt.Println(f)

// 	//constants

// 	// const s string = "constant"

// 	// fmt.Println(s)
// 	const n = 500000000
// 	const h = 3e20 / n
// 	fmt.Println(int64(h))
// 	fmt.Println(math.Sin(n))

// 	//for loop

// 	var i = 0

// 	for i <= 5 {
// 		fmt.Println(i)
// 		i += 1
// 	}

// 	for j := 0; j < 10; j += 1 {
// 		fmt.Println(j)
// 	}

// 	for k := range 13 {
// 		fmt.Println("range", k)
// 	}
// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	//if-else

// 	if 7%2 == 0 {
// 		fmt.Println("7 is even")
// 	} else {
// 		fmt.Println("7 is odd")
// 	}

// 	if 8%4 == 0 {
// 		fmt.Println("8 is divisible by 4")
// 	}

// 	if 8%4 == 0 || 7%2 == 0 {
// 		fmt.Println("8 is divisible by 4")
// 	}

// 	if num := 9; num < 0 {
// 		fmt.Println(num, " is negative")
// 	} else if num < 10 {
// 		fmt.Println(num, " has 1 digit")
// 	} else {
// 		fmt.Println(num, " has multiple digit")
// 	}

// 	//switch statement
// 	z := 2
// 	switch z {
// 	case 1:
// 		fmt.Println("num is 1")

// 	case 2:
// 		fmt.Println("num is 2")
// 	}

// 	whatAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case bool:
// 			fmt.Println("I'm a bool")
// 		case int:
// 			fmt.Println("I'm an int")
// 		default:
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}
// 	whatAmI(true)
// 	whatAmI(1)
// 	whatAmI("hey")

// 	//arrays

// 	var y [5]int
// 	y[4] = 100
// 	fmt.Println("set:", y)
// 	fmt.Println("get:", y[4])
// 	var x [5]string
// 	fmt.Println("emp:", x)
// 	var w [5]bool
// 	fmt.Println("emp:", w)

// 	v := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(v)
// 	t := [...]int{1, 2, 3, 4, 5}
// 	fmt.Println(t)
// 	s := [...]int{1, 5: 50}
// 	fmt.Println(s)

// 	var twoD [2][3]int
// 	for i := 0; i < len(twoD); i++ {
// 		for j := 0; j < len(twoD[0]); j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}

// 	fmt.Println(twoD)

// 	//Slices

// 	var m []string
// 	fmt.Println("unint:", m, m == nil, len(m) == 0)

// 	m = make([]string, 10, 11)
// 	fmt.Println("emp:", m, "len:", len(m), "cap:", cap(m))

// 	m = append(m, "a", "b", "c", "d")
// 	fmt.Println(m)
// 	fmt.Println("emp:", m, "len:", len(m), "cap:", cap(m))

// 	o := make([]string, len(m))
// 	copy(o, m)
// 	fmt.Println("emp:", o, "len:", len(o), "cap:", cap(o))

// 	//Maps

// 	var _map = make(map[string]int)
// 	_map["1"] = 1
// 	_map["2"] = 2
// 	_, prss := _map["2"]
// 	fmt.Println(prss)
// 	_, prs := _map["3"]
// 	fmt.Println(prs)

// 	fmt.Println(_map)

// 	//Functions
// 	fmt.Println(plus(2, 3))
// 	fmt.Println(plusPlus(2, 3, 5))

// 	// Multiple return value

// 	fmt.Println(returnMulValues())

// 	// Variadic Functions

// 	sum(1, 2)
// 	sum(1, 2, 3)
// 	nums := []int{1, 2, 3, 4}
// 	sum(nums...)

// 	// Closures
// 	nextInt := intSec()
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	nextInts := intSec()
// 	fmt.Println(nextInts())

// 	// Recursion
// 	fmt.Println(findFactorial(10))

// 	// Find fibonacci number
// 	var fib func(n int) int
// 	fib = func(n int) int {
// 		if n < 2 {
// 			return n

// 		}
// 		return fib(n-1) + fib(n-2)
// 	}
// 	fmt.Println(fib(10))

// 	//Range over Build-in Types

// 	numbers := []int{1, 2, 3, 4, 5}
// 	sum := 0
// 	for index, num := range numbers {
// 		sum += num
// 		fmt.Print(index, sum, " ")
// 	}
// 	kvs := map[string]string{"a": "apple", "b": "ball"}
// 	for k, v := range kvs {
// 		fmt.Printf("%s -> %s\n", k, v)
// 	}

// 	for i, c := range "naveen" {
// 		fmt.Println(i, c)
// 	}

// 	//Pointers

// 	val := 1
// 	fmt.Println("inital", val)
// 	zeroVal(val)
// 	fmt.Println("zeroVal", val)
// 	ptr := 1
// 	fmt.Println("inital", ptr)
// 	zeroPointer(&ptr)
// 	fmt.Println("zeroVal", ptr)
// 	fmt.Println("pointer", &ptr)

// 	const anString = "Hello Naveen"

// 	for i := range len(anString) {
// 		fmt.Printf("%x ", anString[i])
// 	}
// 	fmt.Println()

// 	fmt.Println("Rune count: ", utf8.RuneCountInString(anString))

// 	for idx, runeValue := range anString {
// 		fmt.Printf("%#U starts at %d\n", runeValue, idx)
// 	}
// 	fmt.Println("\nUsing DecodeRuneInString")

// 	for i, w := 0, 0; i < len(anString); i += w {
// 		runeValue, width := utf8.DecodeRuneInString(anString[i:])
// 		fmt.Printf("%#U starts at %d\n", runeValue, i)
// 		w = width
// 		examineRune(runeValue)
// 	}

// 	//Structs
// 	fmt.Println(person{name: "bol", age: 20})
// 	fmt.Println(person{name: "Alice", age: 21})
// 	fmt.Println(&person{name: "Naveen"})
// 	fmt.Println(newPerson("Nik"))

// 	randomP := person{name: "Random", age: 40}
// 	fmt.Println(randomP)
// 	fmt.Println(randomP.name, randomP.age)
// 	randomP.age = 18
// 	fmt.Println(randomP.name, randomP.age)

// 	dog := struct {
// 		name   string
// 		isGood bool
// 	}{
// 		"Alex",
// 		true,
// 	}
// 	fmt.Println(dog)
// 	//Methods

// 	r := rect{width: 10, height: 20}
// 	fmt.Println("area: ", r.area())
// 	fmt.Println("perim: ", r.perim())
// 	rp := &r
// 	fmt.Println("area: ", rp.area())
// 	fmt.Println("perim: ", rp.perim())

// 	//Interface
// 	rectangle := rectangle{width: 2.1, height: 3.1111}
// 	circle := circle{radius: 3.3}

// 	detectCircle(circle)
// 	detectCircle(rectangle)
// 	measure(circle)
// 	measure(rectangle)

// 	//Enums
// 	ns := transition(StateIdle)
// 	fmt.Println(ns)
// 	ns2 := transition(ns)
// 	fmt.Println(ns2)
// 	// Struct Embedding

// 	co := container{
// 		base: base{
// 			num: 1,
// 		},
// 		str: "Naveen",
// 	}
// 	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
// 	fmt.Println("also num:", co.base.num)
// 	fmt.Println("describe: ", co.describe())

// 	type describer interface {
// 		describe() string
// 	}
// 	var de describer = co
// 	fmt.Println("describer:", de.describe())
// }

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// type ServerState int

// const (
// 	StateIdle ServerState = iota
// 	StateConnected
// 	StateError
// 	StateRetrying
// )

// var stateName = map[ServerState]string{
// 	StateIdle:      "idle",
// 	StateConnected: "connected",
// 	StateError:     "error",
// 	StateRetrying:  "retrying",
// }

// func (ss ServerState) String() string {
// 	return stateName[ss]
// }

// func transition(s ServerState) ServerState {
// 	switch s {
// 	case StateIdle:
// 		return StateConnected
// 	case StateConnected, StateRetrying:
// 		return StateIdle
// 	case StateError:
// 		return StateError
// 	default:
// 		panic(fmt.Errorf("unknown state: %s", s))
// 	}
// }

// type geometry interface {
// 	area() float64
// 	perim() float64
// }

// type rectangle struct {
// 	width, height float64
// }
// type circle struct {
// 	radius float64
// }

// func (r rectangle) area() float64 {
// 	return r.height * r.width

// }
// func (r rectangle) perim() float64 {
// 	return 2 * (r.height + r.width)
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// func detectCircle(g geometry) {
// 	if c, ok := g.(circle); ok {
// 		fmt.Println("circle with radius r: ", c.radius)
// 	}
// }

// type rect struct {
// 	width, height int
// }

// func (r *rect) area() int {
// 	return r.height * r.width
// }
// func (r rect) perim() int {
// 	return 2 * (r.width + r.height)
// }

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string) *person {
// 	p := person{name: name}
// 	p.age = 42
// 	return &p
// }
// func examineRune(r rune) {
// 	if r == 't' {
// 		fmt.Println("found tee")
// 	} else if r == 'e' {
// 		fmt.Println(("found e"))
// 	}
// }

// func zeroVal(a int) {
// 	a = 0
// }
// func zeroPointer(a *int) {
// 	*a = 0
// }

// // find Factorial of a number
// func findFactorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}

// 	return n * findFactorial(n-1)
// }

// func intSec() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}

// }

// func plus(a int, b int) int {
// 	return a + b
// }

// func plusPlus(a, b, c int) int {
// 	return a + b + c
// }

// func returnMulValues() (int, string) {

// 	return 1, "naveen"
// }

// func sum(nums ...int) {

// 	fmt.Print(nums, " ")
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }
