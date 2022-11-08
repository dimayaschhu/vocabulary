package utils

import "encoding/json"

func Marshal(input interface{}) string {
	data, _ := json.MarshalIndent(input, "", "    ")

	return string(data)
}
