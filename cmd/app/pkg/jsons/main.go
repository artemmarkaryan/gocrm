package jsons

import "encoding/json"

func MarshalToString(v interface{}) (string, error) {
	r, err := json.Marshal(v)
	return string(r), err
}
