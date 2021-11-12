package envgentest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joho/godotenv"
)

func Test_EnvGenTestGet(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "..", "test", ".env.test")
	godotenv.Load(testfile)
	err := Load()
	if err != nil {
		t.Error(err.Error())
	}
	if true != Get().BOOL_TEST1() {
		t.Errorf("BOOL_TEST1 invalid")
	}
	if false != Get().BOOL_TEST2() {
		t.Errorf("BOOL_TEST2 invalid")
	}
	if true != Get().BOOL_TEST3() {
		t.Errorf("BOOL_TEST3 invalid")
	}
	if 0 != Get().INT64_TEST1() {
		t.Errorf("INT64_TEST1 invalid")
	}
	if -9223372036854775808 != Get().INT64_TEST2() {
		t.Errorf("INT64_TEST2 invalid")
	}
	if 9223372036854775807 != Get().INT64_TEST3() {
		t.Errorf("INT64_TEST3 invalid")
	}
	if 9223372036854775808 != Get().FLOAT64_TEST1() {
		t.Errorf("FLOAT64_TEST1 invalid")
	}
	if 0.1 != Get().FLOAT64_TEST2() {
		t.Errorf("FLOAT64_TEST2 invalid")
	}
	if -0.1234567890 != Get().FLOAT64_TEST3() {
		t.Errorf("FLOAT64_TEST3 invalid")
	}
	if "" != Get().STRING_TEST1() {
		t.Errorf("STRING_TEST1 invalid")
	}
	if "-" != Get().STRING_TEST2() {
		t.Errorf("FLOAT64_TEST2 invalid")
	}
	if "100m" != Get().STRING_TEST3() {
		t.Errorf("STRING_TEST3 invalid")
	}
	if "lowercasetest" != Get().Lowercasetest() {
		t.Errorf("lowercasetest invalid")
	}
	if "UPPERCASETEST" != Get().UPPERCASETEST() {
		t.Errorf("UPPERCASETEST invalid")
	}
	if "CamelCaseTest" != Get().CamelCaseTest() {
		t.Errorf("CamelCaseTest invalid")
	}
	if "lower_snake_case_test" != Get().Lower_snake_case_test() {
		t.Errorf("lower_snake_case_test invalid")
	}
	if "{\"\"}" != Get().INVALID_JSON_TEST() {
		t.Errorf("INVALID_JSON_TEST invalid")
	}
	if "underscore_test" != Get().UNDERSCORE_TEST() {
		t.Errorf("_TEST invalid")
	}
	if diff := cmp.Diff([]string{"true", "100", "abc"}, Get().S_STRING_TEST2()); diff != "" {
		t.Errorf("S_STRING_TEST2 invalid: %s", diff)
	}
	if diff := cmp.Diff([]int64{1, 10, 100, 1000}, Get().S_INT64_TEST()); diff != "" {
		t.Errorf("S_INT64_TEST invalid: %s", diff)
	}
	if diff := cmp.Diff([]float64{0.1, 0.01, 0.001, 0.0001}, Get().S_FLOAT64_TEST()); diff != "" {
		t.Errorf("S_FLOAT64_TEST invalid: %s", diff)
	}
	if diff := cmp.Diff([]bool{true, false, false}, Get().S_BOOL_TEST()); diff != "" {
		t.Errorf("S_BOOL_TEST invalid: %s", diff)
	}
	if diff := cmp.Diff(JSON_TEST{}, Get().JSON_TEST()); diff != "" {
		t.Errorf("JSON_TEST invalid: %s", diff)
	}
	if diff := cmp.Diff(JSON_TEST2{Name: "Test", R: 1.23456789, ID: 1}, Get().JSON_TEST2()); diff != "" {
		t.Errorf("JSON_TEST2 invalid: %s", diff)
	}
	if diff := cmp.Diff(JSON_TEST3{
		Category: struct {
			ID       int64  "json:\"ID\""
			Name     string "json:\"Name\""
			Products []struct {
				Name string "json:\"Name\""
				ID   string "json:\"_id\""
			} "json:\"Products\""
			Tags []float64 "json:\"Tags\""
		}{
			ID:   10,
			Name: "Test",
			Products: []struct {
				Name string "json:\"Name\""
				ID   string "json:\"_id\""
			}{{Name: "Test100", ID: "XV"}},
			Tags: []float64{0.1, 0.01, 0.001},
		},
		R:  []float64{1.23456789},
		ID: 1,
	}, Get().JSON_TEST3()); diff != "" {
		t.Errorf("JSON_TEST3 invalid: %s", diff)
	}
	if diff := cmp.Diff(Lower_snake_case_test_json{Ids: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, Get().Lower_snake_case_test_json()); diff != "" {
		t.Errorf("lower_snake_case_test_json invalid: %s", diff)
	}
}

