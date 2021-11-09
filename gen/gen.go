package gen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/ChimeraCoder/gojson"
	. "github.com/dave/jennifer/jen"
	"github.com/joho/godotenv"
	"github.com/potproject/goenvgen/model"
)

func GenerateFile(fileName string, packageName string, forceType map[string]string) error {
	var envs map[string]string
	var err error
	if fileName == "" {
		envs, err = godotenv.Read()
	} else {
		envs, err = godotenv.Read(fileName)
	}
	if err != nil {
		return err
	}
	files, err := Generate(envs, packageName, forceType)
	if err != nil {
		return err
	}

	// make Directory & File
	p := filepath.Join(".", packageName)
	os.MkdirAll(p, os.ModePerm)
	for path, body := range files {
		ioutil.WriteFile(path, body, 0644)
	}
	return nil
}

func Generate(envs map[string]string, packageName string, forceType map[string]string) (map[string][]byte, error) {
	outputs := map[string][]byte{}

	// sort & validate
	sortedEnvs := sortedKeys(envs)
	err := validate(sortedEnvs)
	if err != nil {
		return outputs, err
	}

	f := NewFile(packageName)
	f.ImportName("encoding/json", "json")

	// Force Set Type
	forceSetKind, err := forceTypeSetter(forceType)
	if err != nil {
		return outputs, err
	}

	// struct & Load
	structCode := make([]Code, len(envs))
	interfaceCode := make([]Code, len(envs))
	setCode := make([]Code, len(envs))
	setDict := Dict{}
	for _, i := range sortedEnvs {
		v := envs[i]
		ks, ok := forceSetKind[i]
		if !ok {
			ks, _ = checker(v)
		}
		if ks.Kind == model.JSON {
			on, o := genStructJSON(i, packageName, v)
			outputs[on] = o
		}
		structCode = append(structCode, genStructCode(i, ks))
		interfaceCode = append(interfaceCode, genInterfaceCode(i, ks))
		setCode = append(setCode, genSetCode(i, ks)...)
		genGetter(f, i, ks)
		genSetter(f, i, ks)
		setDict[Id(i)] = Id(i)
	}

	f.Type().Id("environment").Struct(structCode...)
	f.Var().Id("env").Id("environment")
	setCode = append(setCode,
		Id("env").Op("=").Id("environment").Values(setDict),
	)
	setCode = append([]Code{Var().Id("err").Id("error")}, setCode...)
	setCode = append(setCode, Return(Id("err")))
	f.Func().Id("Load").Params().Error().Block(
		setCode...,
	)

	// GetterInterface
	f.Type().Id("getterInterface").Interface(
		interfaceCode...,
	)

	// Getter
	f.Type().Id("getter").Struct(
		Id("getterInterface"),
	)

	// Get
	f.Func().Id("Get").Params().Id("getter").Block(
		Return().Id("getter").Block(),
	)

	// SetterInterface
	f.Type().Id("setterInterface").Interface(
		interfaceCode...,
	)

	// Setter
	f.Type().Id("setter").Struct(
		Id("setterInterface"),
	)

	// Set
	f.Func().Id("Set").Params().Id("setter").Block(
		Return().Id("setter").Block(),
	)

	// Render...
	buf := &bytes.Buffer{}
	err = f.Render(buf)
	if err != nil {
		return outputs, err
	}
	fileName := filepath.Join(".", packageName, packageName+".go")
	outputs[fileName] = buf.Bytes()
	return outputs, err
}

func genGetter(f *File, s string, ks model.KindWithSlice) {
	funcS := varNormalize(s)
	i := f.Func().Params(Id("g").Id("getter")).Id(funcS).Params()
	if ks.Slice {
		i = i.Index()
	}
	i = kindToStatement(s, ks, i)
	i.Block(
		Return().Id("env." + s),
	)
}

func genSetter(f *File, s string, ks model.KindWithSlice) {
	funcS := varNormalize(s)
	i := f.Func().Params(Id("s").Id("setter")).Id(funcS)
	p := Id("value")
	if ks.Slice {
		p = p.Index()
	}
	p = kindToStatement(s, ks, p)
	i = i.Params(p)
	i.Block(
		Id("env."+s).Op("=").Id("value"),
		Return(),
	)
}

func genStructCode(s string, ks model.KindWithSlice) *Statement {
	i := Id(s)
	if ks.Slice {
		i = i.Index()
	}
	return kindToStatement(s, ks, i)
}

func genStructJSON(s string, pkgName string, body string) (outputName string, output []byte) {
	structName := varNormalize(s)
	input := strings.NewReader(body)
	tagList := []string{"json"}
	output, _ = gojson.Generate(input, gojson.ParseJson, structName, pkgName, tagList, false, true)
	outputName = filepath.Join(".", pkgName, s+".go")
	return
}

func genInterfaceCode(s string, ks model.KindWithSlice) *Statement {
	funcS := varNormalize(s)
	i := Id(funcS).Params()
	if ks.Slice {
		i = i.Index()
	}
	return kindToStatement(s, ks, i)
}

func sortedKeys(m interface{}) []string {
	v := reflect.ValueOf(m).MapKeys()
	r := make([]string, len(v))
	for i, v1 := range v {
		r[i] = v1.String()
	}
	sort.Strings(r)
	return r
}

func validate(keys []string) error {
	for _, v := range keys {
		// duplicate check
		if isFirstLower(v) {
			vt := strings.Title(v)
			i := sort.SearchStrings(keys, vt)
			if i != len(keys) && keys[i] == vt {
				return fmt.Errorf("duplicate Env fields. %s:%s", v, vt)
			}
		}
	}
	return nil
}

func isFirstLower(v string) bool {
	if v == "" {
		return false
	}
	r := rune(v[0])
	return unicode.IsLower(r)
}

func varNormalize(v string) string {
	if v == "" {
		return ""
	}
	if strings.HasPrefix(v, "_") {
		return "UNDERLINE" + v
	}
	return strings.Title(v)
}

func forceTypeSetter(m map[string]string) (map[string]model.KindWithSlice, error) {
	r := map[string]model.KindWithSlice{}
	for i, v := range m {
		slice := false
		if strings.HasPrefix(v, "[]") {
			v = v[2:]
			slice = true
		}
		k, ok := model.KindSupportForceType[v]
		if ok {
			r[i] = model.KindWithSlice{
				Kind:  k,
				Slice: slice,
			}
		} else {
			return r, fmt.Errorf("Unsupported Type: %s=%s", i, m[i])
		}
	}
	return r, nil
}

func kindToStatement(s string, ks model.KindWithSlice, i *Statement) *Statement {
	switch ks.Kind {
	case model.Bool:
		return i.Bool()
	case model.Int:
		return i.Int()
	case model.Int8:
		return i.Int8()
	case model.Int16:
		return i.Int16()
	case model.Int32:
		return i.Int32()
	case model.Int64:
		return i.Int64()
	case model.Uint8:
		return i.Uint8()
	case model.Uint16:
		return i.Uint16()
	case model.Uint32:
		return i.Uint32()
	case model.Uint64:
		return i.Uint64()
	case model.Float32:
		return i.Float32()
	case model.Float64:
		return i.Float64()
	case model.Interface:
		return i.Interface()
	case model.JSON:
		return i.Id(varNormalize(s))
	default:
		return i.String()
	}
}
