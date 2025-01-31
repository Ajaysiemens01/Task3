package main

//Convert JSON to ByteSlice
func JsonToByteSlice(inputJson string) []byte {
	return []byte(inputJson)
}

//Convert ByteSlice to JSON
func ByteSliceToJson(byteSlice []byte) string {
	return string(byteSlice)
}
