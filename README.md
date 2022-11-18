# Golang Beginner Workshop @ DevFest 2022

👨‍💻 __Christian Bargmann (cbrgm)__

🐘 @chrisbargmann@det.social

📧 chris@cbrgm.net

---

## How to execute Go code?

* "I have Go installed on my computer" -> Awesome! I show you in a minute..
* "I don't have Go on my computer" -> Try [Go Playground](https://play.golang.org/)
* "I don't want to install Go on my computer" -> Try [Go Docker Image](https://hub.docker.com/_/golang)

---
## Session 01 - Basics Syntax (~45min)

* constants and variables
* for loops and use if conditions
* how types work.
* functions
* packages and imports

---
## Session 02 - Advanced Syntax (~45min)

* how zero values work in Go
* Equality and Type Conversions
* Pointers
* Arrays
* Slices
* Ranges
* Maps
* Structs
* Defer
* Concurrency

---
## Session 03 - Development (until the end)

* Go tool
* Go Modules
* Your first application
    * Hello World
    * What Time is it?
    * Simple Http Server
* Testing
* Error Handling

---
## Online Resouces:

* [Dave Cheney - Resources for new programmers](https://dave.cheney.net/resources-for-new-go-programmers)
* [tour.golang.org](tour.golang.org)
* [Go by Example](https://gobyexample.com/)

---
# Session 01 - Basic Syntax (~45 min)

We will first have a look at the basic Go syntax. You will learn:

* how to declare constants and variables
* how to write for loops and use if.
* how types work.
* how to write your own functions.
* how packages and import statements work.

---
## Introduction

The Go programming language

* Modern (2009) (well, compared to others :D)
* Compact, general-purpose
* Imperative, statically type-checked, dynamically type-safe
* Garbage-collected
* Compiles to native code, statically linked
* Fast compilation, efficient execution

> Designed by programmers for programmers!

---
##  Safety first, a few examples!

Example 1: Typed, and type safe

```go
var i int = -1
var u uint = 200
i = u   // nope, incompatible types
```

Example 2: Array accesses are bounds checked

```go
s := make([]string, 10)
x := s[20] // will panic at runtime
```

Example 3: No implicit conversions; booleans and integers are not aliases

```go
i := 2
if i { ... }    // nope, no coercion to bool
```

---
## Awesome support for concurrency

* Multicore CPUs are a reality.
* Multiprocessing is not a solution.
* Networking support baked into the standard library, integrated into the runtime.

---
## Garbage collected

Go is a garbage collected language.

* Eliminates the bookkeeping errors related to ownership of shared values.
* Eliminates an entire class of use after free and memory leak bugs.
* Enables simpler, cleaner, APIs.

The garbage collector handles heaps into the 100's of GB range!

---
## Opinionated

Go is an opinionated language.

* Unused local variables are an error.
* Unused imports are also an error.
* The compiler does not issue warnings, only errors.
* A single way to format code as defined by `go fmt`.

---
## Let's get started: Constants & Identifier

Here are some examples of constants:

```go
1
"hello"
false
1.3
```

To make a constant, we declare one with the const keyword.

```go
const name = "Christian"
println(name)
```

---
## Let's get started: Constants & Identifier

In Go an identifier is any word that starts with a letter.
All source code in Go is UTF-8, you can use even Emojis!

In Go an identifier is any word that starts with a letter.

```go
const participants = 22
const workshop = "DevFest! 😁"
println(participants, workshop)
```

Identifiers must start with a unicode letter, or underscore `_`. This does not work:

```go
const 1student = "Christian"
const 😁workshop = "go-workshop"
```

---
## Comments

* Inline comments, which start with a double forward slash, //.
* Block comments, which start with a forward slash and a star, /*, and end with a star and forward slash, */.

```go
// const a = 1

/*
const b = 2
const c = 3
*/

println(a, b, c)
```

---
## Declarations

There are six kinds of declarations in Go

* `const`: declares a new constant.
* `var`: declares a new variable.
* `type`: declares a new type.
* `func`: declares a new function, or method.
* `package`: declares the package this .go source file belongs to.
* `import`: declares that this package imports declarations from another.

---
## Variables

* A variable holds a value that can be changed over time.
* You declare a new variable with the `var` declaration.

Example:
```go
var π = 3.14159
var radius = 6371.0 // radius of the Earth in km
var circumference = 2 * π * radius

println("Circumference of the earth is", circumference, "km")
```

---
## Variables

Variables can be assigned in different ways:

```go
var x int         // Variable x of type int with a zero value
var x int = v     // Variable x of type int with value v
var x = v       // Variable x with value v, implicit typing

x := v          // Short variable declaration (type inferred)
x, y := v1, v2  // Double declaration (similar with var)
```

* Unused variables are often the source of bugs.
* If you declare a variable in the scope of your function but do not use it, the Go compiler will complain.

---
## Variables

* See: `01_variables/task_00.go`

```go
package main

func main() {
	// Todo: Fix this program. Uncomment the lines below, delete or comment out "others"
	var students = 50
	var assistants = 15
	var professors = 5
	var others = 1

	// println(students + assistants)
	// _ = professors
}
```

---
## Incrementing / Decrementing

* Go supports a limited form of variable post-increment and post-decrement
* Examples: x++, x--, there is no pre-de/incrementing
* i++ and i-- are statements, not an expressions, they do not produce a value.

```go
var i = 1
i++
println(i)
```

---
## Incrementing / Decrementing

See: `01_variables/task_01.go`

```go
package main

func main() {
	var i = 1

	// Task: Fix the statement by moving i++ above the declaration
	var j = i++
	println(i, j)
}
```

---
## Looping

Go has a single for loop construct that combines

    while condition { … }
    do { … } while condition
    do { … } until condition

into one syntax. There is only `for` in Go:

```go
for (init statement); condition; (post statement) { … }
```

---
## Looping


```go
for (init statement); condition; (post statement) { … }
```

The parts of a for statement are:

* **init statement**: used to initalise the loop variable; i = 0.
* **condition**: user to test if the loop is done; i < 10, true means keep looping.
* **post statement**: user to increment the loop variable; i++, i = i - 1.

---
## Looping

Our first loop:

```
var i = 0
for i = 0; i < 10; i++ {
    println(i)
}
```

Cool: not need to put ( braces around the for condition )

---
## Looping

* See: `02_loops/task_00-02.go`

```go
package main

func main() {
	// Task: Can you move line 6 into the loop header? What's the difference?
	var i = 0

	for i = 0; i < 10; i++ {
		println(i)
	}
}
```

---
## Looping

* See: `02_loops/task_00-02.go`.
* Here's another loop example:

```go

package main

func main() {
	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		println(i)
		i += 2
	}
}
```

---
## Conditions

* Go has two conditional statements, `if` and `switch`.
* if is used to choose between two choices based on a condition:

```go
if v > 0 {
        println("v is greater than zero")
} else {
        println("v is less than or equal to zero")
}
```

---
## Conditions

* The `else` part may be omitted, for example when checking preconditions:

```go
if v == 0 {
        // nothing to do
        return
}
// handle v
```

---
## Conditions

* See: `03_conditions/task_00.go`
* Here's another conditions example:

```go

package main

func main() {
	// print only even numbers!
	var i = 0
	for i = 1; i < 11; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
}
```

---
## Continue & break

* Unlike languages like Java, if statements in Go are often used as guard clauses.
* We say that when everything is true the code reads from the top to the bottom of the page.
* We can rewrite the previous program using a new statement, `continue`, which skips the body of the loop.
* Also we can use `break` to break out of loops, see: `03_conditions/task_03.go`

---
## Continue & break

* See: `03_conditions/task_01.go`
* Breaking out of an infinite loop:

```go
package main

func main() {
	// Task: Uncomment the lines below, we can break out of our infite loop!
	var i = 1
	for {
		// if i > 10 {
		//    break
		// }
		if i%2 == 0 {
			println(i)
		}
		i++
	}
}

```

---
## Switch Statement

* A switch statement runs the first case equal to the condition expression
* The cases are evaluated from top to bottom, stopping when a case succeeds.
* If no case matches and there is a default case, its statements are executed.

```go
var c rune = 't'

switch c {
    case ' ', 't', 'n', 'f', 'r':
        println("Found!")
    default:
        println("Not found!)
}
```

---
## Switch Statement

* A fallthrough statement transfers control to the next case.
* Example below results in `2,3`

```go
switch 2 { // missing expression means "true"
    case 1:
        fmt.Println("1")
        fallthrough
    case 2:
        fmt.Println("2")
        fallthrough
    case 3:
        fmt.Println("3")
}
```

---
## Types

Go is a strongly typed language, like Java, C, C++, and Python. Go has nine kinds of types, they are:

* `strings`: string.
* `signed integers`: int8, int16, int32, int64.
* `unsigned integers`: uint8, uint6, uint32, uint64.
* `aliases`: byte, rune, int, uint.
* `booleans`: bool.
* `IEEE floating point`: float32, float64.
* `Complex types`: complex64, complex128.
* `Compound types`: array, slice, map, struct.
* `Pointer types`: *int, *bytes.Buffer.

---
## Functions

* You can declare your own functions with the func declaration.
* A function's name must be a valid identifier, just like const and var.
* A simple function definitions looks like this:

```go
func foo() {} // Simple function definition
```

Example:

```go
func moin() {
    println("Moin")
}
```

* Functions may have parameters and return types.
* To pass an argument to a function, the type of the argument and the type of the function's formal parameter must be the same.

---
## Functions

* Let's have a look at `04_functions/task_00-01.go`

```go
package main

func moin() {
	println("moin")
}

func main() {
	var i int = 0
	for i = 1; i < 3; i++ {
		moin()
	}
}

```

---
## Functions

There are many different ways to define a function:

```go
func f1() {}                // Simple function definition
func f2(s string, i int) {} // Function that accepts two args
func f3(s1, s2 string) {}   // Two args of the same type
func f4(s ...string) {}     // Variadic function

func f5() int {             // Return type declaration
	return 42
}

func f6() (int, string) {   // Multiple return values
	return 42, "foo"
}
```

---
## Packages

* A package is the unit in which software is shared and reused in Go. All Go code is arranged into packages.
* Each source file in a package must begin with the same package declaration.
* A package's name must be a valid identifier, just like `const`, `var`, and `func`.
* We already introduced the `main` package. To run a Go program it needs a `main()` function within a `main package`

---
## Packages

* See: `05_packages/task_00.go`

```go
// Task: Make this program runnable, change the package name to main
package testpackage

func add(x int, y int) int {
	return x + y
}

func main() {
	println(add(42, 13))
}
```

---
## Packages

 `Go` ships with a rich standard library of packages. This includes

* file input / output
* string handling
* compression
* encoding and decoding of JSON and XML
* network handling
* HTTP client and server

---
## Imports

* The final declaration we'll cover in this section is the `import` declaration.
* The import declaration allows you to use code from other packages into your package.
* When you import a package, the public types, functions, variables, types, and constants, are available with a prefix of the package's name.
* Go already ships a lot of default packages

```go
time.Now    // denotes the Now function in package time
```

You can import different packages using `import`.

```go
import "fmt"
import "time"
```

---
## Imports

You can also write:

```go
import (
    "fmt"
    "time"
)
```

---
## Imports

Here's a full example:


```go
package main

import "fmt"
import "time"

func main() {
    var now = time.now()
    fmt.println(now)
}
```

---
## End of Section 1 - Time for a recap! :-)


What we've learned so far:

* how to declare constants and variables
* how to write for loops and use if.
* how types work.
* how to write your own functions.
* how packages and import statements work.

See you in ~15 min.

---
# Session 02 - Advanced Syntax (~45 min)

Welcome back! In the second session we will cover:

* how zero values work in Go
* Equality and Type Conversions
* Pointers
* Arrays
* Slices
* Ranges
* Maps
* Structs
* Defer
* Concurrency

---
## Zero Values

* In Go, there is no unitialised memory. The Go runtime will always ensure that the memory allocated for each variable is initalised before use.
* For example: `var name string` Then the memory assigned to the variable `name` will be zeroed, as we have not provided an initaliser.
* Result: Value of `name` will be "" because it's the value f a string with zero length

---
## Zero Values

Zero values summarized:

* The zero value for integer types: int, int8, uint, uint64, etc, is 0.
* The zero value for floating point types: float32, float64, complex128, etc, is 0.0.
* The zero value for arrays is the zero value for each element, ie. [3]int is 0, 0, 0.
* The zero value for slices is nil.
* The zero value for boolean is false.
* The zero value for structs is the zero value for each field.

---
## Equality & Type Conversions

* As Go is a strongly typed language, for two variables to be equal, both their type and their value must be equal.
* See: `00_equality/task_00.go`

```go
package main

import "fmt"

// In this example the assignment of y = x fails because x and y are different integer types.
// This will fail:
func main() {
	var x uint = 700
	var y int = 700
	fmt.Println(x == y)
}

```

---
## Equality & Type Conversions

* Sometimes you have variables of different integer types, you can convert from one type to another using a conversion expression.
* The expression `T(v)` converts the value v to the type T.

Example:

```go
package main

import "fmt"

func main() {
    var x uint = 700
    var y int = int(x) // Type Conversion from uint to int

    fmt.Println(y)
}
```

* When you convert a variable with a larger number of bits to a variable with a smaller number of bits there is a risk of truncation, because there are less bits available to represent your number.

---
## Equality & Type Conversions

* See: `01_types/task_00-01.go`
* Here's another example:

```go
package main

import "fmt"

// In this example the assignment of y = x fails because x and y are different integer types.
// This will fail:
func main() {
	var x uint = 700
	var y int = 700

	fmt.Println(x == y)
}
```

---
## Pointers

* Go has pointers. A pointer holds the memory address of a value.
* The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
* Example: `var p *int`
* The `&` operator generates a pointer to its operand.

```go
i := 42
p = &i
```
* The * operator denotes the pointer's underlying value. Example:

```go
i := 42
p = &i

fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

---
## Pointers

* See: `02_pointers/task_00.go`
* Again, another example

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer

	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

---
## Arrays

* An array type definition specifies a length and an element type.
* For example, the type `[4]int` represents an array of four integers. An array's size is fixed; its length is part of its type.
* Example: `[4]int` and `[5]int` are distinct, incompatible types
* Arrays can be indexed in the usual way, so the expression s[n] accesses the nth element, starting from zero.

Example:

```go
var a [4]int
a[0] = 1
i := a[0]

// i == 1
```

---
## Arrays

To initialize a string array you can write:

```go
b := [2]string{"Christian", "Bargmann"}
```

---
## Arrays

Or, you can have the compiler count the array elements for you:

```go
b := [...]string{"Christian", "Bargmann"}

```

---
## Arrays

See: `03_types/task_00.go`

```go
package main

import "fmt"

// Example: string and int arrays
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

---
## Slices

* Slices are like references to arrays
* Go's slice type provides a convenient and efficient means of working with sequences of typed data
* A slice literal is declared just like an array literal, except you leave out the element count:
* Example: `letters := []string{"a", "b", "c", "d"}`
* A slice is formed by specifying two indices, a low and high bound, separated by a colon:  `a[low : high]`

---
## Slices

See: `04_slices/task_00.go`

```go

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
// Result: [3 5 7]
```

* A slice does not store any data, it just describes a section of an underlying array.
* Changing the elements of a slice modifies the corresponding elements of its underlying array.

---
## Slices

* See: `04_slices/task_01.go`

```go
package main

import "fmt"

// Task: First run the program. Uncomment the line below afterwards. What behavior do you see? Any ideas?
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

---
## Slice Literals

A slice literal is like an array literal without the length.

Example: This is an array literal:

```go
[3]bool{true, true, false}
```

Example: This is a slice literal. It creates the same array as above, then builds a slice that references it:

```go
[]bool{true, true, false}
```

---
## Slice Literals

You can also use Go's build in `make` function to create a new slice:

```go
package main

import "fmt"

func main() {
    i := make([]int, 20)
    fmt.Println(len(i))
}
```

---
## Length & Capacity

* A slice has both a **length** and a **capacity**.
* The length of a slice is the number of elements it contains.
* The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
* The length and capacity of a slice s can be obtained using the expressions `len(s` and `cap(s)`
* The zero value of a slice is `nil`. A `nil` slice has a length and capacity of 0 and has no underlying array.  Example: `var s []int`


---
## Length & Capacity

* See: `04_slices/task_03.go`

```go
package main

import "fmt"

// Example: Comparison of slice length vs capacity
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

---
## Appending to slices

* It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
* The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
* If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

---
## Appending to slices

See: `04_slices/task_04.go`

```go
package main

import "fmt"

// Example: appending items to a slice
func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

```

---
## Range

* The range form of the for loop iterates over a slice or map.
* When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

Example:

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2*%d = %d\n", i, v)
	}
}
```

---
## Range

* You can skip the index or value by assigning to `_`.
* If you only want the index, drop the , value entirely.
* See: `04_range/task_05.go`

---
## Maps

* Go has a built in Hash Map type, called a `map`.
* Maps map values of key type K to values of type V
* Example: `var m map[string]int`
* Just like making a slice, making a map is accomplished with the `make` built-in.

---
## Maps

Example:

```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    fmt.Println(m)
}
```

---
## Maps

Inserting a value into a map looks similar to assigning a value to a slice element.

```go
package main

import "fmt"

func main() {
    days := make(map[int]string)
    days[1] = "Monday"
    days[2] = "Tuesday"
    days[3] = "Wednesday"
    days[4] = "Thursday"
    days[5] = "Friday"
    days[6] = "Saturday"
    days[7] = "Sunday"
    fmt.Println(days)
}

```

---
## Maps

 You can also use the sort literal initialization:

 ```go
package main

import "fmt"

func main() {
     days := map[int]string{
         1: "Monday",
         2: "Tuesday",
         3: "Wednesday",
         4: "Thursday",
         5: "Friday",
         6: "Saturday",
         7: "Sunday",
     }
     fmt.Println(days)
 }
 ```

---
## Retrieve values from a map

 Just like a slice, you can retrieve the value stored in a map with the syntax m[key].
 Example:

```go
package main

import "fmt"

func main() {
    // map of planets to their number of moons
    moons := map[string]int{
        "Mercury": 0,
        "Venus":   0,
        "Earth":   1,
        "Mars":    2,
        "Jupiter": 67,
    }

    fmt.Println("Earth:", moons["Earth"])
    fmt.Println("Neptune:", moons["Neptune"])
}

```

---
## Check if value in map exists

Map look ups support a second syntax.
Example:

```go
package main

import "fmt"

func main() {
    moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

    n, found := moons["Earth"]
    fmt.Println("Earth:", n, "Found:", found)

    n, found = moons["Neptune"]
    fmt.Println("Neptune:", n, "Found:", found)
}
```

---
## Delete a value from map

To delete a value from a map, you use the built in delete function.
Example:

```go
package main

import "fmt"

func main() {
    moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

    const planet = "Mars"

    n, found := moons[planet]
    fmt.Println(planet, n, "Found:", found)

    delete(moons, planet) // Removing element from map

    n, found = moons[planet]
    fmt.Println(planet, n, "Found:", found)
}
```

---
## Iterating over a map

We can use the `range` operator presented earlier :-)

```go
for key, value := range myMap {
    fmt.Println("Key:", key, "Value", value)
}
```

---
## Structs

* A struct is a collection of fields. They allow us to create our "own" types in Go.
* For example we can create a Person Type like this:

```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

```

* The `type` keyword introduces a new type.
* It is followed by the name of the type (Person) and the keyword struct to indicate that we’re defining a struct.
* The struct contains a list of fields inside the curly braces. Each field has a name and a type.

---
## Creating a struct

You can initialize a variable of a struct type using a struct literal like so:

```go
// Initialize a struct by supplying the value of all the struct fields.
var p = Person{"Christian", "Bargmann", 25}
```

* Note: Note that you need to pass the field values in the same order in which they are declared in the struct (This is not Python! :-))

---
## Creating a struct

Go also supports the name: value syntax for initializing a struct (the order of fields is irrelevant when using this syntax).

```go
// Initialize a struct by supplying the value of all the struct fields.
var p = Person{FirstName: "Christian", LastName: "Bargmann", Age: 25}
```

or like this:

```go
var p = Person{
    FirstName: "Christian",
    LastName: "Bargmann",
    Age: 25,
}
```

* You can also use the built-in `new()` function to create an instance of a struct.
* The `new()` function allocates enough memory to fit all the struct fields, sets each of them to their zero value and returns a pointer to the newly allocated struct.

---
## Creating a struct

See: `06_structs/task_01.go`

```go
package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

// Example: Create a new employee using build-in new() function
func main() {
	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	employeePointer := new(Employee)

	employeePointer.Id = 1000
	employeePointer.Name = "Sachin"

	fmt.Println(employeePointer)
}
```

---
## Accessing struct fields

* You can access individual fields of a struct using the dot (.) operator -

```go
c := Car{
	Name:       "Ferrari",
	Model:      "GTC4",
	Color:      "Red",
	WeightInKg: 1920,
}

// Accessing struct fields using the dot operator
fmt.Println("Car Name: ", c.Name)
fmt.Println("Car Color: ", c.Color)
```

---
## Accessing struct fields

You can also access struct fields through pointers

```go
c := Car{
	Name:       "Ferrari",
	Model:      "GTC4",
	Color:      "Red",
	WeightInKg: 1920,
}

pc := &c // Create a pointer to struct c

// Accessing struct fields through a pointer using the dot operator
fmt.Println("Car Name: ", pc.Name)
fmt.Println("Car Color: ", pc.Color)
```

---
## "OOP"

* Using pointer receivers, we can do stuff that feels like "object orientated programming"
* Lets have a look at: `06_structs/task_02-03.go`

---
## Interfaces

* An interface type is defined as a set of method signatures.
* A value of interface type can hold any value that implements those methods.
* **Interfaces are implemented implicitly** - no need for `foo implements bar`
* A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

---
## Interfaces

```go
package main

import "fmt"

type Dense interface { // The Dense interface
  Density() int
}

// IsItDenser takes two interfaces as formal parameters
func IsItDenser(a, b Dense) bool {
  return a.Density() > b.Density()
}

type Rock struct {
  Mass   int
  Volume int
}

// Rock has Density() int method
func (r Rock) Density() int {
  return r.Mass / r.Volume
}

type Geode struct {
}

// Geode has Density() int method
func (g Geode) Density() int {
  return 100
}

func main() {
  r := Rock{10, 1}
  g := Geode{}

  fmt.Println(IsItDenser(g, r)) // We can pass both to IsItDenser, both implement Denser interface implicitly!
}

```

* See: `08_interfaces/task_00.go`

---
## The empty interface

Every type implements the empty interface `interface{}`. This allows us to pass any type to a function. Whoop!

```go
func PrintIt(a interface{}) {
  // a can have any methods. We dont care. :-)
  fmt.Println(a)
}

```

Since Go 1.18 you can also use `any` instead of `interface{}`

---
## Defer

* A defer statement defers the execution of a function until the surrounding function returns.
* The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
* It's like Java's try-with-resources :-)

Example:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// prints hello world
```

---
## Defer

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

```

---
## Defer

Returns:

```bash
counting
done
9
8
7
6
5
4
3
2
1
0
```

* `defer` is useful when opening resources like database connections or file syscalls :-)
* See: `09_defer/task_00.go`

---
## Concurrency

* Concurrent programming is a large topic but it’s also one of the most interesting aspects of the Go language.
* Do not communicate by sharing memory; instead, share memory by communicating.
* Go encourages a different approach in which shared values are passed around on channels and, in fact, never actively shared by separate threads of execution.

---
## Go Routines

* A goroutine is a lightweight thread managed by the Go runtime.
* Example: `go f(x, y, z)` starts a new goroutine running `f(x,y,z)`
* The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
* Goroutines run in the same address space, so access to shared memory must be synchronized.

---
## Channels

* Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

* Like maps and slices, channels must be created before use
* Example: `ch := make(chan int)`
* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
* See: `10_concurrency/task_00.go`

---
## Buffered Channels

* Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
* Example: `ch := make(chan int, 100)`
* Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
* See: `10_concurrency/task_01.go`

---
## Range and close

* A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
* Example: `v, ok := <-ch`
* ok is false if there are no more values to receive and the channel is closed.
* See: `10_concurrency/task_02-03.go`

---
## Select Statement

* The select statement lets a goroutine wait on multiple communication operations.
* A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
* See: `10_concurrency/task_04.go`

---
# End of section 02

In this section we convered:

* how zero values work in Go
* Equality and Type Conversions
* Pointers
* Arrays
* Slices
* Ranges
* Maps
* Structs
* Defer
* Concurrency


---
# What's next?

Do we have time left? There's plenty of stuff we did not talk about yet

* Generics
* Go tooling
* Go Modules
* Project Structure
...

what are you interested in? Let's have discussions!

---
# Thank you!

👨‍💻 __Christian Bargmann (cbrgm)__

🐘 @chrisbargmann@det.social

📧 chris@cbrgm.net
