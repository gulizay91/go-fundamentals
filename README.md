# Go Fundamentals

## CH1- Compiled vs Interpreted
### Compiled Languages
Compiled languages are converted directly into machine code. So, they tend to be faster and efficient to execute than interpreted languages. Compiled programs can be run without access to the original source code, and without access to a compiler. They also give the developer more control over hardware aspects, like memory and CPU.

C#, C, C++, Java, Rust, Go

### Interpreted Language
Interpreters run through a program line by line and execute each command. Interpreted languages were once significantly slower than compiled languages.

Python, Ruby, Javascript, Php

### Strongly Typed and GC
Go is strongly and statically typed. Strongly typed means that variable typing is strictly enforced. Go's garbage collection automatically releases memory that is no longer needed by a program. This is called the Go Runtime which uses a garbage collector to periodically scan the program's heap, looking for objects that are no longer reachable. 

## CH2- Variables and Declarations
* string
```
var strValue string
var anotherStrValue string = "string value"
```
* bool
```
var boolValue, anotherBoolValue = true, false 
```

* numeric types:
    int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
    float32, float64
    complex64, complex128
```
age, pi := 23, 3.14
```

### Inital variables main.go
```sh
package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
```
result:
```
0 0.000000 false ""
```
### String length
Unicode characters can be 1 to 4 bytes

```sh
package main

import "fmt"

func main() {
    var ( 
      name string = "güliz"
      surname string = "ay"
    )
	
    fmt.Println(len(name)) // result: 6 (g: 1, ü: 2, l: 1, i: 1, z: 1 byte)
    fmt.Println(len(surname)) // result: 2
}
```

### Casting variables

```sh
package main

import "fmt"

func main() {
	accountAge := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer
	accountAgeInt := int8(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
```

### Formatting strings

fmt.Printf - Prints a formatted string to standard out.
fmt.Sprintf() - Returns the formatted string

EXAMPLES
These formatting verbs work with both fmt.Printf and fmt.Sprintf.

* %v Interpolate the default representation
```
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old
```
```
s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old
```
* %s Interpolate a string
```
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old
```
* %d Interpolate an integer
```
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
```
* %f Interpolate a decimal
```
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```

### Const and IOTA

```sh
package main

import "fmt"

const (
    Monday = iota + 1
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

func main() {
	// iota
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
```

```
1 2 3 4 5 6 7
```
### Conditionals
```
if numberOne < numberTwo {
		fmt.Println("Number one is smaller than numer two")
	}
```

```sh
package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
```

```
Trying to send a message of length: 10 and a max length of: 20

Message sent
```

## CH3- Functions

```sh
package main

import (
    "fmt"
    "strconv"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func add(i1 int, i2 int) int {
	return i1 + i2
}

func printer(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {
	printer("Lane,", " happy birthday!")
	printer("Your age is now ", strconv.Itoa(add(15,15)))
}

```

```
Lane, happy birthday!
Your age is now 30
```