package main

import(
	"encoding/json"
	"log"
)
// Person struct with JSON tags
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

//Add default values to the object of Person Struct
func (person *Person) UnmarshalJSON(data []byte) error {
	type Alias Person
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(person),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if person.Name == "" {
		person.Name = "Default"
	}
	if person.Age == 0 {
		person.Age = 18
	}
	if person.Email == "" {
		person.Email = "default@example.com"
	}

	return nil
}

// Print nestedjson using printJson function in print_nested_json.go
func main() {
	// Sample JSON input (can be replaced with any valid JSON object)
	jsonData := `{
		"id": 123,
		"tags": ["golang", "json", "programming"],
		"profile": {
			"name": "John",
			"age": 30
		},
		"active": true
	}`

	// Declare a generic map to hold any JSON structure
	var jsonObject map[string]interface{}

	// Unmarshal JSON into the map
	jsonObject,err := UnmarshalJsonToMap(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// Process and print the JSON data
	for keyItem, valueItems := range jsonObject {
		printJson(keyItem, valueItems)
	}
}
