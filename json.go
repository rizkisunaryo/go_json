package go_json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

func PrintJson(header string, w http.ResponseWriter, v interface{}) {
	newHeader := header + "json.PrintJson: "
	jsStr, err := json.Marshal(&v)
	if err!=nil {
		log.Println(header, "Message: Error in marshalling json: ", err.Error())
		fmt.Fprintln(w, "{\"Message\":\"Error in marshalling json\"}")
	} else {
		log.Println(newHeader, string(jsStr))
		fmt.Fprintln(w, string(jsStr))
	}
}