package main

import (
  "fmt"
)

func main() {
  //The concept of pointers and what it means
  // asterisk * could mean a pointers and dereferencing
  a := 5
  b := &a

  //b value is a pointer to allocated memory of a
  fmt.Println(b)

  //change the value of b need to dereferencing first
  *b = 3

  //this would change the b value and a value
  fmt.Println(*b)
  fmt.Println(a)

  //Now let see a slices is actually passed value by copy
  slice := []int{1,2,3}
  fmt.Println(slice)

  //change slice value with for loop
  //this doesnt work
  for _,value := range slice{
    value++
  }
  
  fmt.Println("Changed slice value (doesnt work because of passed by copy)")
  fmt.Println(slice)
  
  //this work because the change is by reference i.e the slice index
  for index, _ := range slice{
    slice[index]++
  }

  fmt.Println("Change slice value")
  fmt.Println(slice)
}

