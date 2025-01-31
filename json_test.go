package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test marshaling a Person struct
func TestMarshaling(t *testing.T) {
	person := Person{Name: "Ajay Kumar", Age: 21, Email: "ajaykumarsiemens@gmail.com"}
	expectedPerson := `{"name":"Ajay Kumar","age":21,"email":"ajaykumarsiemens@gmail.com"}`

	marshaledPerson, err := MarshalPerson(person)
	assert.NoError(t, err)
	assert.JSONEq(t, expectedPerson, string(marshaledPerson))
}


// Test unmarshaling a JSON byte slice into a Person struct
func TestUnmarshaling(t *testing.T) {
	jsonData := []byte(`{"name":"Ajay Kumar","age":21,"email":"ajaykumarsiemens@gmail.com"}`)
	expectedPerson := Person{Name: "Ajay Kumar", Age: 21, Email: "ajaykumarsiemens@gmail.com"}

	person, err := UnmarshalPerson(jsonData)
	assert.NoError(t, err)
	assert.Equal(t, expectedPerson, person)
}

// Test unmarshaling with an error (invalid data type)
func TestUnmarshalingWithError(t *testing.T) {
	jsonData := []byte(`{"name":"Jane Smith", "age":"25", "email":"jane.smith@example.com"}`)
	expectedError := "json: cannot unmarshal string into Go struct field .age of type int"
	expectedPerson := Person{}
	person, err := UnmarshalPerson(jsonData)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)
	assert.Equal(t,expectedPerson, person)
}

// Test unmarshaling an array of persons
func TestUnmarshalPersonsArray(t *testing.T) {
	personsArray := `[{"name": "Emily", "age": 25, "email": "emily@example.com"},
					  {"name": "Tom", "age": 30, "email": "tom@example.com"},
					  {"name": "Jessica", "age": 35, "email": "jessica@example.com"}]`

	expectedPersons := []Person{
		{Name: "Emily", Age: 25, Email: "emily@example.com"},
		{Name: "Tom", Age: 30, Email: "tom@example.com"},
		{Name: "Jessica", Age: 35, Email: "jessica@example.com"},
	}

	persons, err := UnmarshalPersonsArray(personsArray)
	assert.NoError(t, err)
	assert.Equal(t, expectedPersons, persons)
}

// Test for JSON to byte slice conversion
func TestConvertJsonToBytes(t *testing.T) {
	personJsonString := `{"name": "David", "age": 40, "email": "david@example.com"}`
	expectedByteSlice := []byte(personJsonString)

	byteSlice := JsonToByteSlice(personJsonString)
	assert.Equal(t, expectedByteSlice, byteSlice)
}

// Test for ByteSlice to JSON string conversion
func TestConvertBytesToJson(t *testing.T) {
	personByteSlice := []byte(`{"name": "David", "age": 40, "email": "david@example.com"}`)
	expectedJsonString := `{"name": "David", "age": 40, "email": "david@example.com"}`

	jsonString := ByteSliceToJson(personByteSlice)
	assert.Equal(t, expectedJsonString, jsonString)
}

// Test handling of JSON with missing fields
func TestJsonWithMissingKey(t *testing.T) {
	person := []byte(`{ "name": "Samuel", "age": 22 }`)
	expectedPerson := Person{Name: "Samuel", Age: 22, Email: "default@example.com"}

	actualPerson, err := UnmarshalPerson(person)
	assert.NoError(t, err)
	assert.Equal(t, expectedPerson, actualPerson)
}


// Test to find Value for a key in nested json
func TestFindingKeyInNestedJson(t *testing.T) {
	nestedJson := `{"id": 101, "profile": {"name": "Alice Johnson", "details": {"age": 28, "email": "alice.johnson@example.com"}}}`

	// Unmarshal JSON to map
	nestedJsonData, err := UnmarshalJsonToMap(nestedJson)
	assert.NoError(t, err)

	// Test for the "email" key
	expectedEmail := "alice.johnson@example.com"
	email, err := FindKeyInJson(nestedJsonData, "email")
	assert.NoError(t, err)
	assert.Equal(t, expectedEmail, email)
}

