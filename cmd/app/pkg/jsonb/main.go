package jsonb

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type JSONB map[string]interface{}

func (j JSONB) String() string {
	s, _ := jsons.MarshalToString(j)
	return s
}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	var payload []byte
	if v, ok := value.(string); ok {
		payload = []byte(v)
	}
	if err := json.Unmarshal(payload, &j); err != nil {
		return err
	}
	return nil
}
