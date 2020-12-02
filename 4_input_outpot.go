package main

import (
	"fmt"
	"bufio"
	"os"
);

func main() {
	// fmt.Print("Hola Mundo");
	edad := 22;
	boolean := false;
	float := 14.32;

	fmt.Printf("Imprimir un string despues un entero %d\n ", edad);
	fmt.Printf("Imprimir un string despues un booleano %t\n ", boolean);
	fmt.Printf("Imprimir un string despues un flotante %f\n ", float);
	fmt.Printf("Imprimir un string despues un flotante %e\n ", float);
	fmt.Printf("Imprimir un string despues un flotante %b\n ", float);
	fmt.Printf("Imprimir un string concatenados a un string %s ", "Roberto");

	var edados int;
	fmt.Println("Coloca la edad: ");
	fmt.Scanf("%d\n",&edados);
	fmt.Printf("Mi edad es %d\n", edados);

	var nombre string;
	fmt.Println("Coloca la nombre: ");
	fmt.Scanf("%v\n",&nombre);
	fmt.Printf("Mi nombre es %s\n", nombre);
}

func readNameWithosStdin()  {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Ingresa tu nombre: ");
	nombre,err := reader.ReadString('\n');
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}
}