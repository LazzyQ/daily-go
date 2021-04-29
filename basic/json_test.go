package basic

import (
	"encoding/json"
	"testing"
)

type StructValue struct {
	IntValue    int    `json:"int_value,omitempty"`
	StringValue string `json:"string_value,omitempty"`
}

type JsonOmitempty struct {
	IntValue      int         `json:"int_value,omitempty"`
	BoolValue     bool        `json:"bool_value,omitempty"`
	StringValue   string      `json:"string_value,omitempty"`
	IntSliceValue []int       `json:"int_slice_value,omitempty"`
	StructValue   StructValue `json:"struct_value,omitempty"`
}

func TestJsonOmitempty(t *testing.T) {
	jsonOmitempty := JsonOmitempty{}
	bytes, _ := json.Marshal(jsonOmitempty)
	t.Logf("序列化后的结果: %s", string(bytes))

	var jsonOmitemptyUnmarshal JsonOmitempty
	_ = json.Unmarshal(bytes, &jsonOmitemptyUnmarshal)
	t.Logf("反序列化后的结果: %+v", jsonOmitemptyUnmarshal)
}
