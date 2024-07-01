package main

import "fmt"

func main() {
  // var myName string = "Melkey"
  
  //infered type
  newName := "Melkey too"

  fmt.Printf("Hello my name is %s\n", newName)

  //This show go zero value concept which is the default value
  //For undeclare variable value
  //Go would assign a default value
  var myBool bool
  var myInt int
  var myFloat float32

  fmt.Printf("mybool is %t\nmyint is %d\nmyFloatis %f\n",myBool,myInt,myFloat)
}
