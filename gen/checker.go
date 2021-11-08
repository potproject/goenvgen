package gen

import (
	"strconv"
	"strings"

	"github.com/potproject/goenvgen/model"
)

func checker(text string) (model.Kind, interface{}, bool) {
	k, i, isStruct := sliceChecker(text)
	if isStruct {
		return k, i, isStruct
	}
	k, i = typeChecker(text)
	return k, i, isStruct
}

func sliceChecker(text string) (model.Kind, interface{}, bool) {
	if strings.Index(text, ",") != -1 {
		slice := strings.Split(text, ",")
		typeC, _ := typeChecker(slice[0])
		for i := 1; i < len(slice); i++ {
			k, _ := typeChecker(slice[i])
			if typeC != k {
				typeC = model.String
			}
		}
		if typeC == model.Bool {
			var bs []bool
			for _, v := range slice {
				if strings.ToLower(v) == "true" {
					bs = append(bs, true)
				} else {
					bs = append(bs, false)
				}
			}
			return typeC, bs, true
		}
		if typeC == model.Int64 {
			var is []int64
			for _, v := range slice {
				i, _ := strconv.ParseInt(v, 10, 64)
				is = append(is, i)
			}
			return typeC, is, true
		}
		if typeC == model.Float64 {
			var fs []float64
			for _, v := range slice {
				f, _ := strconv.ParseFloat(v, 64)
				fs = append(fs, f)
			}
			return typeC, fs, true
		}
		return typeC, slice, true
	}
	return model.Invalid, 0, false
}

func typeChecker(text string) (model.Kind, interface{}) {
	if strings.ToLower(text) == "true" {
		return model.Bool, true
	}
	if strings.ToLower(text) == "false" {
		return model.Bool, false
	}
	i, err := strconv.ParseInt(text, 10, 64)
	if err == nil {
		return model.Int64, i
	}
	f, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return model.Float64, f
	}
	return model.String, text
}
