package main

import ("fmt"
"os"
"bufio"
"encoding/json"
"strings")

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("enter a name: ")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("enter an address: ")
	address, _ := in.ReadString('\n')
	address = strings.TrimSpace(address)

	fmt.Println(name)
	fmt.Println(address)

	myMap := make(map[string]string)
	myMap["name"] = name
	myMap["address"] = address

	fmt.Println(myMap)

	barr, _ := json.Marshal(myMap)
	fmt.Printf("%s", barr)	
}