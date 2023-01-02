package main

/**
Printing . Print Values to Console
doc: https://go.dev/doc/effective_go#printing
*/

/**
Package: fmt (format)
*/
import (
	"fmt"
	"os"
)

func main() {
	// Print
	fmt.Print("A")
	fmt.Print("B")
	fmt.Print("C")
	fmt.Println("D")
	fmt.Printf("Hello, my name is %s, i am %d years old!", "Kunal", 20)
	fmt.Println("Hey", "Kunal", "Singh")
	fmt.Println("Hello", 23)

	// Fprint
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stderr, "Hello, %s\n", "Kunal")
	fmt.Fprintln(os.Stderr, "Hello")

	// Sprint
	x := fmt.Sprint("Kunal", "Sin9h")
	y := fmt.Sprintln("Kunal", "Sin9h")
	z := fmt.Sprintf("Username: %s", "KunalSin9h")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Print(z)
}
