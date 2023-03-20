package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var allZones = map[string]string{
	"America/New_York": "US",
	"Asia/Kolkata":     "IN",
}

func isZoneExists(zone string) bool {
	_, ok := allZones[zone]
	return ok
}

// HandleTimeRequest is a handler for local time request of given timezone
func HandleTimeRequest(writer http.ResponseWriter, request *http.Request) {
	invalid := false
	tzparams := request.URL.Query().Get("tz")
	unescapedQuery, err := url.QueryUnescape(tzparams)
	if err != nil {
		invalid = true
	}
	tzs := strings.Split(unescapedQuery, ",")
	times := make(map[string]string, len(tzs))
	if !invalid {
		for _, value := range tzs {
			if strings.TrimSpace(value) != "" {
				if isZoneExists(value) {
					loc, _ := time.LoadLocation(value)
					times[value] = time.Now().In(loc).String()
				} else {
					invalid = true
				}
			}

		}
	}
	writer.Header().Set("Content-Type", "application/json")
	if invalid {
		errorPayload, _ := json.Marshal("invalid message")
		http.NotFound(writer, request)
		writer.Write([]byte(errorPayload))
	} else {
		if len(times) > 0 {
			payload, _ := json.Marshal(times)
			writer.Write([]byte(payload))
		} else {
			utctime := make(map[string]string, 1)
			utcloc, _ := time.LoadLocation("UTC")
			utctime["UTC"] = time.Now().In(utcloc).String()
			utcpayload, _ := json.Marshal(utctime)
			writer.Write([]byte(utcpayload))
		}
	}

}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	http.HandleFunc("/api/time", HandleTimeRequest)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
