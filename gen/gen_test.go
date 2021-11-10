package gen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/potproject/goenvgen/model"
)

func Test_GenerateFile(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage, map[string]string{
		"FORCE_BOOL":        "bool",
		"FORCE_INT":         "int",
		"FORCE_INT8":        "int8",
		"FORCE_INT16":       "int16",
		"FORCE_INT32":       "int32",
		"FORCE_INT64":       "int64",
		"FORCE_UINT8":       "uint8",
		"FORCE_UINT16":      "uint16",
		"FORCE_UINT32":      "uint32",
		"FORCE_UINT64":      "uint64",
		"FORCE_FLOAT32":     "float32",
		"FORCE_FLOAT64":     "float64",
		"FORCE_INTERFACE":   "interface",
		"FORCE_S_BOOL":      "[]bool",
		"FORCE_S_INT":       "[]int",
		"FORCE_S_INT8":      "[]int8",
		"FORCE_S_INT16":     "[]int16",
		"FORCE_S_INT32":     "[]int32",
		"FORCE_S_INT64":     "[]int64",
		"FORCE_S_UINT8":     "[]uint8",
		"FORCE_S_UINT16":    "[]uint16",
		"FORCE_S_UINT32":    "[]uint32",
		"FORCE_S_UINT64":    "[]uint64",
		"FORCE_S_FLOAT32":   "[]float32",
		"FORCE_S_FLOAT64":   "[]float64",
		"FORCE_S_INTERFACE": "[]interface",
	})
	if err != nil {
		t.Error(err.Error())
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateFile_NotExist(t *testing.T) {
	testpackage := "envgen"
	err := GenerateFile("", testpackage, map[string]string{})
	if err == nil || err.Error() != "open .env: The system cannot find the file specified." {
		t.Error("Error: Not Exist Pattern Failed")
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateFile_Empty(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".empty.env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage, map[string]string{})
	if err == nil || err.Error() != "Dotenv file is empty." {
		t.Error("Error: Not Empty Pattern Failed")
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateFile_Duplicate(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".duplicate.env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage, map[string]string{})
	if err == nil || err.Error() != "duplicate Env fields. duplicate_string_test:Duplicate_string_test" {
		t.Error("duplicate validate Failed")
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateFile_ForceError(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage, map[string]string{
		"FORCE_BOOL": "json",
	})
	if err == nil || err.Error() != "Unsupported Type: FORCE_BOOL=json" {
		t.Error("Force Error Failed")
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_sortedKeys(t *testing.T) {
	m := map[string]string{
		"C": "3",
		"A": "1",
		"B": "2",
	}
	expect := []string{"A", "B", "C"}
	actual := sortedKeys(m)
	if diff := cmp.Diff(expect, actual); diff != "" {
		t.Errorf("User value is mismatch (-got +want):\n%s", diff)
	}

}

func Test_isFirstLower(t *testing.T) {
	if isFirstLower("test") != true || isFirstLower("Test") != false || isFirstLower("") != false {
		t.Error("isFirstLower Test Failed")
	}
}

func Test_varNormalize(t *testing.T) {
	if varNormalize("") != "" || varNormalize("test") != "Test" || varNormalize("Test_2") != "Test_2" || varNormalize("_Test") != "UNDERLINE_Test" {
		t.Error("varNormalize Test Failed")
	}
}

func Test_forceTypeSetter(t *testing.T) {
	m := map[string]string{
		"ENV_BOOL":      "bool",
		"ENV_INT8":      "int8",
		"ENV_STRING":    "string",
		"ENV_UINT64":    "uint64",
		"ENV_INTERFACE": "interface",
	}
	expect := map[string]model.KindWithSlice{
		"ENV_BOOL":      model.KindWithSlice{Kind: model.Bool, Slice: false},
		"ENV_INT8":      model.KindWithSlice{Kind: model.Int8, Slice: false},
		"ENV_STRING":    model.KindWithSlice{Kind: model.String, Slice: false},
		"ENV_UINT64":    model.KindWithSlice{Kind: model.Uint64, Slice: false},
		"ENV_INTERFACE": model.KindWithSlice{Kind: model.Interface, Slice: false},
	}
	actual, err := forceTypeSetter(m)
	if err != nil {
		t.Errorf(err.Error())
	}
	if diff := cmp.Diff(expect, actual); diff != "" {
		t.Errorf("value is mismatch (-expect +actual):\n%s", diff)
	}
}

func Test_forceTypeSetterSlice(t *testing.T) {
	m := map[string]string{
		"ENV_S_BOOL":      "[]bool",
		"ENV_S_INT8":      "[]int8",
		"ENV_S_STRING":    "[]string",
		"ENV_S_UINT64":    "[]uint64",
		"ENV_S_INTERFACE": "[]interface",
	}
	expect := map[string]model.KindWithSlice{
		"ENV_S_BOOL":      model.KindWithSlice{Kind: model.Bool, Slice: true},
		"ENV_S_INT8":      model.KindWithSlice{Kind: model.Int8, Slice: true},
		"ENV_S_STRING":    model.KindWithSlice{Kind: model.String, Slice: true},
		"ENV_S_UINT64":    model.KindWithSlice{Kind: model.Uint64, Slice: true},
		"ENV_S_INTERFACE": model.KindWithSlice{Kind: model.Interface, Slice: true},
	}
	actual, err := forceTypeSetter(m)
	if err != nil {
		t.Errorf(err.Error())
	}
	if diff := cmp.Diff(expect, actual); diff != "" {
		t.Errorf("value is mismatch (-expect +actual):\n%s", diff)
	}
}

func Test_forceTypeSetterUnsupported(t *testing.T) {
	m := map[string]string{
		"ENV_BOOL":            "bool",
		"ENV_UNSUPPORTEDTYPE": "json",
	}
	_, err := forceTypeSetter(m)
	if err.Error() != "Unsupported Type: ENV_UNSUPPORTEDTYPE=json" {
		t.Error("forceTypeSetterUnsupported Test Failed")
	}
}

func Test_forceTypeSetterUnsupportedSlice(t *testing.T) {
	m := map[string]string{
		"ENV_S_BOOL":            "[]bool",
		"ENV_S_UNSUPPORTEDTYPE": "[]json",
	}
	_, err := forceTypeSetter(m)
	if err.Error() != "Unsupported Type: ENV_S_UNSUPPORTEDTYPE=[]json" {
		t.Error("forceTypeSetterUnsupported Test Failed")
	}
}
