package safety

import (
	"fmt"
	"passgen/helpers"
)

// Check is tests password for bunch of requirements
func Check(password string) {
	length := helpers.ProperCharacterCounter(password)
	upperCount, lowerCount, numberCount, symbolCount, nonStandardCount := characterTypeCounter(password)

	fmt.Print("Is password at least 8 characters? : ")
	if length >= 8 {
		helpers.PositivePrintf("Yes! It's %v characters long!\n", length)
	} else {
		helpers.NegativePrintln("No! You may want to consider using longer password!")
	}

	fmt.Print("Is password really long enough? : ")
	if length >= 16 {
		helpers.PositivePrintln("Yes! It's great!")
	} else {
		helpers.NegativePrintln("No! But it's ok")
	}

	fmt.Print("Is password contains any upper-case characters? : ")
	if upperCount > 0 {
		helpers.PositivePrintln("Yes!")
	} else {
		helpers.NegativePrintln("No! You may want to consider using upper-case characters in your password!")
	}

	fmt.Print("Is password contains any lower-case characters? : ")
	if lowerCount > 0 {
		helpers.PositivePrintln("Yes!")
	} else {
		helpers.NegativePrintln("No! You may want to consider using lower-case characters in your password!")
	}

	fmt.Print("Is password contains any digits? : ")
	if numberCount > 0 {
		helpers.PositivePrintln("Yes!")
	} else {
		helpers.NegativePrintln("No! You may want to consider using digits in your password!")
	}

	fmt.Print("Is password contains any symbols? : ")
	if symbolCount > 0 {
		helpers.PositivePrintln("Yes!")
	} else {
		helpers.NegativePrintln("No! You may want to consider using symbols in your password!")
	}

	fmt.Print("Is password contains any non-standard characters? : ")
	if nonStandardCount > 0 {
		helpers.PositivePrintln("Yes! It's great!")
	} else {
		helpers.NegativePrintln("No! But it's ok")
	}

	fmt.Print("Is password contains any repeating characters? : ")
	if !isThereAnyRepeatingRunes(password) {
		helpers.PositivePrintln("No! So, it's great!")
	} else {
		helpers.NegativePrintln("Yes and it can be dangerous!")
	}

	fmt.Print("Is password contains any repeating character groups? : ")
	if !isThereAnyRepeatingRuneGroups(password) {
		helpers.PositivePrintln("No! It's good!")
	} else {
		helpers.NegativePrintln("Yes and you must avoid it!")
	}
}
