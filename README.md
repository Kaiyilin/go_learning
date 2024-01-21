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

# Go Variable Naming Rules
A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

<ul>
    <li> A variable name must start with a letter or an underscore character (_)
    <li> A variable name cannot start with a digit
    <li> A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
    <li> Variable names are case-sensitive (age, Age and AGE are three different variables)
    <li> There is no limit on the length of the variable name
    <li> A variable name cannot contain spaces
    <li> The variable name cannot be any Go keywords
</ul>

# Like C, Go can declare constant

```go
// syntax
const CONSTNAME type = value

// e.g.
const myName string = "Kai"

```
## const rules
<ul>
    <li>Constant names follow the same naming rules as variables
    <li>Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    <li>Constants can be declared both inside and outside of a function
</ul>

## types of const
<ul>
    <li>Typed constants (declared with a defined type)
    <li>Untyped constants (declared without a type)
</ul>

### Note for const
<ul>
<li> When a constant is declared, it is not possible to change the value later
<li> Multiple constants can be grouped together into a block for readability

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
</ul>