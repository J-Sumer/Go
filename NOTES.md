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
