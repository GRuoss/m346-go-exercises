package main

import "fmt"

func main() {
	modules := make(map[int]uint, 0)
	
	modules[104] = 3 
	modules[117] = 4 
	modules[346] = 5

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104)
	fmt.Println("Nach Löschen von Modul 104:", modules)

	modules[215] = 6 
	fmt.Println("Nach Hinzufügen von Modul 215:", modules)

	modules[117] = 7 
	fmt.Println("Nach Ersetzen von Modul 117:", modules)
}
