package gen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GenerateFile(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage)
	if err != nil {
		t.Error(err.Error())
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateFile_NotExist(t *testing.T) {
	testpackage := "envgen"
	err := GenerateFile("", testpackage)
	if err == nil || err.Error() != "open .env: The system cannot find the file specified." {
		t.Error("Error: Not Exist Pattern Failed")
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateDuplicateEnvTest(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".duplicate.env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage)
	if err == nil || err.Error() != "duplicate Env fields. duplicate_string_test:Duplicate_string_test" {
		t.Error("duplicate validate Failed")
	}
	os.RemoveAll(filepath.Join(".", testpackage))
}

func Test_GenerateUnderlineEnvTest(t *testing.T) {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".underline.env.test")
	testpackage := "envgen"
	err := GenerateFile(testfile, testpackage)
	if err != nil {
		t.Error(err.Error())
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
