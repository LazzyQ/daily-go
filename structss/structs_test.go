package structss

import (
	"testing"

	"github.com/fatih/structs"
)

type Teacher struct {
	Name  string `json:"name,omitempty"`
	Class string `json:"class,omitempty"`
}

func TestStructsMap(t *testing.T) {
	teacher := Teacher{
		Name:  "马老师",
		Class: "",
	}
	t.Logf("%v", structs.Map(teacher))
}