func Test_envgentestGetForce(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "..", "test", ".env.test")
	godotenv.Load(testfile)
	err := Load()
	if err != nil {
		t.Error(err.Error())
	}
	if true != Get().FORCE_BOOL() {
		t.Errorf("FORCE_BOOL invalid")
	}
	if 1 != Get().FORCE_INT() {
		t.Errorf("FORCE_INT invalid")
	}
	if 127 != Get().FORCE_INT8() {
		t.Errorf("FORCE_INT8 invalid")
	}
	if 32767 != Get().FORCE_INT16() {
		t.Errorf("FORCE_INT16 invalid")
	}
	if 2147483647 != Get().FORCE_INT32() {
		t.Errorf("FORCE_INT32 invalid")
	}
	if 0 != Get().FORCE_INT64() {
		t.Errorf("FORCE_INT64 invalid")
	}
	if 255 != Get().FORCE_UINT8() {
		t.Errorf("FORCE_UINT8 invalid")
	}
	if 65535 != Get().FORCE_UINT16() {
		t.Errorf("FORCE_UINT16 invalid")
	}
	if 4294967295 != Get().FORCE_UINT32() {
		t.Errorf("FORCE_UINT32 invalid")
	}
	if 0 != Get().FORCE_UINT64() {
		t.Errorf("FORCE_UINT64 invalid")
	}
	if 0.1 != Get().FORCE_FLOAT32() {
		t.Errorf("FORCE_FLOAT32 invalid")
	}
	if 0.01 != Get().FORCE_FLOAT64() {
		t.Errorf("FORCE_FLOAT64 invalid")
	}
	if "interface" != Get().FORCE_INTERFACE() {
		t.Errorf("FORCE_INTERFACE invalid")
	}
	if diff := cmp.Diff([]bool{true, false}, Get().FORCE_S_BOOL()); diff != "" {
		t.Errorf("FORCE_S_BOOL invalid: %s", diff)
	}
	if diff := cmp.Diff([]int{100, 200, 300}, Get().FORCE_S_INT()); diff != "" {
		t.Errorf("FORCE_S_INT invalid: %s", diff)
	}
	if diff := cmp.Diff([]int8{16, 32, 64}, Get().FORCE_S_INT8()); diff != "" {
		t.Errorf("FORCE_S_INT8 invalid: %s", diff)
	}
	if diff := cmp.Diff([]int16{200, 400, 800}, Get().FORCE_S_INT16()); diff != "" {
		t.Errorf("FORCE_S_INT16 invalid: %s", diff)
	}
	if diff := cmp.Diff([]int32{300, 900, 2700}, Get().FORCE_S_INT32()); diff != "" {
		t.Errorf("FORCE_S_INT32 invalid: %s", diff)
	}
	if diff := cmp.Diff([]int64{400, 1600, 6400}, Get().FORCE_S_INT64()); diff != "" {
		t.Errorf("FORCE_S_INT64 invalid: %s", diff)
	}
	if diff := cmp.Diff([]uint8{2, 4, 8}, Get().FORCE_S_UINT8()); diff != "" {
		t.Errorf("FORCE_S_UINT8 invalid: %s", diff)
	}
	if diff := cmp.Diff([]uint16{3, 9, 27}, Get().FORCE_S_UINT16()); diff != "" {
		t.Errorf("FORCE_S_UINT16 invalid: %s", diff)
	}
	if diff := cmp.Diff([]uint32{4, 16, 64}, Get().FORCE_S_UINT32()); diff != "" {
		t.Errorf("FORCE_S_UINT32 invalid: %s", diff)
	}
	if diff := cmp.Diff([]uint64{5, 25, 125}, Get().FORCE_S_UINT64()); diff != "" {
		t.Errorf("FORCE_S_UINT64 invalid: %s", diff)
	}
	if diff := cmp.Diff([]float32{0.1, 0.01, 0.001}, Get().FORCE_S_FLOAT32()); diff != "" {
		t.Errorf("FORCE_S_FLOAT32 invalid: %s", diff)
	}
	if diff := cmp.Diff([]float64{0.01, 0.001, 0.0001}, Get().FORCE_S_FLOAT64()); diff != "" {
		t.Errorf("FORCE_S_FLOAT64 invalid: %s", diff)
	}
	if diff := cmp.Diff([]interface{}{"inter", "face"}, Get().FORCE_S_INTERFACE()); diff != "" {
		t.Errorf("FORCE_S_INTERFACE invalid: %s", diff)
	}
}
