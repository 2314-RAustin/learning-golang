package main

import (
	"fmt"
);

func main()  {
	x := 10;
	y := 5;

	if x >= y {
		fmt.Println("%d es mayot que %d \n", x,y);
	} else if y > x {
		fmt.Println("%d es mayot que %d \n", y,x);
	} else {
		fmt.Println("Son iguales")
	}
}