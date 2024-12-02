package main

import (
	"fmt"
)

type Person struct {
    Name string
    Age  int
}

func runBasics() {
	var name string = "Alice"
    var age int = 30

    // Shorthand declaration
    city := "New York"
	city = "Sydney";
	fmt.Println(name, age, city);

	const Pi float64 = 3.14
    const Greeting = "Hello, Go!"
	fmt.Println(Pi, Greeting);

	var a int = 42
    var b float64 = 3.14
    var c string = "Hello"
    var d bool = true
	fmt.Println(a, b, c, d)

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	slice := []int{1, 2, 3, 4, 5}
    slice = append(slice, 6)
    fmt.Println(slice)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	num := 10
    if num > 5 {
        fmt.Println("Greater than 5")
    } else {
        fmt.Println("5 or less")
    }

    // Switch statement
    switch num {
    case 1:
        fmt.Println("One")
    case 10:
        fmt.Println("Ten")
    default:
        fmt.Println("Unknown")
    }

	scores := map[string]int{"Alice": 90, "Bob": 85}
    scores["Charlie"] = 88
    fmt.Println(scores)

	p := Person{Name: "Jonatan", Age: 29}
	fmt.Println(p)

	defer fmt.Println("Executed last") // Runs when the function exits
    fmt.Println("Executed first")
    
}