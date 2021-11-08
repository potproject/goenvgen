package gen

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_checker(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name  string
		text  string
		want  reflect.Kind
		want1 interface{}
		want2 bool
	}{
		{
			name:  "bool test1",
			text:  "TRUE",
			want:  reflect.Bool,
			want1: true,
			want2: false,
		},
		{
			name:  "bool test2",
			text:  "false",
			want:  reflect.Bool,
			want1: false,
			want2: false,
		},
		{
			name:  "bool test3",
			text:  "True",
			want:  reflect.Bool,
			want1: true,
			want2: false,
		},
		{
			name:  "int64 test1",
			text:  "0",
			want:  reflect.Int64,
			want1: int64(0),
			want2: false,
		},
		{
			name:  "int64 test2",
			text:  "-9223372036854775808",
			want:  reflect.Int64,
			want1: int64(-9223372036854775808),
			want2: false,
		},
		{
			name:  "int64 test3",
			text:  "9223372036854775807",
			want:  reflect.Int64,
			want1: int64(9223372036854775807),
			want2: false,
		},
		{
			name:  "float64 test1",
			text:  "9223372036854775808",
			want:  reflect.Float64,
			want1: float64(9223372036854775808),
			want2: false,
		},
		{
			name:  "float64 test2",
			text:  "0.1",
			want:  reflect.Float64,
			want1: float64(0.1),
			want2: false,
		},
		{
			name:  "float64 test3",
			text:  "-0.1234567890",
			want:  reflect.Float64,
			want1: float64(-0.1234567890),
			want2: false,
		},
		{
			name:  "string test1",
			text:  "",
			want:  reflect.String,
			want1: "",
			want2: false,
		},
		{
			name:  "string test2",
			text:  "-",
			want:  reflect.String,
			want1: "-",
			want2: false,
		},
		{
			name:  "string test3",
			text:  "100m",
			want:  reflect.String,
			want1: "100m",
			want2: false,
		},
		{
			name:  "[]string test",
			text:  ",",
			want:  reflect.String,
			want1: []string{"", ""},
			want2: true,
		},
		{
			name:  "[]int64 test",
			text:  "1,10,100,1000",
			want:  reflect.Int64,
			want1: []int64{1, 10, 100, 1000},
			want2: true,
		},
		{
			name:  "[]float64 test",
			text:  "0.1,0.01,0.001,0.0001",
			want:  reflect.Float64,
			want1: []float64{0.1, 0.01, 0.001, 0.0001},
			want2: true,
		},
		{
			name:  "[]bool test",
			text:  "true,false,false",
			want:  reflect.Bool,
			want1: []bool{true, false, false},
			want2: true,
		},
		{
			name:  "[]string test2",
			text:  "true,100,abc",
			want:  reflect.String,
			want1: []string{"true", "100", "abc"},
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := checker(tt.text)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("User value is mismatch (-got +tt.want):\n%s", diff)
			}
			if diff := cmp.Diff(got1, tt.want1); diff != "" {
				t.Errorf("User value is mismatch (-got1 +tt.want1):\n%s", diff)
			}
			if diff := cmp.Diff(got2, tt.want2); diff != "" {
				t.Errorf("User value is mismatch (-got1 +tt.want1):\n%s", diff)
			}
		})
	}
}