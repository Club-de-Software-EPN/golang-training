// First line of a Go program
package main

// import
import (
  "fmt" // format package
  "math"
)

// main function
func main(){

  fmt.Println("Hi World, EPN!")

  // Variables
  var name string = "Alejandro"
  lastname := "Llanganate"
  const PI float64 = 3.1415121 // JS

  // Zero Values // Declare variables
  var count float64 // 0
  var state bool // false
  var id string // ""
  var age int64 // 0

  fmt.Println(count, state, id, age)

  // Println
  fmt.Println(name, lastname)
  fmt.Println("Pi: ", PI)

  // Printf similiar C
  var team string = "Cyber security"
  var total int = 19

  fmt.Printf("%s --- total members: %d \n",team, total)

  // Square
  const anyNumber float64 = 4.0
  fmt.Println("Square of 4.0: ",math.Pow(anyNumber, 2))
}