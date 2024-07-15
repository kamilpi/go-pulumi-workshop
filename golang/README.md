Download and install: https://go.dev/doc/install


### 1. What is Go?
- Go is a cross-platform, open source programming language
- Go is a fast, statically typed, compiled language
- Go was developed in 2007
- Go's syntax is similar to C++

### 2. Why Use Go?
- Go has fast run time and compilation time
- Go supports concurrency
- Go has memory management
- Go works on different platforms (Windows, Mac, Linux, Raspberry Pi, etc.)

### 3. Go Compared to Python

| Go  | Python  |
|---|---|
| Statically typed  | Dynamically typed  |
| Fast run time  | Slow run time  |
| Compiled  | Interpreted  |
| Fast compile time  | Interpreted  |
| Supports concurrency through goroutines and channel  | No built-in concurrency mechanism  |
| Has automatic garbage collection  | Has automatic garbage collection  |
| Does not support classes and objects  | Has classes and objects  |
| Does not support inheritance  | Supports inheritance  |

### 4. Syntax

<pre>
package main // - <b>Package declaration</b>
import ("fmt") // - <b>Import packages</b>

func main() { // - <b>Functions</b>
  fmt.Println("Hello World!") // - <b>Statements and expressions </b>
}
</pre>

> **_NOTE:_**  In Go, any executable code belongs to the main package.


### 5. Comments

#### Single-line
<pre>
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!") <b>// This is a comment</b>
}
</pre>

#### Multi-line

<pre>
package main
import ("fmt")

func main() {
  <b>/* The code below will print Hello World
  to the screen, and it is amazing */</b>
  fmt.Println("Hello World!")
}
</pre>

### 6. Variables

#### Types:
- <b>int</b>- stores integers (whole numbers), such as 123 or -123
- <b>float32</b> - stores floating point numbers, with decimals, such as 19.99 or -19.99
- <b>string</b> - stores text, such as "Hello World". String values are surrounded by double quotes
- <b>bool</b> - stores values with two states: true or false


#### Declaration:
- with the `var` keyword:
<pre>
var component string = "Jenkins"
</pre>
- with the `:=` sign:
<pre>
component := "Jenkins"
</pre>

#### Differences
| var  | := |
|---|---|
| Can be used inside and outside of functions | Can only be used inside functions |
| Variable declaration and value assignment can be done separately  | Variable declaration and value assignment cannot be done separately (must be done in the same line)|


#### Examples
<pre>
var a, b, c, d int = 1, 3, 5, 7
var a, b = 6, "Hello"
var (
    a int
    b int = 1
    c string = "hello"
)
</pre>

#### Rules
- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

#### Styles
- camelCase
- PascalCase
- snake_case

### 7. Constants
<pre>
const CONSTNAME type = value // - <b>READ ONLY!!!</b>
</pre>  

- rules as variables
- can be declared both inside and outside of a function
- names are usually written in uppercase letters

### 8. Output

- Print() - prints its arguments with their default format.
<pre>
package main
import ("fmt")

func main() {
  var i string = "Hello"
  fmt.Print(i)
}
</pre>

- Println() - function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end
<pre>
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
</pre>

- Printf() - formats its argument based on the given formatting verb and then prints them
<pre>
package main
import ("fmt")

func main() {
  var i string = "Hello"

  fmt.Printf("i has value: %v and type: %T\n", i, i)
}
</pre>

#### Formatting verbs examples

- %v - prints the value in the default format
- %s - prints the value as plain string
- %g - print float only necessary digits
- %d - integer - base 10
- %b - integer - base 2

### 9. Arrays
#### Declaration
<pre>
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred
</pre>

#### Change element of Array
<pre>
package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}
</pre>

#### len() function
<pre>
package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Println(len(arr1))
}
</pre>

### 10. Slices
> **_NOTE:_**  the length of a slice can grow and shrink as you see fit

create a slice:
- Using the []datatype{values} format
    ```
    myslice := []int{1,2,3}

    ```
- Create a slice from an array
    ```
    var myarray = [length]datatype{values} // An array
    myslice := myarray[start:end] // A slice made from the array
    ```
- Using the make() function
    ```
    slice_name := make([]type, length, capacity)
    myslice1 := make([]int, 5, 10)
    ```

### 11. Operators

- `+, -, *, /, %, ++, --`
- `==, !=, >, < >=, <=`
- `&&, ||, !`
- `&, |, ^, <<, >>`  
---
- `=, +=`
- `*, *=`

### 12. Conditions

- 
  ```
  func main() {
    if 20 > 18 {
      fmt.Println("20 is greater than 18")
    }
  }
  ```

- 
  ```
  if condition {
    // code to be executed if condition is true
  } else {
    // code to be executed if condition is false
  }
  ```

- 
  ```
  if condition1 {
    // code to be executed if condition1 is true
  } else if condition2 {
    // code to be executed if condition1 is false and condition2 is true
  } else {
    // code to be executed if condition1 and condition2 are both false
  }
  ```

- 
  ```
  if condition1 {
    // code to be executed if condition1 is true
    if condition2 {
      // code to be executed if both condition1 and condition2 are true
    }
  }
  ```

13. Switch

```
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}
```

14. Loops

```
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
```

```
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

15. Functions

- A function is a block of statements that can be used repeatedly in a program.

- A function will not execute automatically when a page loads.

- A function will be executed by a call to the function.

```
func FunctionName() {
  // code to be executed
}
```

```
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}
```

rules:
- A function name must start with a letter
- A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
- Function names are case-sensitive
- A function name cannot contain spaces

#### Parameters

```
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}
```

```
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
```

#### Return

```
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
```

```
package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}
```

#### Recursion

```
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}
```

### 13. Struct
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

```
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
```

```
type Person struct {
  name string
  age int
  job string
  salary int
}
```

- You can also pass a structure as a function argument
- To access any member of a structure, use the dot operator (.) between the structure variable name and the structure member.
  ```
  var pers1 Person
  pers1.name = "Hege"
  ```

### 14 Maps

- Maps are used to store data values in key:value pairs.
- Each element in a map is a key:value pair.
- A map is an unordered and changeable collection that does not allow duplicates.
- The length of a map is the number of its elements. You can find it using the len() function.
- The default value of a map is nil.
- Maps hold references to an underlying hash table.
- Go has multiple ways for creating maps.

```
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
```

```
package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}
```

### 15 Compilation
- The `go build` command compiles the packages, along with their dependencies, but it doesn't install the results.
- The `go install` command compiles and installs the packages.

### 16 Formatting

Gofmt is a tool that automatically formats Go source code.

```
gofmt -w yourcode.go
```

```
go fmt path/to/your/package
```



In Go language, the main package is a special package which is used with the programs that are executable and this package contains main() function. The main() function is a special type of function and it is the entry point of the executable programs.

init() function is just like the main function, does not take any argument nor return anything. This function is present in every package and this function is called when the package is initialized. This function is declared implicitly, so you cannot reference it from anywhere and you are allowed to create multiple init() function in the same program and they execute in the order they are created. 

# Exercises!
