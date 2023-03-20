package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	// Defining t for MarshalJSON method
	t := time.Date(2013, 7, 6, 12, 34, 33, 01, time.UTC)

	// Calling MarshalJSON() method
	encoding, _ := t.MarshalJSON()

	fmt.Println("time marchal is ", string(encoding))

	jsonText, _ := json.Marshal(t)
	fmt.Println("json marchal is ", string(jsonText))
	// Defining tm for UnmarshalJSON() method
	var tm time.Time

	// Calling UnmarshalJSON method with its parameters
	decode := tm.UnmarshalJSON(encoding)

	// Prints output
	fmt.Printf("Error: %v\n", decode)
	var mt time.Time
	if err := json.Unmarshal(([]byte(jsonText)), &mt); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json unmarchalr: %v\n", mt)

}
