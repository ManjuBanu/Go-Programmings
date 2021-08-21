package main

import "fmt"
import "runtime"

//Ninja Level 1
//Hands-on exercise #1
func main() {
	Ninja1_1()
	Ninja1_2()
	Ninja1_3()
	Ninja1_4()
	Ninja1_5()
	Ninja1_6()
	numeric_types()
	numeric_types1()
	runtime_package()
	string_type()
	Ninja2_1()
	Ninja2_2()
	Ninja2_3()
	Ninja2_4()
	Ninja2_5()
	Ninja2_6()
	Ninja3_1()
	Ninja3_2()
	Ninja3_3()
	Ninja3_4()
	Ninja3_5()
	Ninja3_6()
	Ninja3_7()
	Ninja3_8()
	Ninja3_9()
	Ninja3_10()
}

// Exercises - [Ninja Level 1]
//1.	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 1.	42
// 2.	“James Bond”
// 3.	true
// 2.	Now print the values stored in those variables using
// 1.	a single print statement
// 2.	multiple print statements
func Ninja1_1() {
	fmt.Println("** Ninja1_1 **")
	x := 42
	fmt.Println(x)
	y := "James Bond"
	fmt.Println(y)
	z := true
	fmt.Println(z)

	fmt.Println(x, y, "\t", z)
}

//1.	Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
// 1.	identifier “x” type int
// 2.	identifier “y” type string
// 3.	identifier “z” type bool
// 2.	in func main
// 1.	print out the values for each identifier
// 2.	The compiler assigned values to the variables. What are these values called?
var x int
var y string
var z bool

func Ninja1_2() {
	fmt.Println("** Ninja1_2 **")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

//Using the code from the previous exercise,
// 1.	At the package level scope, assign the following values to the three variables
// 1.	for x assign 42
// 2.	for y assign “James Bond”
// 3.	for z assign true
// 2.	in func main
// 1.	use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
// 2.	print out the value stored by variable “s”
var a int = 42
var b string = "James Bond"
var c bool = true

func Ninja1_3() {
	fmt.Println("** Ninja1_3 **")
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)
}

// •	FYI - nice documentation and new terminology “underlying type”
// o	https://golang.org/ref/spec#Types
// For this exercise
// 1.	Create your own type. Have the underlying type be an int.
// 2.	create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// 3.	in func main
// 1.	print out the value of the variable “x”
// 2.	print out the type of the variable “x”
// 3.	assign 42 to the VARIABLE “x” using the “=” OPERATOR
// 4.	print out the value of the variable “x”
type hotdog int

var k hotdog

func Ninja1_4() {
	fmt.Println("** Ninja1_4 **")
	fmt.Println(k)
	fmt.Printf("%T\n", k)
	k = 42
	fmt.Println(k)
}

//Building on the code from the previous example
// 1.	at the package level scope, using the “var” keyword,
// create a VARIABLE with the IDENTIFIER “y”. The variable should be of the
// UNDERLYING TYPE of your custom TYPE “x”
// 2.	in func main
// a.	this should already be done
// i.	print out the value of the variable “x”
// ii.	print out the type of the variable “x”
// iii.	assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// iv.	print out the value of the variable “x”
// b.	now do this
//  .	now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
// 1.	then use the “=” operator to ASSIGN that value to “y”
// 2.	print out the value stored in “y”
// 3.	print out the type of “y”\
type hotdog1 int

var m hotdog1
var l int

func Ninja1_5() {
	fmt.Println("** Ninja1_5 **")
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	m = 45
	fmt.Println(m)
	l = int(m)
	fmt.Println(l)
	fmt.Printf("%T", l)
}

//Hands-on exercise #6
var x1 bool

func Ninja1_6() {
	fmt.Println("** Ninja1_6 **")
	fmt.Println(x1)
	x1 = true
	fmt.Println(x1)

	a1 := 7
	b1 := 42
	fmt.Println(a1 == b1)
}

// Numeric types
var x2 int
var y2 float64

func numeric_types() {
	fmt.Println("** Numeric_types **")
	x := 42
	y := 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	y2 = 42.34534
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Printf("%T\n", x2)
	fmt.Printf("%T\n", y2)
}

// Numeric types
var x3 int8 = -128

func numeric_types1() {
	fmt.Println("** Numeric_types1 **")
	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
}

//runtime package
func runtime_package() {
	fmt.Println("** Runtime_package **")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

//String Type
func string_type() {
	fmt.Println("** String_type **")
	s1 := "Hello, 世界"
	fmt.Printf("%s\n", s1)
	fmt.Printf("%q\n", s1)
	fmt.Printf("%x\n", s1)
	fmt.Printf("---%x\n", "世")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%x ", s1[i])
	}

	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}

	s3 := "Hello, playground"
	fmt.Println(s3)
	s3 = "Hello, Hawaii"
	fmt.Println(s3)
	fmt.Printf("%T\n", s3)

	bs := []byte(s3)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%#U ", s3[i])
	}

	fmt.Println("")

	for i, v := range s3 {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}

// Exercises -[ Ninja Level 2 ]
//Write a program that prints a number in decimal, binary, and hex
func Ninja2_1() {
	fmt.Println("** Ninja2_1 **")
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}

//Using the following operators, write expressions and assign their values to variables:
// g.	==
// h.	<=
// i.	>=
// j.	!=
// k.	<
// l.	>
// Now print each of the variables.
func Ninja2_2() {
	fmt.Println("** Ninja2_2 **")
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
}

//Create TYPED and UNTYPED constants. Print the values of the constants.
const (
	a1     = 42
	b1 int = 43
)

