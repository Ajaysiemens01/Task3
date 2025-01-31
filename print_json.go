package main

import (
	"fmt"
)

// Print person details
func(person Person) printPerson(){
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
}

//Print Persons Array details
func printPersonsArray(persons []Person) {
   for _, person := range persons {
	   fmt.Printf("Name: %s , Age: %d\n", person.Name, person.Age)
   }
}

//Print the nested json object with fields that can hold multiple types
func printJson(keyItem string, valueItems interface{}) {
	// Check the type of the valueItems and handle accordingly
	switch value := valueItems.(type) {
	case string:
		fmt.Printf("%s: %s\n", keyItem, value)
	case float64: // JSON numbers are decoded as float64 by default
		fmt.Printf("%s: %v\n", keyItem, value)
	case bool:
		fmt.Printf("%s: %v\n", keyItem, value)
	case []interface{}:
		// If the valueItems is an array, print its elements
		fmt.Printf("%s: [", keyItem)
		for index, item := range value {
			if index > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%v", item)
			// recursively print array elements
		}
		fmt.Print("]\n")
	case map[string]interface{}:
		// If the valueItems is an object, recursively process its fields	
		for key, subValue := range value {
			fmt.Printf("%s ", keyItem)
			printJson(key, subValue)
			fmt.Printf("\n")
		}
	default:
		// Handle unknown types (can add more cases if needed)
		fmt.Printf("%s: %v (unknown type)\n", keyItem, value)
	}
}