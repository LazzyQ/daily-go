package basic

import "time"

type Student struct {
	Name     string
	Age      int
	Birthday time.Time `json:"birthday"`
}