func Ninja2_3() {
	fmt.Println("** Ninja2_3 **")
	fmt.Println(a1, b1)
}

// Write a program that
// •	assigns an int to a variable
// •	prints that int in decimal, binary, and hex
// •	shifts the bits of that int over 1 position to the left, and assigns that to a variable
// •	prints that variable in decimal, binary, and hex
func Ninja2_4() {
	fmt.Println("** Ninja2_4 **")
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}

//Create a variable of type string using a raw string literal. Print it.
func Ninja2_5() {
	fmt.Println("** Ninja2_5 **")
	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}

//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
const (
	a2 = 2017 + iota
	b2 = 2017 + iota
	c2 = 2017 + iota
	d2 = 2017 + iota
)

func Ninja2_6() {
	fmt.Println("** Ninja2_6 **")
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2)
}

// Exercises - [Ninja Level 3]
// Print every number from 1 to 10,000
func Ninja3_1() {
	fmt.Println("** Ninja3_1 **")
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

// Print every rune code point of the uppercase alphabet three times. Your output should look like this:
// 65
// 	U+0041 'A'
// 	U+0041 'A'
// U+0041 'A'
// 66
// 	U+0042 'B'
// 	U+0042 'B'
// 	U+0042 'B'
//  … through the rest of the alphabet characters
func Ninja3_2() {
	fmt.Println("** Ninja3_2 **")
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

// Create a for loop using this syntax
// •	for condition { }
// Have it print out the years you have been alive.
func Ninja3_3() {
	fmt.Println("** Ninja3_3 **")
	bd := 1985
	for bd <= 2017 {
		fmt.Println(bd)
		bd++
	}
}

// Create a for loop using this syntax
// •	for { }
// Have it print out the years you have been alive.
func Ninja3_4() {
	fmt.Println("** Ninja3_4 **")
	bd := 1985
	for {
		if bd > 2017 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
func Ninja3_5() {
	fmt.Println("** Ninja3_5 **")
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

// Create a program that shows the “if statement” in action.
func Ninja3_6() {
	fmt.Println("** Ninja3_6 **")
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}
}

// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
func Ninja3_7() {
	fmt.Println("** Ninja3_7 **")
	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}

// Create a program that uses a switch statement with no switch expression specified.
func Ninja3_8() {
	fmt.Println("** Ninja3_8 **")
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

// Create a program that uses a switch statement with the switch expression
//  specified as a variable of TYPE string with the IDENTIFIER “favSport”.
func Ninja3_9() {
	fmt.Println("** Ninja3_9 **")
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}

// Write down what these print:
// •	fmt.Println(true && true)
// •	fmt.Println(true && false)
// •	fmt.Println(true || true)
// •	fmt.Println(true || false)
// •	fmt.Println(!true)
func Ninja3_10() {
	fmt.Println("** Ninja3_10 **")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// func main() {
// 	//declar and assign
// 	x := 10
// 	fmt.Println(x)
// 	x = 1210
// 	fmt.Println(x)
// 	y := 10
// 	fmt.Println(y)
// 	z := "manju muthu"
// 	fmt.Println(z)
// }

//[example for data types , declaring, assigning ]

// var z int

// func main() {
// 	z = 21
// 	fmt.Println(z)
// }

// var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

// var z string = "Shaken, not stirred"

// var a string = `James said,
// "Shaken,

// not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

// func main() {
// 	fmt.Println(y)
// 	fmt.Printf("%T\n", y)
// 	fmt.Println(z)
// 	fmt.Printf("%T\n", z)
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)
// 	// z = 43
// 	// fmt.Println(z)
// 	// fmt.Printf("%T\n", z)
// }

// [zero values => it will take default value of zero]
// var y string
// var z int

// func main() {
// DECLARE a variable to be of a certain TYPE
// and then ASSIGN a VALUE of that type to the variable

// 	fmt.Println("printing the value of y", y, "ending")
// 	fmt.Printf("%T\n", y)

// 	y = "Bond, James Bond"

// 	fmt.Println(y)
// 	fmt.Printf("%T\n", y)

// 	fmt.Println(z)
// 	fmt.Printf("%T", z)
// }

//[formated printing]
// var a int
// var b string = "James Bond"
// func main() {
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%#v\n", a)
// 	fmt.Printf("%#v\n", b)
// 	fmt.Printf("%T\n", a)
// 	fmt.Printf("%T\n", b)
// 	fmt.Printf("%T\t%T\n", a, b)

// 	s := fmt.Sprint(a, " something more ", b)
// 	fmt.Println(s)
// 	s2 := fmt.Sprintf("%v\t%T\t%T\n", "to pass in", a, b)
// 	fmt.Println(s2)
// }

//[formated printing 2]
// var y = 42

// func main() {
// 	fmt.Println(y)
// 	fmt.Printf("%T\n", y)
// 	fmt.Printf("%b\n", y)
// 	fmt.Printf("%x\n", y)
// 	fmt.Printf("%#x\n", y)
// 	y = 911
// 	fmt.Printf("%#x\n", y)
// 	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

// 	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
// 	fmt.Println(s)
// 	fmt.Printf("%v", y)
// }

//[creating own type]
// var a int

// type hotdog int

// var b hotdog

// func main() {
// 	a = 42
// 	b = 43
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)
// 	fmt.Println(b)
// 	fmt.Printf("%T\n", b)
// }

//[conversion - type]
// var a int
// type hotdog int
// var b hotdog
// func main() {
// 	a = 42
// 	b = 43
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)
// 	fmt.Println(b)
// 	fmt.Printf("%T\n", b)
// 	a = int(b) //convertiong to int from hotdog
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)
// }
