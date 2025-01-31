package main

import (
	"fmt"
	"encoding/json"
)

// marshalPerson takes a Person struct and converts it into a JSON byte slice.
func MarshalPerson(person Person) ([]byte, error) {
	marshaledPerson, err := json.Marshal(person)
	if err != nil {
		return nil, fmt.Errorf("error marshalling person: %v", err)
	}
	return marshaledPerson, nil
}

// unmarshalPerson takes a JSON byte slice and converts it back into a Person struct.
func UnmarshalPerson(PersonInBytes []byte) (Person, error) {
	var person Person
	if err := json.Unmarshal(PersonInBytes, &person); err != nil {
		return Person{}, fmt.Errorf("error unmarshalling person: %v", err)
	}
	person.printPerson()

	return person, nil
}

// unmarshalPersonsArray takes a JSON string representing an array of Persons and converts it into a slice of Person structs.
func UnmarshalPersonsArray(personsArray string) ([]Person, error) {
	var persons []Person
	if err := json.Unmarshal([]byte(personsArray), &persons); err != nil {
		return nil, fmt.Errorf("error unmarshalling persons array: %v", err)
	}
	// Iterate over the persons slice and print each person's name and age.
	printPersonsArray(persons)
	return persons, nil
}

//unmarshalJsonToMap takes a JSON string and converts to a map 
func UnmarshalJsonToMap(jsonStr string) (map[string]interface{},error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		return nil, fmt.Errorf("error unmarshalling persons array: %v", err)
	}
	return jsonData,nil
}
