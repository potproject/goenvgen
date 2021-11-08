package gen

import (
	"reflect"
	"sort"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/joho/godotenv"
)

func Gen(fileName string) error {
	f := NewFile("envgen")

	var envs map[string]string
	var err error
	if fileName == "" {
		envs, err = godotenv.Read()
	}
	envs, err = godotenv.Read(fileName)
	if err != nil {
		return err
	}
	// struct & Load
	structCode := make([]Code, len(envs))
	interfaceCode := make([]Code, len(envs))
	setCode := make([]Code, len(envs))
	setDict := Dict{}
	for _, i := range sortedKeys(envs) {
		v := envs[i]
		k, _, isS := checker(v)

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
	setCode = append(setCode, Return(Nil()))
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

	// End...
	return f.Save("envgen/envgen.go")
}

func genGetter(f *File, s string, k reflect.Kind, isS bool) {
	funcS := strings.Title(s)
	i := f.Func().Params(Id("g").Id("getter")).Id(funcS).Params()
	if isS {
		i = i.Index()
	}
	switch k {
	case reflect.Bool:
		i = i.Bool()
	case reflect.Int64:
		i = i.Int64()
	case reflect.Float64:
		i = i.Float64()
	default:
		i = i.String()
	}
	i.Block(
		Return().Id("env." + s),
	)
}

func genSetter(f *File, s string, k reflect.Kind, isS bool) {
	funcS := strings.Title(s)
	i := f.Func().Params(Id("s").Id("setter")).Id(funcS)
	p := Id("value")
	if isS {
		p = p.Index()
	}
	switch k {
	case reflect.Bool:
		p = p.Bool()
		i = i.Params(p)
	case reflect.Int64:
		p = p.Int64()
		i = i.Params(p)
	case reflect.Float64:
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

func genStructCode(s string, k reflect.Kind, isS bool) *Statement {
	i := Id(s)
	if isS {
		i = i.Index()
	}
	switch k {
	case reflect.Bool:
		return i.Bool()
	case reflect.Int64:
		return i.Int64()
	case reflect.Float64:
		return i.Float64()
	default:
		return i.String()
	}
}

func genInterfaceCode(s string, k reflect.Kind, isS bool) *Statement {
	funcS := strings.Title(s)
	i := Id(funcS).Params()
	if isS {
		i = i.Index()
	}
	switch k {
	case reflect.Bool:
		return i.Bool()
	case reflect.Int64:
		return i.Int64()
	case reflect.Float64:
		return i.Float64()
	default:
		return i.String()
	}
}

func genSetCode(s string, k reflect.Kind, isS bool) []Code {
	var codes []Code
	if isS {
		switch k {
		case reflect.Bool:
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
		case reflect.Int64:
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
		case reflect.Float64:
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
		case reflect.Bool:
			codes = append(codes, Id(s).Op(":=").Lit(false))
			codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
			codes = append(codes,
				If(Qual("strings", "ToLower").Call(Id(s+"__S")).Op("==").Lit("true")).Block(
					Id(s).Op("=").Lit(true),
				),
			)
		case reflect.Int64:
			codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
			codes = append(codes, List(Id(s), Err()).Op(":=").Qual("strconv", "ParseInt").Call(Id(s+"__S"), Lit(10), Lit(64)))
			codes = append(codes,
				If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				),
			)
		case reflect.Float64:
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
