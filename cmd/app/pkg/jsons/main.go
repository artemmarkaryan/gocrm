package jsons

import "encoding/json"

func MarshallToString(v interface{}) (string, error) {
	r, err := json.Marshal(v)
	return string(r), err
}
