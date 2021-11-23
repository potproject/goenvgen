package main

import (
	"os"
	"path/filepath"

	"github.com/potproject/goenvgen/gen"
)

func main() {
	dir, _ := os.Getwd()
	testfile := filepath.Join(dir, "..", "test", ".env.test")
	testpackage := "envgentest"
	err := gen.GenerateFile(testfile, testpackage, map[string]string{
		"FORCE_BOOL":          "bool",
		"FORCE_INT":           "int",
		"FORCE_INT8":          "int8",
		"FORCE_INT16":         "int16",
		"FORCE_INT32":         "int32",
		"FORCE_INT64":         "int64",
		"FORCE_UINT8":         "uint8",
		"FORCE_UINT16":        "uint16",
		"FORCE_UINT32":        "uint32",
		"FORCE_UINT64":        "uint64",
		"FORCE_FLOAT32":       "float32",
		"FORCE_FLOAT64":       "float64",
		"FORCE_INTERFACE":     "interface",
		"FORCE_S_BOOL":        "[]bool",
		"FORCE_S_INT":         "[]int",
		"FORCE_S_INT8":        "[]int8",
		"FORCE_S_INT16":       "[]int16",
		"FORCE_S_INT32":       "[]int32",
		"FORCE_S_INT64":       "[]int64",
		"FORCE_S_UINT8":       "[]uint8",
		"FORCE_S_UINT16":      "[]uint16",
		"FORCE_S_UINT32":      "[]uint32",
		"FORCE_S_UINT64":      "[]uint64",
		"FORCE_S_FLOAT32":     "[]float32",
		"FORCE_S_FLOAT64":     "[]float64",
		"FORCE_S_INTERFACE":   "[]interface",
		"R_FORCE_INTERFACE":   "interface",
		"R_FORCE_UINT8":       "uint8",
		"R_FORCE_INT8":        "int8",
		"R_FORCE_INT":         "int",
		"R_FORCE_S_INTERFACE": "[]interface",
		"R_FORCE_S_UINT8":     "[]uint8",
		"R_FORCE_S_INT8":      "[]int8",
		"R_FORCE_S_INT":       "[]int",
		"R_FORCE_S_FLOAT64":   "[]float64",
	}, []string{
		"R_STRING_TEST",
		"R_INT64_TEST",
		"R_FLOAT64_TEST",
		"R_BOOL_TEST",
		"R_JSON_TEST",
		"SR_STRING_TEST",
		"SR_INT64_TEST",
		"SR_FLOAT64_TEST",
		"SR_BOOL_TEST",
		"R_FORCE_INTERFACE",
		"R_FORCE_UINT8",
		"R_FORCE_INT8",
		"R_FORCE_INT",
		"R_FORCE_S_INTERFACE",
		"R_FORCE_S_UINT8",
		"R_FORCE_S_INT8",
		"R_FORCE_S_INT",
		"R_FORCE_S_FLOAT64",
	})
	if err != nil {
		panic(err)
	}
}
