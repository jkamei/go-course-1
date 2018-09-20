package main
import ("fmt"
	"strconv"
	"os"
	"sort"
	"strings")



func main() {
	var sortedList = make([]int, 3)
	for {
		var inputString string
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&inputString)
		
		if strings.ToLower(inputString) == "x" {
			break
		}
		inputInt, err := strconv.Atoi(inputString)
		if err != nil {
	        // handle error
	        fmt.Println(err)
	        os.Exit(2)
	    }
		fmt.Println(inputInt)

		sortedList = append(sortedList, inputInt)
		fmt.Println(sortedList)

		sort.Ints(sortedList)
		fmt.Println(sortedList)
	}
}