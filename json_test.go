package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test marshaling a Person struct
func TestMarshalPerson(t *testing.T) {
	person := Person{Name: "Ajay Kumar", Age: 21, Email: "ajaykumarsiemens@gmail.com"}
	expectedPerson := `{"name":"Ajay Kumar","age":21,"email":"ajaykumarsiemens@gmail.com"}`

	marshaledPerson, err := MarshalPerson(person)
	assert.NoError(t, err)
	assert.JSONEq(t, expectedPerson, string(marshaledPerson))
}


// Test unmarshaling a JSON byte slice into a Person struct
func TestUnmarshalPerson(t *testing.T) {
	person := []byte(`{"name":"Ajay Kumar","age":21,"email":"ajaykumarsiemens@gmail.com"}`)
	expectedPerson := Person{Name: "Ajay Kumar", Age: 21, Email: "ajaykumarsiemens@gmail.com"}

	actualPerson, err := UnmarshalPerson(person)
	assert.NoError(t, err)
	assert.Equal(t, expectedPerson, actualPerson)
}

// Test unmarshaling with an error (invalid data type)
func TestUnmarshalWithError(t *testing.T) {
	person := []byte(`{"name":"Jane Smith", "age":"25", "email":"jane.smith@example.com"}`)
	expectedError := "json: cannot unmarshal string into Go struct field .age of type int"
	expectedPerson := EmptyPerson
	actualPerson, err := UnmarshalPerson(person)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)
	assert.Equal(t,expectedPerson, actualPerson)
}

// Test unmarshaling an array of persons
func TestUnmarshalPersonsArray(t *testing.T) {
	persons := `[{"name": "Emily", "age": 25, "email": "emily@example.com"},
					  {"name": "Tom", "age": 30, "email": "tom@example.com"},
					  {"name": "Jessica", "age": 35, "email": "jessica@example.com"}]`

	expectedPersons := []Person{
		{Name: "Emily", Age: 25, Email: "emily@example.com"},
		{Name: "Tom", Age: 30, Email: "tom@example.com"},
		{Name: "Jessica", Age: 35, Email: "jessica@example.com"},
	}

	actualPersons, err := UnmarshalPersonsArray(persons)
	assert.NoError(t, err)
	assert.Equal(t, expectedPersons, actualPersons)
}

// Test for JSON to byte slice conversion
func TestConvertJsonToBytes(t *testing.T) {
	personJsonString := `{"name": "David", "age": 40, "email": "david@example.com"}`
	expectedByteSlice := []byte(personJsonString)

	actualByteSlice := JsonToByteSlice(personJsonString)
	assert.Equal(t, expectedByteSlice, actualByteSlice)
}

// Test for ByteSlice to JSON string conversion
func TestConvertBytesToJson(t *testing.T) {
	personByteSlice := []byte(`{"name": "David", "age": 40, "email": "david@example.com"}`)
	expectedJsonString := `{"name": "David", "age": 40, "email": "david@example.com"}`

	actualJsonString := ByteSliceToJson(personByteSlice)
	assert.Equal(t, expectedJsonString, actualJsonString)
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
func TestCheckKeyInNestedJson(t *testing.T) {
	nestedJson := `{"id": 101, "profile": {"name": "Alice Johnson", "details": {"age": 28, "email": "alice.johnson@example.com"}}}`

	// Unmarshal JSON to map
	nestedJsonObject, err := UnmarshalJsonToMap([]byte(nestedJson))
	assert.NoError(t, err)

	// Test for the "email" key
	expectedEmail := "alice.johnson@example.com"
	email, err := FindKeyInJson(nestedJsonObject, "email")
	assert.NoError(t, err)
	assert.Equal(t, expectedEmail, email)
}

