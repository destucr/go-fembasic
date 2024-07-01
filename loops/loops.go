package main 

import (
  "fmt"
  "slices"
)

func main() {
  //Array in go is fixed size and the size should be initlaize first
  //Thus you cant append an array to have a new length
  animals := [2]string{}

  animals[0] = "dog"
  animals[1] = "cat"
  
  fmt.Println("Print the array of animals")
  fmt.Println(animals)
  
  //If you want a dynamic size array, use slices
  //Declare like array but you dont initialize the size
  
  //slices
  foods := []string{
    "cake",
    "chips",
  }
  
  fmt.Println("Print the foods slices")
  fmt.Println(foods)

  //Append the value of slices
  foods = append(foods, "grape")

  fmt.Println("Print the appended slices")
  fmt.Println(foods)

  //Go couldnt delete specific index in slices
  //i.e you cant do delete(array[2]) naturally
  //Go 1.21 version above you can use "slices" stdlib

  //delete the index [1] in foods which is chips
  //we want delete 1
  //and inclusive till 2 (index 2 not deleted)
  foods = slices.Delete(foods,1,2)

  fmt.Println("Print the deleted slices")
  fmt.Println(foods)
  
  //add snacks
  foods = append(foods, "snacks")

  //Next iterate through array with loops
  //Print foods array with for loops
  fmt.Println()
  fmt.Println("Print foods array with for loop")
  for i := 0; i < len(foods); i++{
    fmt.Printf("This is my foods %s\n", foods[i])
  }

  //Prior go ver 1.22, you can use for range
  fmt.Println()
  fmt.Println("Print foods array with for range")
  for index, values := range foods {
    fmt.Printf("This is my foods index %d values %s\n", index, values)
  }

  //or iterate the range with an integer
  for value := range 10 {
    fmt.Println(value)
  }

  //Go doesnt have while keyword
  //Instead use for but with conditional
  //loops until is false
  fmt.Println("\nUsing while loop flow control")
  whileBool := 0

  for whileBool < 10 {
    fmt.Printf("valid %d\n",whileBool)
    whileBool ++
  }
}
