package datatypes

import (
	"database/sql/driver"
	"strings"
)

type StringArray []string

func (s *StringArray) Scan(value interface{}) (err error) {
	*s = strings.Split(value.(string), ",")
	return
}

func (s *StringArray) Value() (driver.Value, error) {
	return strings.Join(*s, ","), nil
}
