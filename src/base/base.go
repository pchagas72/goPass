package base

import (
	"math/rand"
	"strings"
	"time"
)

var lettersArray []rune = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var numbersArray []rune = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var specialCharactersArray []rune = []rune{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.'}
var all_arrays [][]rune = [][]rune{lettersArray, numbersArray, specialCharactersArray}

func Check(err error) {
	if err != nil {
		println("An error ocurred.")
		println(err)
	}
}

// This function fills a master key with predictable characters based on the key length, not so secure. But it works.
// In the future I may change this to the number of lowercase letters, numbers and symbols.
func AddKeyChars(key string) string {
	var leng int = len(key)
	var llen int = 32 - leng
	var fill_element_key = llen
	var counter int = 0
	var i int = 0
	var new_key string = key
	for i < llen {
		if counter == 0 {
			if llen-i < len(lettersArray) {
				new_key += string(lettersArray[fill_element_key-i])
			} else {
				new_key += string(lettersArray[0])
			}
			counter++
		} else if counter == 1 {
			if llen-i < len(numbersArray) {
				new_key += string(numbersArray[fill_element_key-i])
			} else {
				new_key += string(numbersArray[0])
			}
			counter++
		} else if counter == 2 {
			if llen-i < len(specialCharactersArray) {
				new_key += string(specialCharactersArray[fill_element_key-i])
			} else {
				new_key += string(specialCharactersArray[0])
			}
			counter = 0
		}
		i++
	}
	return new_key
}

func getRandomCharacter() rune {
	var random_array_index int = rand.Intn(len(all_arrays))
	var random_array []rune = all_arrays[random_array_index]
	var random_rune_index int = rand.Intn(len(random_array))
	return random_array[random_rune_index]
}

func addToKey(password string) string {
	var leng int = len(password)
	var llen int = 32 - leng
	for i := 0; i < llen; i++ {
		password = AddKeyChars(password)
	}
	return password
}

func CreateKey(key string) string {
	var leng int = len(key)
	if leng < 32 {
		key = addToKey(key) // Change this function for something more secure.
	}
	return key
}

// This function creates an random (and somewhat secure) password.
func GeneratePassword(leng int, minSpecialChars int, minMumbers int, minUppercase int) string {
	if leng < 16 {
		println("The minimun length for an password is 16 characters.")
	}
	var password strings.Builder
	var counter int = 0
	rand.Seed(time.Now().UnixNano())
	for counter <= minSpecialChars {
		var random_index int = rand.Intn(len(specialCharactersArray))
		var special_char string = string(specialCharactersArray[random_index])
		password.WriteString(special_char)
		counter++
	}
	for counter <= minMumbers {
		var random_index int = rand.Intn(len(numbersArray))
		var number string = string(numbersArray[random_index])
		password.WriteString(number)
		counter++
	}
	for counter <= minUppercase {
		var random_index int = rand.Intn(len(lettersArray))
		var random_uppercase string = strings.ToUpper(string(lettersArray[random_index]))
		password.WriteString(random_uppercase)
		counter++
	}
	for counter <= leng {
		var character rune = getRandomCharacter()
		password.WriteString(string(character))
		counter++
	}
	return password.String()
}
