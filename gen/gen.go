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

func GenerateFile(fileName string, packageName string) error {
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
	files, err := Generate(envs, packageName)
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

func Generate(envs map[string]string, packageName string) (map[string][]byte, error) {
	outputs := map[string][]byte{}

	// sort & validate
	sortedEnvs := sortedKeys(envs)
	err := validate(sortedEnvs)
	if err != nil {
		return outputs, err
	}

	f := NewFile(packageName)
	f.ImportName("encoding/json", "json")

	// struct & Load
	structCode := make([]Code, len(envs))
	interfaceCode := make([]Code, len(envs))
	setCode := make([]Code, len(envs))
	setDict := Dict{}
	for _, i := range sortedEnvs {
		v := envs[i]
		k, _, isS := checker(v)
		if k == model.JSON {
			on, o := genStructJSON(i, packageName, v)
			outputs[on] = o
		}
		structCode = append(structCode, genStructCode(i, k, isS))
		interfaceCode = append(interfaceCode, genInterfaceCode(i, k, isS))
		setCode = append(setCode, genSetCode(i, k, isS)...)
		genGetter(f, i, k, isS)
		genSetter(f, i, k, isS)
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

func genGetter(f *File, s string, k model.Kind, isS bool) {
	funcS := varNormalize(s)
	i := f.Func().Params(Id("g").Id("getter")).Id(funcS).Params()
	if isS {
		i = i.Index()
	}
	switch k {
	case model.JSON:
		i = i.Id(varNormalize(s))
	case model.Bool:
		i = i.Bool()
	case model.Int64:
		i = i.Int64()
	case model.Float64:
		i = i.Float64()
	default:
		i = i.String()
	}
	i.Block(
		Return().Id("env." + s),
	)
}

func genSetter(f *File, s string, k model.Kind, isS bool) {
	funcS := varNormalize(s)
	i := f.Func().Params(Id("s").Id("setter")).Id(funcS)
	p := Id("value")
	if isS {
		p = p.Index()
	}
	switch k {
	case model.JSON:
		p = p.Id(varNormalize(s))
		i = i.Params(p)
	case model.Bool:
		p = p.Bool()
		i = i.Params(p)
	case model.Int64:
		p = p.Int64()
		i = i.Params(p)
	case model.Float64:
		p = p.Float64()
		i = i.Params(p)
	default:
		p = p.String()
		i = i.Params(p)
	}
	i.Block(
		Id("env."+s).Op("=").Id("value"),
		Return(),
	)
}

func genStructCode(s string, k model.Kind, isS bool) *Statement {
	i := Id(s)
	if isS {
		i = i.Index()
	}
	switch k {
	case model.JSON:
		return i.Id(varNormalize(s))
	case model.Bool:
		return i.Bool()
	case model.Int64:
		return i.Int64()
	case model.Float64:
		return i.Float64()
	default:
		return i.String()
	}
}

func genStructJSON(s string, pkgName string, body string) (outputName string, output []byte) {
	structName := varNormalize(s)
	input := strings.NewReader(body)
	tagList := []string{"json"}
	output, _ = gojson.Generate(input, gojson.ParseJson, structName, pkgName, tagList, false, true)
	outputName = filepath.Join(".", pkgName, s+".go")
	return
}

func genInterfaceCode(s string, k model.Kind, isS bool) *Statement {
	funcS := varNormalize(s)
	i := Id(funcS).Params()
	if isS {
		i = i.Index()
	}
	switch k {
	case model.JSON:
		return i.Id(varNormalize(s))
	case model.Bool:
		return i.Bool()
	case model.Int64:
		return i.Int64()
	case model.Float64:
		return i.Float64()
	default:
		return i.String()
	}
}

func genSetCode(s string, k model.Kind, isS bool) []Code {
	var codes []Code
	if isS {
		switch k {
		case model.Bool:
			codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
				Qual("os", "Getenv").Call(Lit(s)),
				Lit(","),
			))
			codes = append(codes, Var().Id(s).Index().Bool())
			codes = append(codes,
				For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
					If(Qual("strings", "ToLower").Call(Id("v")).Op("==").Lit("true")).Block(
						Id(s).Op("=").Append(Id(s), Lit(true)),
					).Else().Block(
						Id(s).Op("=").Append(Id(s), Lit(false)),
					),
				),
			)
		case model.Int64:
			codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
				Qual("os", "Getenv").Call(Lit(s)),
				Lit(","),
			))
			codes = append(codes, Var().Id(s).Index().Int64())
			codes = append(codes,
				For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
					List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(Id("v"), Lit(10), Lit(64)),
					If(Id("err").Op("!=").Nil()).Block(
						Return(Id("err")),
					),
					Id(s).Op("=").Append(Id(s), Id("i")),
				),
			)
		case model.Float64:
			codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
				Qual("os", "Getenv").Call(Lit(s)),
				Lit(","),
			))
			codes = append(codes, Var().Id(s).Index().Float64())
			codes = append(codes,
				For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
					List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseFloat").Call(Id("v"), Lit(64)),
					If(Id("err").Op("!=").Nil()).Block(
						Return(Id("err")),
					),
					Id(s).Op("=").Append(Id(s), Id("i")),
				),
			)
		default:
			codes = append(codes, Id(s).Op(":=").Qual("strings", "Split").Call(
				Qual("os", "Getenv").Call(Lit(s)),
				Lit(","),
			))
		}
	} else {
		switch k {
		case model.JSON:
			codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
			codes = append(codes, Var().Id(s).Id(varNormalize(s)))
			codes = append(codes, Err().Op("=").Qual("encoding/json", "Unmarshal").Call(Id("[]byte").Call(Id(s+"__S")), Op("&").Id(s)))
		case model.Bool:
			codes = append(codes, Id(s).Op(":=").Lit(false))
			codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
			codes = append(codes,
				If(Qual("strings", "ToLower").Call(Id(s+"__S")).Op("==").Lit("true")).Block(
					Id(s).Op("=").Lit(true),
				),
			)
		case model.Int64:
			codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
			codes = append(codes, List(Id(s), Err()).Op(":=").Qual("strconv", "ParseInt").Call(Id(s+"__S"), Lit(10), Lit(64)))
			codes = append(codes,
				If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				),
			)
		case model.Float64:
			codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
			codes = append(codes, List(Id(s), Err()).Op(":=").Qual("strconv", "ParseFloat").Call(Id(s+"__S"), Lit(64)))
			codes = append(codes,
				If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				),
			)
		default:
			codes = append(codes, Id(s).Op(":=").Qual("os", "Getenv").Call(Lit(s)))
		}
	}
	return codes
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
