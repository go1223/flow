package main

import "fmt"

func main() {
	fmt.Println("single flow client")
	for i := 0; i < 1; i++ {
		go func(index int) {
			fmt.Printf("client %d\n", index)
			var key = fmt.Sprintf("%d", index)
			err := GetMessageFromServer("124.70.197.102", "8890", key)
			if err != nil {
				fmt.Println(key, err)
			}
		}(i + 1)
	}
	select {}
}
