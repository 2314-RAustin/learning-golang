package main

import (
	"fmt"
)

func main() {
	//iteracion clasica
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//iteracion opcional
	i := 0
	for {
		fmt.Println(i)
		if i == 2 {
			i++
			continue
		}
		if i == 10 {
			break
		}
		i++
	}

	// for _, var := 10 var {
	// }

}
