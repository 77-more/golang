/*package main
import "fmt"
func main(){
	fmt.Println("Hello world")
}
*/
/*
package main 
import (
	"fmt"
	"os"
)
func main(){
	// declaration of variable
  var s string 
  // os.Args represents the arguments provided through the command line, in the example we ran the code as ./first test . Here ./first will be os.Args[0] and os.Args[1] will be test.  
  for i := 1; i<len(os.Args); i++ {
	s += os.Args[i]
	fmt.Println(os.Args[i])
  }
  fmt.Println(s)
}
*/
// a while loop can be written as follows 
// if you don't give any value to integer variable the default value is 0.
// if you don't give any value to a string varuable the default value is an empty space. 
/*
package main
import (
	"fmt"
)
func main() {
var i int
var s string
i = 0
fmt.Println(i)
fmt.Println(s)
for i < 5 {
	fmt.Println("Hello world");
	i +=1;
}
}*/
/*
// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
	// range will print the index value and the value 
	// _ element is called blank identifier. Golang has a special feature to define and use the unused variable using Blank Identifier. 
	// go doesn't allow unused variables 
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

below will show the different ways to declare the variables 
s := ""   ----> This way can be used only within a function
var s string
var s = ""
var s string = ""
*/
/*
package main

import (
    "fmt"
    "os"
	"strings"
)
func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
*/
/*
Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.

Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.

Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.) 
*/
// Below code prints the index and value of each element in os.Args.
/*
package main 
import (
    "fmt"
    "os"
)
func main(){
     for i,v := range os.Args{
        fmt.Println("index:",i,"value:",v)
     }
}
*/
// Below code prints the time taken by the program
// Println will print and start a new line after printing it its like using \n at the end of what you want to print. 
/*
package main
import (
    "fmt"
    "os"
    "time"
) 

func main(){

    start := time.Now()
    for i,v := range os.Args{
        fmt.Println("index:",i,"value:",v)
     }
    fmt.Printf("%.10f s elapsed\n", time.Since(start).Seconds())
}
*/
/*
variables and constants:
variables can have different values throughout the program but constants cannot have different values, the values are constant and the golang will not allow you to change the values. 

Declaring an integer type variable:

var a = 5

Declaring a string type variable:

var a = "test string"
*/
/*
printf:

with printf you can use placeholders for your variables. %v can be used for any type of 

> var a = 5

> var b = "test string"
> import "fmt" 

> fmt.Printf("printing %v and %v\n",a,b)
printing 5 and test string

You need to use \n with the printf statement because it is not included automatically. 
*/

/*
Data types:
In go if you don't assign the value while declaring the variable then it is gonna error out. So either declare the value or declare the type of the variable like below. 

var testVar string

To detect the type of the variable you can use the below line of code. 

printf("variable a is %T and variable b is %T",variable1,variable2)

If you want to explicitly declare the variable type you can use do the below. 

var testVar int = 5

Alternative way to declare a variable

testVar := 5

The above syntax cannot be followed for declaring constants
*/

/*
Asking user for the input:

fmt.Scan(testVar)
The above line will not stop for the user to enter the input. The below line will 
fmt.Scan(&testVar)


using & in front of the variable will print the memory location of the variable.  & is nothing but a pointer
fmt.Println(&testVar)


*/

/*
Arrays and slices in Go

How to define an array? 
The number 50 in the brackets is nothing but how many elements this array can hold. In this case 50. 
Next is the type of the elements that are stored in this variable. Here is it string. 
Next comes in the curl brackets are the actual elements. If you just declare {} this will initiate empty array. 
You cannot mix and match type of elements in array. If you say string everything has to be string. 

var bookings = [50]string{"Nana","Nicole","Peter"}
var bookings [50]string

adding elements to array. 

if you try to add 51st element it will error out because you have set the size of the array. So we use slices to overcome this issue. 

declaring a slice:

bookings := []string{}

Appending to a slice:

bookings = append(bookings, firstName+ " "+ lastName)

adding value to an array:

bookings[0] = firstName + " " + lastName
*/

package main 
import "fmt"
var p = f()
var a = 1 



func f() *int {
    v := 1
    return &v
}

func main(){
    fmt.Println(p)
    fmt.Println("this is a",a)
    var b = &a // This is how you turn b in to a pointer of variable a by writing &a 
    fmt.Println("this is b value",*b) // this is how you get the actual value that is present in the address location that variable b holds. 
}