package main

import (
	"flag"
	"fmt"
)

func main(){
	age := flag.Int("a", 0, "Age of Person")
	name := flag.String("n", "NA", "Name of person" )
	student := flag.Bool("s", false, "If person is a student")

	flag.Parse()

	fmt.Printf("Age: %d \n", *age)
	fmt.Printf("Name: %s \n", *name)
	fmt.Printf("Student: %v \n", *student)
	fmt.Printf("Tail: %v \n", flag.Args())
}