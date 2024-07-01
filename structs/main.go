package main

import (
  "fmt"
)

type Person struct {
  Name string
  Age int
}

func newPerson(name string, age int) Person{
  return Person{
    Name: name,
    Age: age,
  }
}

func wrongChangeName(person Person, newName string){
  person.Name = newName
}

func (p *Person) changeName(newName string){
  p.Name = newName
}

func main() {
  var myPerson Person
  
  myPerson.Name = "Melkey"
  myPerson.Age = 4
  
  fmt.Println("Create struct without function")
  fmt.Println(myPerson)

  //formating stuct print with %+v
  fmt.Printf("This is myPerson print with formating %%+v \n%+v\n",myPerson)

  //Passed a name and age with newPerson function
  secondPerson := newPerson("Bob", 12)
  fmt.Println("Create struct with newPerson() function")
  fmt.Println(secondPerson)

  //change name with function
  //This wouldnt work it would still "bob"
  //Because of passing values as copies not references i.e need pointer
  wrongChangeName(secondPerson, "Julie") //different memory address to the first secondPerson line 40
  fmt.Println(secondPerson)

  //This function used Pointer receiver (p Person)
  //And it change the allocated memory on secondPerson struct
  secondPerson.changeName("Peter") //same memory address to the first secondPerson 
  fmt.Println(secondPerson)

}
