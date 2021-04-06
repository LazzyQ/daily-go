package basic

import (
	"strings"
	"testing"
)

func TestStringsReplaceEmpty(t *testing.T) {
	replace := ""
	t.Log(strings.Replace("交流社区", replace, "<mark></mark>", -1))
}
