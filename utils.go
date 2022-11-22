package main

import (
	"reflect"
	"strings"
)

var found bool
func _walk(record map[string]interface{}, term string) {
	// fmt.Println("Term: ", term)
	for _, element := range record {
		// fmt.Println("Key: ", key)
		// fmt.Println("Element:", element)
		if element == nil {
			continue
		} else if reflect.TypeOf(element).Kind() == reflect.String {
			if strings.Contains(element.(string), term) {
				found = true
			}
		} else if reflect.TypeOf(element) == reflect.TypeOf(record) {
			_walk(element.(map[string]interface{}), term)
		}
    }
}

func searchData(jsonData []map[string]interface{}, searchKey string, delimiter string) []map[string]interface{} {
	var transformedData []map[string]interface{}
	for _, s := range jsonData {
		terms := s[searchKey].(string)
		delete(s,searchKey)
		res := strings.Split(terms, ",")
		for _, term := range res {
			_walk(s, strings.TrimSpace(term))
			if found  {
				transformedData = append(transformedData, s)
			}
			found = false
			
		}
	}
	return transformedData
}