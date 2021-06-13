package main

import (
	"flag"
	"fmt"
	"strconv"
)

type roll struct{
	batchYear int
	uniqueNumber int
	branchCode string
}

func (r *roll) Set(s string) error{
	r.batchYear, _ = strconv.Atoi(s[:4])
	r.uniqueNumber, _ = strconv.Atoi(s[4:7])
	r.branchCode = s[7:]
	return nil
}

func (r *roll) String() string{
	return fmt.Sprint(r.batchYear) + fmt.Sprint(r.uniqueNumber) + r.branchCode 
}

func main(){
	var regn roll
	flag.Var(&regn, "r", "registration numberer")
	age := flag.Int("a", 0, "Age of Person")
	name := flag.String("n", "NA", "Name of person" )
	student := flag.Bool("s", false, "If person is a student")

	flag.Parse()

	fmt.Printf("Age: %d \n", *age)
	fmt.Printf("Name: %s \n", *name)
	fmt.Printf("Student: %v \n", *student)
	fmt.Printf("Registration: %v \n", &regn)
	fmt.Printf("Tail: %v \n", flag.Args())
}