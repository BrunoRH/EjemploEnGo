package main

import (
	"fmt"
)

func LongitudCadena(){
	
	var cadena string
	fmt.Println("Ingrese una cadena : ")
	fmt.Scanf("%s", &cadena)
	
	fmt.Printf("La cantidad de caracteres de la cadena es %d",len(cadena))

}

func main() {

	LongitudCadena()
	
	

}
