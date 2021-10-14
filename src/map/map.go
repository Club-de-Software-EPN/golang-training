package main

import (
	"fmt" // format package
	"sort"
)

func getProvinceCodes()(provinceCodes map[int]string){
	// Create an empty map   ; 17: "Pichincha"
	provinceCodes = make(map[int]string)

	provinceCodes[1]="Azuay"
	provinceCodes[2]="Bolivar"
	provinceCodes[3]="Cañar"
	provinceCodes[4]="Carchi"
	provinceCodes[5]="Cotopaxi"
	provinceCodes[6]="Chimborazo"
	provinceCodes[7]="El Oro"
	provinceCodes[8]="Esmeraldas"
	provinceCodes[9]="Guayas"
	provinceCodes[10]="Imbabura"
	provinceCodes[11]="Loja"
	provinceCodes[12]="Los Ríos"
	provinceCodes[13]="Manabí"
	provinceCodes[14]="Morona Santiago"
	provinceCodes[15]="Napo"
	provinceCodes[16]="Pastaza"
	provinceCodes[17]="Pichincha"
	provinceCodes[18]="Tungurahua"
	provinceCodes[19]="Zamora Chinchipe"
	provinceCodes[20]="Galápagos"
	provinceCodes[21]="Sucumbios"	
	provinceCodes[22]="Orellana"
	provinceCodes[23]="Santo Domingo de los Tsachilas"
	provinceCodes[24]="Santa Elena"
	
	return
}

func iteratingCodesAndProvinces(provinceCodes map[int]string){
	for key, value := range provinceCodes {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	} 
}

func iteratingCodesAndProvincesInOrder(provinceCodes map[int]string){
	var keys []int

	for key,_ := range provinceCodes {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("Key: %d, Value: %s\n", key, provinceCodes[key])
	} 
}


func main() {
	// codesAndProvinces := getProvinceCodes()
	// fmt.Println(codesAndProvinces)
	//iteratingCodesAndProvinces(codesAndProvinces)
	// iteratingCodesAndProvincesInOrder(codesAndProvinces)

	// Create a map with keys and values
	teams := map[string][]string{
		"CyberSecurity": {"Alejandro", "Ana", "Juan", "Ginno"},
		"Web": {"Juanjo", "Loreley", "Lesly", "Frankz"},
		"AI": {"Anderson", "Javier", "Gabriel"},
		"Videogames": {"Andrés", "Edison", "Daliana"},
	}

	fmt.Println("--- Before using delete ---")
	fmt.Println("len", len(teams)) // number of keys --> 4
	fmt.Println(teams)

	delete(teams, "Videogames")
	delete(teams, "Web")

	fmt.Println("--- After using delete ---")
	fmt.Println("len", len(teams)) // number of keys --> 2
	fmt.Println(teams)
}