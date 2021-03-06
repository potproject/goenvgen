package model

// Kind is Support Types
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Interface
	JSON
)

// KindSupportForceType is types supported by manually type definition
var KindSupportForceType = map[string]Kind{
	"bool":      Bool,
	"int":       Int,
	"int8":      Int8,
	"int16":     Int16,
	"int32":     Int32,
	"int64":     Int64,
	"uint8":     Uint8,
	"uint16":    Uint16,
	"uint32":    Uint32,
	"uint64":    Uint64,
	"float32":   Float32,
	"float64":   Float64,
	"interface": Interface,
	"string":    String,
}

// KindWithSliceAndRequired is struct to determine type and slice and required
type KindWithSliceAndRequired struct {
	Kind     Kind
	Slice    bool
	Required bool
}
