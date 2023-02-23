## To create a module in go
$ go mod init example/hello 
	- initializes a go.mod files that tracks the dependencies requried for the the code

modules --contains--> packages 

$ go mod tidy
	- if we add any packages we need to import them in go.mod file. Running above command will do that for us (It will download the requried dependencies)

$ go mod edit -replace example.com/greetings=../greetings
	- to get the package locally
$ go mod tidy
	- to refresh the go.mod file and make the changes. synchronize the example.com/hello module's dependencies

main package and main func are go's starting point 

$ go build
	- to build the package (creates executable file ) (.exe in windows)
	- ./hello or ./hello.exe in windows

$ go list -f '{{.Target}}'
	- to get go installed path

$ set PATH=C:\path\to\your\install\directory

$ go install
	- command to compile and install the package

$ hello
	- to run the package from terminal 

	
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. pizza and pi do not start with a capital letter, so they are not exported.

```
func main() {
	fmt.Println(math.pi) // Throws an error
}
```

``` x int, y int  ```
is simillar to
``` x, y int ```

A function can return any number of results. The swap function returns two strings.

```
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

### Named returns

Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values.

```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

var is used to declare a list of variables

```
var c, python, java bool
var k, j int = 1, 2

func main() {
	var i int
	var c, python, java = true, false, "no!"
	fmt.Println(i, c, python, java) // 0 true false no!
}
```

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

Basic types in go

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128

```
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) \\ Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) \\ Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", z, z) \\ Type: complex128 Value: (2+3i)
}
``` 


Variables declared without an explicit initial value are given their zero value.

The zero value is:
* 0 for numeric types,
* false for the boolean type, and
* "" (the empty string) for strings.


The expression T(v) converts the value v to the type T.

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
More simply 
```
i := 42
f := float64(i)
u := uint(f)

```

### Constants
Constants are declared like variables, but with the const keyword. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the := syntax.

# Flow control

for loop (doesn't need () but should have {})
```
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

For is Go's "while"
```
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```
or
```
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

Like for, the if statement can start with a short statement to execute before the condition. Variables declared by the statement are only in scope until the end of the if.
```
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
```

# Complete this exercies https://go.dev/tour/flowcontrol/8

To declare int as float64
```
z := 1.0
z := float64(1)
```

Switch case. Switch without a condition is the same as switch true.

```
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

### Defer
A defer statement defers the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

counting
done
9
8
7
.
.
```

### Pointers

```
var p *int
i := 42
p = &i
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

While initilalizing we give star, when assigning we give & and when accessing its value we use * again.

```
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

### Structs
Struct fields are accessed using a dot.
```
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

If we don't specify the value it will go to default (like 0 for int, false for bool etc).
```
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```

### Arrays (Fixed length)
var a [10]int
```
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

### Slice (Dynamic length)

A slice gives flexible view into the elements of an array. In practice

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

``` a[low : high] ```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

``` a[1:4] ```

```
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

```

Slices are like references to arrays. A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.

```
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b) // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}
```

Slice literals
A slice literal is like an array literal without the length.

This is an array literal:

[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}

```
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

```
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]
}
```
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, ``` counting from the first element in the slice. ```

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

```
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)  // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

```
var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!") // This gets printed
	}
```

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

``` a := make([]int, 5)  // len(a)=5 ```
To specify a capacity, pass a third argument to make:
```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

Slices can contain any type, including other slices.

```
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

```
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


### Range for looping slice

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

```
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

Prints:
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128

```

You can skip the index or value by assigning to _.
```
for i, _ := range pow
for _, value := range pow
```
If you only want the index, you can omit the second variable.

``` for i := range pow ```

# Complete this exercise

https://go.dev/tour/moretypes/18

### Maps

m = make(map[<key type>]<Value type>)

```
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

```
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

```
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

```
func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.

```
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

### Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.

```
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

Print:
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

### Methods in GO

Go does not have classes. However, you can define methods on types. A method is a function with a special receiver argument. The receiver appears in its own argument list between the func keyword and the method name. In this example, the Abs method has a receiver of type Vertex named v.

```
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```


```
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

### Pointer receivers
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.


```
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// if * is removed, value printed is 5 not 50
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs()) // 50
}
```
Another way
```
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```


### Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
```
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

```
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:
```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```
while methods with value receivers take either a value or a pointer as the receiver when they are called:
```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
In this case, the method call p.Abs() is interpreted as (*p).Abs().

```
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
```

##### Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)


```
type Vertex struct {
	X, Y float64
}

// This * is needed
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// This * is not
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```





































