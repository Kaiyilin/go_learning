# Go setup

## Go runtime download

## Checking the go version
```bash
$ go version
```

## Execute a basic go programme
```go
// helloWorld.go
package main
import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
```

```bash
$ go run ./helloWorld.go
Hello World
```
## Build the programme as executable
```bash
$ go build ./helloWorld.go
./helloworld
```

## Go programming structure
* In go, every programme is part of the package, we use ***package*** to define this
* ***import*** allows us to import files within certain package
* {}, code inside the curly bracket is executable

# Go Variables
## data types
- integer
- float
- bool
- string

## Declaring variables in 2 ways
1. With the ***var*** keyword
2. With the := sign

```go
// example
var variable_name type = value

// Note: In this case, the type of the variable is inferred from the value.
// Note: It is not possible to declare a variable using :=, without assigning a value to it.
// := can only be use in the function
variable_name := value

// It is also possible to assign a value to a variable after it is declared.

package main
import ("fmt")

func main() {
  var student_name string // declare first 
  student_name = "myName" // assign later
  fmt.Println(student_name)
}
```

## Go declaring multiple variables
```go
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```

## Go Variable Naming Rules
A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

# Like C, Go can declare constant

```go
// syntax
const CONSTNAME type = value

// e.g.
const myName string = "Kai"
```
## const rules
- Constant names follow the same naming rules as variables
- Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
- Constants can be declared both inside and outside of a function

## types of const
- Typed constants (declared with a defined type)
- Untyped constants (declared without a type)

### Note for const
- When a constant is declared, it is not possible to change the value later
- Multiple constants can be grouped together into a block for readability

```go
package main
import ("fmt")

const (
  A int = 1 // typed
  B = 3.14 // untyped
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```

# Go data structures
## Go Arrays

Arrays are used to store multiple values of the **same type** in a single variable

Arrays are fixed in size.

```go
var array_name = [length]datatype{values} // here length is defined
var array_name = [...]datatype{values} // here length is inferred

or you can use := as well

array_name := [length]datatype{values}
array_name := [...]datatype{values}

// e.g.
var cars_branding = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

var cars_branding = [...]string{"Volvo", "BMW", "Ford", "Mazda"}
```

## Go slices
slices are similar to arrays, slices also used to store multiple values of the **same type** in a single variable.

Slices can grow and shrink.

In Go, there are two functions that can be used to return the length and capacity of a slice:

- len() function - returns the length of the slice (the number of elements in the slice)
- cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

```go
package main
import ("fmt")

func main() {
  // syntax 
  // slice_name := []datatype{values}

  my_empty_slices := []int{}
  my_slices := []int{1, 2, 3}

}
```
### create slices with the make function

```go
// syntax
// if capacity is empty, it will automatically equal to length
slice_name := make([]type, length, capacity)

package main
import ("fmt")

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}
```
