package basic

import (
	"regexp"
	"testing"
)

var (
	reLink = `src="(https?://[\s\S]+?)"`
)

func TestRegrex(t *testing.T) {
	str := `<p>软的手机壳，炒鸡好用</p>\n\n<p><img align="absmiddle" src="https://img.alicdn.com/imgextra/i1/624425305/O1CN015H5enE1p3jnw0Oe8K_!!624425305.gif" style="max-width:750.0px;"><img align="absmiddle" src="https://img.alicdn.com/imgextra/i4/624425305/O1CN0121h4Ry1p3jnkliQEE_!!624425305.gif" style="max-width:750.0px;"><img align="absmiddle" src="https://img.alicdn.com/imgextra/i1/624425305/O1CN013Cdbxv1p3jn3VFpTT_!!624425305.jpg" style="max-width:750.0px;"><img align="absmiddle" src="https://img.alicdn.com/imgextra/i1/624425305/O1CN010Ma4dz1p3jnAY5FT8_!!624425305.jpg" style="max-width:750.0px;"><img align="absmiddle" src="https://img.alicdn.com/imgextra/i3/624425305/O1CN01O5wuQZ1p3jmDdY4oV_!!624425305.jpg" style="max-width:750.0px;"><img align="absmiddle" src="https://img.alicdn.com/imgextra/i3/624425305/O1CN01wfuZBa1p3jm7WdQFe_!!624425305.jpg" style="max-width:750.0px;"></p>`

	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(str, -1)
	for _, result := range results {
		t.Logf("%v", result[1])
	}
}
