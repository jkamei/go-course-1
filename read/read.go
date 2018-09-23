package main

import ("fmt"
"strings"
"io/ioutil")

type Name struct {
	fname string
	lname string
}


func main() {
	dat, _ := ioutil.ReadFile("names.txt")

	fullNames := strings.Split(string(dat), "\n")

	nameSlice  := []Name{}

	for _, line := range fullNames {
		if line != "" {
			splitNames := strings.Split(line, " ")
			firstName, lastName := splitNames[0], splitNames[1]
			nameSlice = append(nameSlice, Name{fname: firstName, lname: lastName})
		}
	}

	for _, name := range nameSlice {
		fmt.Printf("first: %s, last: %s\n", name.fname, name.lname)
	}
}