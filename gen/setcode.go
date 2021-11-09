package gen

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/potproject/goenvgen/model"
)

func setCodeSliceBool(s string, ks model.KindWithSlice) []Code {
	var codes []Code
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
	return codes
}

func setCodeSliceInt(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
		Qual("os", "Getenv").Call(Lit(s)),
		Lit(","),
	))
	codes = append(codes, Var().Id(s).Index().Int())
	codes = append(codes,
		For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
			List(Id("i"), Id("err")).Op(":=").Qual("strconv", "Atoi").Call(Id("v")),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Id("err")),
			),
			Id(s).Op("=").Append(Id(s), Id("i")),
		),
	)
	return codes
}

func setCodeSliceIntBits(s string, ks model.KindWithSlice, bits int) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
		Qual("os", "Getenv").Call(Lit(s)),
		Lit(","),
	))
	codes = append(codes, Var().Id(s).Index().Id(fmt.Sprintf("int%d", bits)))
	codes = append(codes,
		For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
			List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(Id("v"), Lit(10), Lit(bits)),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Id("err")),
			),
			Id(s).Op("=").Append(Id(s), Id(fmt.Sprintf("int%d", bits)).Call(Id("i"))),
		),
	)
	return codes
}

func setCodeSliceUintBits(s string, ks model.KindWithSlice, bits int) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
		Qual("os", "Getenv").Call(Lit(s)),
		Lit(","),
	))
	codes = append(codes, Var().Id(s).Index().Id(fmt.Sprintf("uint%d", bits)))
	codes = append(codes,
		For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
			List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseUint").Call(Id("v"), Lit(10), Lit(bits)),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Id("err")),
			),
			Id(s).Op("=").Append(Id(s), Id(fmt.Sprintf("uint%d", bits)).Call(Id("i"))),
		),
	)
	return codes
}

func setCodeSliceFloatBits(s string, ks model.KindWithSlice, bits int) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
		Qual("os", "Getenv").Call(Lit(s)),
		Lit(","),
	))
	codes = append(codes, Var().Id(s).Index().Id(fmt.Sprintf("float%d", bits)))
	codes = append(codes,
		For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
			List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseFloat").Call(Id("v"), Lit(bits)),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Id("err")),
			),
			Id(s).Op("=").Append(Id(s), Id(fmt.Sprintf("float%d", bits)).Call(Id("i"))),
		),
	)
	return codes
}

func setCodeSliceInterface(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__A").Op(":=").Qual("strings", "Split").Call(
		Qual("os", "Getenv").Call(Lit(s)),
		Lit(","),
	))
	codes = append(codes, Var().Id(s).Index().Interface())
	codes = append(codes,
		For(List(Id("_"), Id("v")).Op(":=").Range().Id(s+"__A")).Block(
			Id(s).Op("=").Append(Id(s), Id("v")),
		),
	)
	return codes
}

func setCodeSliceString(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s).Op(":=").Qual("strings", "Split").Call(
		Qual("os", "Getenv").Call(Lit(s)),
		Lit(","),
	))
	return codes
}

func setCodeJSON(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	codes = append(codes, Var().Id(s).Id(varNormalize(s)))
	codes = append(codes, Err().Op("=").Qual("encoding/json", "Unmarshal").Call(Id("[]byte").Call(Id(s+"__S")), Op("&").Id(s)))
	return codes
}

func setCodeBool(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s).Op(":=").Lit(false))
	codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	codes = append(codes,
		If(Qual("strings", "ToLower").Call(Id(s+"__S")).Op("==").Lit("true")).Block(
			Id(s).Op("=").Lit(true),
		),
	)
	return codes
}

func setCodeInt(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	codes = append(codes, List(Id(s), Err()).Op(":=").Qual("strconv", "Atoi").Call(Id(s+"__S")))
	codes = append(codes,
		If(Id("err").Op("!=").Nil()).Block(
			Return(Id("err")),
		),
	)
	return codes
}

