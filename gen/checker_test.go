package gen

import (
	"testing"

	"github.com/potproject/goenvgen/model"

	"github.com/google/go-cmp/cmp"
)

func Test_checker(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name                     string
		text                     string
		KindWithSliceAndRequired model.KindWithSliceAndRequired
		value                    interface{}
	}{
		{
			name:                     "bool test1",
			text:                     "TRUE",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Bool, Slice: false, Required: false},
			value:                    true,
		},
		{
			name:                     "bool test2",
			text:                     "false",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Bool, Slice: false, Required: false},
			value:                    false,
		},
		{
			name:                     "bool test3",
			text:                     "True",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Bool, Slice: false, Required: false},
			value:                    true,
		},
		{
			name:                     "int64 test1",
			text:                     "0",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Int64, Slice: false, Required: false},
			value:                    int64(0),
		},
		{
			name:                     "int64 test2",
			text:                     "-9223372036854775808",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Int64, Slice: false, Required: false},
			value:                    int64(-9223372036854775808),
		},
		{
			name:                     "int64 test3",
			text:                     "9223372036854775807",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Int64, Slice: false, Required: false},
			value:                    int64(9223372036854775807),
		},
		{
			name:                     "float64 test1",
			text:                     "9223372036854775808",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Float64, Slice: false, Required: false},
			value:                    float64(9223372036854775808),
		},
		{
			name:                     "float64 test2",
			text:                     "0.1",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Float64, Slice: false, Required: false},
			value:                    float64(0.1),
		},
		{
			name:                     "float64 test3",
			text:                     "-0.1234567890",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Float64, Slice: false, Required: false},
			value:                    float64(-0.1234567890),
		},
		{
			name:                     "string test1",
			text:                     "",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.String, Slice: false, Required: false},
			value:                    "",
		},
		{
			name:                     "string test2",
			text:                     "-",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.String, Slice: false, Required: false},
			value:                    "-",
		},
		{
			name:                     "string test3",
			text:                     "100m",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.String, Slice: false, Required: false},
			value:                    "100m",
		},
		{
			name:                     "[]string test",
			text:                     ",",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.String, Slice: true, Required: false},
			value:                    []string{"", ""},
		},
		{
			name:                     "[]int64 test",
			text:                     "1,10,100,1000",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Int64, Slice: true, Required: false},
			value:                    []int64{1, 10, 100, 1000},
		},
		{
			name:                     "[]float64 test",
			text:                     "0.1,0.01,0.001,0.0001",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Float64, Slice: true, Required: false},
			value:                    []float64{0.1, 0.01, 0.001, 0.0001},
		},
		{
			name:                     "[]bool test",
			text:                     "true,false,false",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.Bool, Slice: true, Required: false},
			value:                    []bool{true, false, false},
		},
		{
			name:                     "[]string test2",
			text:                     "true,100,abc",
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.String, Slice: true, Required: false},
			value:                    []string{"true", "100", "abc"},
		},
		{
			name:                     "JSON test1",
			text:                     `{}`,
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.JSON, Slice: false, Required: false},
			value:                    map[string]interface{}{},
		},
		{
			name:                     "JSON test2",
			text:                     `{"id": 1, "Name": "Test" , "_r": 1.234567890}`,
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.JSON, Slice: false, Required: false},
			value:                    map[string]interface{}{"Name": string("Test"), "_r": float64(1.23456789), "id": float64(1)},
		},
		{
			name:                     "JSON test3",
			text:                     `{"IDs": [1,2,"500","ABC"]}`,
			KindWithSliceAndRequired: model.KindWithSliceAndRequired{Kind: model.JSON, Slice: false, Required: false},
			value:                    map[string]interface{}{"IDs": []interface{}{float64(1), float64(2), string("500"), string("ABC")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KindWithSliceAndRequired, value := checker(tt.text)
			if diff := cmp.Diff(KindWithSliceAndRequired, tt.KindWithSliceAndRequired); diff != "" {
				t.Errorf("User value is mismatch (-got +tt.KindWithSliceAndRequired):\n%s", diff)
			}
			if diff := cmp.Diff(value, tt.value); diff != "" {
				t.Errorf("User value is mismatch (-got1 +tt.value):\n%s", diff)
			}
		})
	}
}
