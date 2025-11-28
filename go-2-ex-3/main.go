package main

import "fmt"

func main() {
	var modules = make(map[int]string)

	modules[104] = "Datenmodell implementieren"
	modules[117] = "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren"
	modules[346] = "Cloud Lösungen konzipieren und realisieren"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)

	modules[999] = "Module 999"

	modules[346] = "Cloud Lösungen konzipieren und realisieren V.2"

	fmt.Println(modules)
}
