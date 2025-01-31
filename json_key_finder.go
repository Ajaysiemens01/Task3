// Reads a JSON string (it can contain nested JSON objects and arrays).
// Iterates through the structure to find and print the value of a given key (e.g., "email").
// If the key is not found, print a message stating that the key doesn't exist.

package main

import(
	"fmt"
)
// FindKey recursively searches for a key in a nested JSON structure.
func FindKeyInJson(jsonData map[string]interface{}, requiredKey string) (interface{}, error) {
	for key, value := range jsonData {
		// Check if the current key matches the required key.
		if key == requiredKey {
			return value, nil
		}
		// If the value is a nested map, recursively search within it.
		if nestedMap, ok := value.(map[string]interface{}); ok {
			requiredValue, err := FindKeyInJson(nestedMap, requiredKey)
		    if err == nil{
				return requiredValue, nil
			} 
			
		}
	}
	// Return an error if the key is not found.
	return nil, fmt.Errorf("key '%s' not found", requiredKey)
}