func setCodeIntBits(s string, ks model.KindWithSlice, bits int) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	codes = append(codes, List(Id(s+"__64"), Err()).Op(":=").Qual("strconv", "ParseInt").Call(Id(s+"__S"), Lit(10), Lit(bits)))
	codes = append(codes,
		If(Id("err").Op("!=").Nil()).Block(
			Return(Id("err")),
		),
	)
	codes = append(codes, Id(s).Op(":=").Id(fmt.Sprintf("int%d", bits)).Call(Id(s+"__64")))
	return codes
}

func setCodeUintBits(s string, ks model.KindWithSlice, bits int) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	codes = append(codes, List(Id(s+"__64"), Err()).Op(":=").Qual("strconv", "ParseUint").Call(Id(s+"__S"), Lit(10), Lit(bits)))
	codes = append(codes,
		If(Id("err").Op("!=").Nil()).Block(
			Return(Id("err")),
		),
	)
	codes = append(codes, Id(s).Op(":=").Id(fmt.Sprintf("uint%d", bits)).Call(Id(s+"__64")))
	return codes
}

func setCodeFloatBits(s string, ks model.KindWithSlice, bits int) []Code {
	var codes []Code
	codes = append(codes, Id(s+"__S").Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	codes = append(codes, List(Id(s+"__64"), Err()).Op(":=").Qual("strconv", "ParseFloat").Call(Id(s+"__S"), Lit(bits)))
	codes = append(codes,
		If(Id("err").Op("!=").Nil()).Block(
			Return(Id("err")),
		),
	)
	codes = append(codes, Id(s).Op(":=").Id(fmt.Sprintf("float%d", bits)).Call(Id(s+"__64")))
	return codes
}

func setCodeInterface(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Var().Id(s).Interface().Op("=").Qual("os", "Getenv").Call(Lit(s)))
	return codes
}

func setCodeString(s string, ks model.KindWithSlice) []Code {
	var codes []Code
	codes = append(codes, Id(s).Op(":=").Qual("os", "Getenv").Call(Lit(s)))
	return codes
}

func genSetCode(s string, ks model.KindWithSlice) []Code {
	if ks.Slice {
		switch ks.Kind {
		case model.Bool:
			return setCodeSliceBool(s, ks)
		case model.Int:
			return setCodeSliceInt(s, ks)
		case model.Int8:
			return setCodeSliceIntBits(s, ks, 8)
		case model.Int16:
			return setCodeSliceIntBits(s, ks, 16)
		case model.Int32:
			return setCodeSliceIntBits(s, ks, 32)
		case model.Int64:
			return setCodeSliceIntBits(s, ks, 64)
		case model.Uint8:
			return setCodeSliceUintBits(s, ks, 8)
		case model.Uint16:
			return setCodeSliceUintBits(s, ks, 16)
		case model.Uint32:
			return setCodeSliceUintBits(s, ks, 32)
		case model.Uint64:
			return setCodeSliceUintBits(s, ks, 64)
		case model.Float32:
			return setCodeSliceFloatBits(s, ks, 32)
		case model.Float64:
			return setCodeSliceFloatBits(s, ks, 64)
		case model.Interface:
			return setCodeSliceInterface(s, ks)
		default:
			return setCodeSliceString(s, ks)
		}
	} else {
		switch ks.Kind {
		case model.JSON:
			return setCodeJSON(s, ks)
		case model.Bool:
			return setCodeBool(s, ks)
		case model.Int:
			return setCodeInt(s, ks)
		case model.Int8:
			return setCodeIntBits(s, ks, 8)
		case model.Int16:
			return setCodeIntBits(s, ks, 16)
		case model.Int32:
			return setCodeIntBits(s, ks, 32)
		case model.Int64:
			return setCodeIntBits(s, ks, 64)
		case model.Uint8:
			return setCodeUintBits(s, ks, 8)
		case model.Uint16:
			return setCodeUintBits(s, ks, 16)
		case model.Uint32:
			return setCodeUintBits(s, ks, 32)
		case model.Uint64:
			return setCodeUintBits(s, ks, 64)
		case model.Float32:
			return setCodeFloatBits(s, ks, 32)
		case model.Float64:
			return setCodeFloatBits(s, ks, 64)
		case model.Interface:
			return setCodeInterface(s, ks)
		default:
			return setCodeString(s, ks)
		}
	}
}
