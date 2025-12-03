// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-12-02
// Fileoverview This program prints out the letters of the alphabet from A to Z.

package main

import "fmt"

func main() {
// variable declaration
var characterASCIINumber int
var characterASCIIString string

// start out printing the CAPITAL letters
fmt.Println("The letters of the alaphbet from A to Z are:")

// loop to print CAPITAL letters from A to Z
for characterASCIINumber = 65; characterASCIINumber <= 90; characterASCIINumber++ {
characterASCIIString = string(rune(characterASCIINumber))
fmt.Printf("%s\n", characterASCIIString)
}

// new print for the small letters
fmt.Println("\nThe letters of the alphabet from a to z are:")

// loop to print small letters from a to z
for characterASCIINumber = 97; characterASCIINumber <= 122; characterASCIINumber++ { 
characterASCIIString = string(rune(characterASCIINumber))
fmt.Printf("%s\n", characterASCIIString)
}

fmt.Println("\nDone.")
}