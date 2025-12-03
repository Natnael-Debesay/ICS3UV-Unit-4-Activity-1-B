/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-12-02
 * @fileoverview This porgram prints out all the letters of the alphabet from A to Z.
 */

// variables
let charactersASCIINumber: number = 0;

// start out printing the CAPITAL letters
console.log("The letters of the alphabet from A to Z are: ");

// loop to print CAPITAL letters from A to Z
for (
  let characterASCIINumber = 65;
  characterASCIINumber <= 90;
  characterASCIINumber++
) {
  console.log(String.fromCharCode(characterASCIINumber));
}

// now print out the small letters from a to z
for (
  let characterASCIINumber = 97;
  characterASCIINumber <= 122;
  characterASCIINumber++
) {
  console.log(String.fromCharCode(characterASCIINumber));
}

console.log("\nDone.");
