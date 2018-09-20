package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main() {

	for{
		in := bufio.NewReader(os.Stdin)
		fmt.Printf("enter a string: ")
		userInput, _ := in.ReadString('\n')
		loweredInput := strings.ToLower(strings.TrimSpace(userInput))

		if strings.Index(loweredInput, "i") == 0 && strings.Index(loweredInput, "a") != -1 && strings.LastIndex(loweredInput, "n") == len(loweredInput)-1 {
			} else {
				fmt.Println("Not Found!")
			}
	}	
}