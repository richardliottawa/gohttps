package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const expiryDateLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(expiryDateLayout, s)
	return
}

func ParseData(message string) MyType {
	var mt MyType
	// the commented line procude error because of Z - UTC
	// in golang, z can only be on its own, the last char
	if err := json.Unmarshal(([]byte(jsonText)), &mt); err != nil {
		log.Fatal(err)
	}
	return mt
}

// type MyType struct {
// 	Name     string    `json:"name"`
// 	Expiring time.Time `json:"expiring"`
// }

// var jsonText = `
// {
//   "name": "foobar",
//   "expiring": "2020-11-20"
// }`

// func main() {

// 	fmt.Printf("%+v\n", ParseData(jsonText))
// }

type MyType struct {
	Name     string     `json:"name"`
	Expiring CustomTime `json:"expiring"`
}

var jsonText = []byte(`
  {
	"name": "foobar",
	"expiring": "2020-11-30"
  }`)

func main() {
	var mt MyType
	fmt.Println("before", string(jsonText))
	if err := json.Unmarshal(jsonText, &mt); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", mt)
}
