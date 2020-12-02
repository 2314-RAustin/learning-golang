package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22;
	fmt.Println("mi edad es " + strconv.Itoa(edad));

	edaddos := "22";
	edad_int, _ := strconv.Atoi(edaddos);

	fmt.Println(edaddos, edad_int)
} 