package envgentest

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func (g getter) BOOL_TEST1() bool {
	return env.BOOL_TEST1
}
func (s setter) BOOL_TEST1(value bool) {
	env.BOOL_TEST1 = value
	return
}
func (g getter) BOOL_TEST2() bool {
	return env.BOOL_TEST2
}
func (s setter) BOOL_TEST2(value bool) {
	env.BOOL_TEST2 = value
	return
}
func (g getter) BOOL_TEST3() bool {
	return env.BOOL_TEST3
}
func (s setter) BOOL_TEST3(value bool) {
	env.BOOL_TEST3 = value
	return
}
func (g getter) CamelCaseTest() string {
	return env.CamelCaseTest
}
func (s setter) CamelCaseTest(value string) {
	env.CamelCaseTest = value
	return
}
func (g getter) FLOAT64_TEST1() float64 {
	return env.FLOAT64_TEST1
}
func (s setter) FLOAT64_TEST1(value float64) {
	env.FLOAT64_TEST1 = value
	return
}
func (g getter) FLOAT64_TEST2() float64 {
	return env.FLOAT64_TEST2
}
func (s setter) FLOAT64_TEST2(value float64) {
	env.FLOAT64_TEST2 = value
	return
}
func (g getter) FLOAT64_TEST3() float64 {
	return env.FLOAT64_TEST3
}
func (s setter) FLOAT64_TEST3(value float64) {
	env.FLOAT64_TEST3 = value
	return
}
func (g getter) FORCE_BOOL() bool {
	return env.FORCE_BOOL
}
func (s setter) FORCE_BOOL(value bool) {
	env.FORCE_BOOL = value
	return
}
func (g getter) FORCE_FLOAT32() float32 {
	return env.FORCE_FLOAT32
}
func (s setter) FORCE_FLOAT32(value float32) {
	env.FORCE_FLOAT32 = value
	return
}
func (g getter) FORCE_FLOAT64() float64 {
	return env.FORCE_FLOAT64
}
func (s setter) FORCE_FLOAT64(value float64) {
	env.FORCE_FLOAT64 = value
	return
}
func (g getter) FORCE_INT() int {
	return env.FORCE_INT
}
func (s setter) FORCE_INT(value int) {
	env.FORCE_INT = value
	return
}
func (g getter) FORCE_INT16() int16 {
	return env.FORCE_INT16
}
func (s setter) FORCE_INT16(value int16) {
	env.FORCE_INT16 = value
	return
}
func (g getter) FORCE_INT32() int32 {
	return env.FORCE_INT32
}
func (s setter) FORCE_INT32(value int32) {
	env.FORCE_INT32 = value
	return
}
func (g getter) FORCE_INT64() int64 {
	return env.FORCE_INT64
}
func (s setter) FORCE_INT64(value int64) {
	env.FORCE_INT64 = value
	return
}
func (g getter) FORCE_INT8() int8 {
	return env.FORCE_INT8
}
func (s setter) FORCE_INT8(value int8) {
	env.FORCE_INT8 = value
	return
}
func (g getter) FORCE_INTERFACE() interface{} {
	return env.FORCE_INTERFACE
}
func (s setter) FORCE_INTERFACE(value interface{}) {
	env.FORCE_INTERFACE = value
	return
}
func (g getter) FORCE_S_BOOL() []bool {
	return env.FORCE_S_BOOL
}
func (s setter) FORCE_S_BOOL(value []bool) {
	env.FORCE_S_BOOL = value
	return
}
func (g getter) FORCE_S_FLOAT32() []float32 {
	return env.FORCE_S_FLOAT32
}
func (s setter) FORCE_S_FLOAT32(value []float32) {
	env.FORCE_S_FLOAT32 = value
	return
}
func (g getter) FORCE_S_FLOAT64() []float64 {
	return env.FORCE_S_FLOAT64
}
func (s setter) FORCE_S_FLOAT64(value []float64) {
	env.FORCE_S_FLOAT64 = value
	return
}
func (g getter) FORCE_S_INT() []int {
	return env.FORCE_S_INT
}
func (s setter) FORCE_S_INT(value []int) {
	env.FORCE_S_INT = value
	return
}
func (g getter) FORCE_S_INT16() []int16 {
	return env.FORCE_S_INT16
}
func (s setter) FORCE_S_INT16(value []int16) {
	env.FORCE_S_INT16 = value
	return
}
func (g getter) FORCE_S_INT32() []int32 {
	return env.FORCE_S_INT32
}
func (s setter) FORCE_S_INT32(value []int32) {
	env.FORCE_S_INT32 = value
	return
}
func (g getter) FORCE_S_INT64() []int64 {
	return env.FORCE_S_INT64
}
func (s setter) FORCE_S_INT64(value []int64) {
	env.FORCE_S_INT64 = value
	return
}
func (g getter) FORCE_S_INT8() []int8 {
	return env.FORCE_S_INT8
}
func (s setter) FORCE_S_INT8(value []int8) {
	env.FORCE_S_INT8 = value
	return
}
func (g getter) FORCE_S_INTERFACE() []interface{} {
	return env.FORCE_S_INTERFACE
}
func (s setter) FORCE_S_INTERFACE(value []interface{}) {
	env.FORCE_S_INTERFACE = value
	return
}
func (g getter) FORCE_S_UINT16() []uint16 {
	return env.FORCE_S_UINT16
}
func (s setter) FORCE_S_UINT16(value []uint16) {
	env.FORCE_S_UINT16 = value
	return
}
func (g getter) FORCE_S_UINT32() []uint32 {
	return env.FORCE_S_UINT32
}
func (s setter) FORCE_S_UINT32(value []uint32) {
	env.FORCE_S_UINT32 = value
	return
}
func (g getter) FORCE_S_UINT64() []uint64 {
	return env.FORCE_S_UINT64
}
func (s setter) FORCE_S_UINT64(value []uint64) {
	env.FORCE_S_UINT64 = value
	return
}
func (g getter) FORCE_S_UINT8() []uint8 {
	return env.FORCE_S_UINT8
}
func (s setter) FORCE_S_UINT8(value []uint8) {
	env.FORCE_S_UINT8 = value
	return
}
func (g getter) FORCE_UINT16() uint16 {
	return env.FORCE_UINT16
}
func (s setter) FORCE_UINT16(value uint16) {
	env.FORCE_UINT16 = value
	return
}
func (g getter) FORCE_UINT32() uint32 {
	return env.FORCE_UINT32
}
func (s setter) FORCE_UINT32(value uint32) {
	env.FORCE_UINT32 = value
	return
}
func (g getter) FORCE_UINT64() uint64 {
	return env.FORCE_UINT64
}
func (s setter) FORCE_UINT64(value uint64) {
	env.FORCE_UINT64 = value
	return
}
func (g getter) FORCE_UINT8() uint8 {
	return env.FORCE_UINT8
}
func (s setter) FORCE_UINT8(value uint8) {
	env.FORCE_UINT8 = value
	return
}
func (g getter) INT64_TEST1() int64 {
	return env.INT64_TEST1
}
func (s setter) INT64_TEST1(value int64) {
	env.INT64_TEST1 = value
	return
}
func (g getter) INT64_TEST2() int64 {
	return env.INT64_TEST2
}
func (s setter) INT64_TEST2(value int64) {
	env.INT64_TEST2 = value
	return
}
func (g getter) INT64_TEST3() int64 {
	return env.INT64_TEST3
}
func (s setter) INT64_TEST3(value int64) {
	env.INT64_TEST3 = value
	return
}
func (g getter) INVALID_JSON_TEST() string {
	return env.INVALID_JSON_TEST
}
func (s setter) INVALID_JSON_TEST(value string) {
	env.INVALID_JSON_TEST = value
	return
}
func (g getter) JSON_TEST() JSON_TEST {
	return env.JSON_TEST
}
func (s setter) JSON_TEST(value JSON_TEST) {
	env.JSON_TEST = value
	return
}
func (g getter) JSON_TEST2() JSON_TEST2 {
	return env.JSON_TEST2
}
func (s setter) JSON_TEST2(value JSON_TEST2) {
	env.JSON_TEST2 = value
	return
}
func (g getter) JSON_TEST3() JSON_TEST3 {
	return env.JSON_TEST3
}
func (s setter) JSON_TEST3(value JSON_TEST3) {
	env.JSON_TEST3 = value
	return
}
func (g getter) STRING_TEST1() string {
	return env.STRING_TEST1
}
func (s setter) STRING_TEST1(value string) {
	env.STRING_TEST1 = value
	return
}
func (g getter) STRING_TEST2() string {
	return env.STRING_TEST2
}
func (s setter) STRING_TEST2(value string) {
	env.STRING_TEST2 = value
	return
}
func (g getter) STRING_TEST3() string {
	return env.STRING_TEST3
}
func (s setter) STRING_TEST3(value string) {
	env.STRING_TEST3 = value
	return
}
func (g getter) S_BOOL_TEST() []bool {
	return env.S_BOOL_TEST
}
func (s setter) S_BOOL_TEST(value []bool) {
	env.S_BOOL_TEST = value
	return
}
func (g getter) S_FLOAT64_TEST() []float64 {
	return env.S_FLOAT64_TEST
}
func (s setter) S_FLOAT64_TEST(value []float64) {
	env.S_FLOAT64_TEST = value
	return
}
func (g getter) S_INT64_TEST() []int64 {
	return env.S_INT64_TEST
}
func (s setter) S_INT64_TEST(value []int64) {
	env.S_INT64_TEST = value
	return
}
func (g getter) S_STRING_TEST() []string {
	return env.S_STRING_TEST
}
func (s setter) S_STRING_TEST(value []string) {
	env.S_STRING_TEST = value
	return
}
func (g getter) S_STRING_TEST2() []string {
	return env.S_STRING_TEST2
}
func (s setter) S_STRING_TEST2(value []string) {
	env.S_STRING_TEST2 = value
	return
}
func (g getter) UPPERCASETEST() string {
	return env.UPPERCASETEST
}
func (s setter) UPPERCASETEST(value string) {
	env.UPPERCASETEST = value
	return
}
func (g getter) UNDERLINE_TEST() string {
	return env._TEST
}
func (s setter) UNDERLINE_TEST(value string) {
	env._TEST = value
	return
}
func (g getter) Lower_snake_case_test() string {
	return env.lower_snake_case_test
}
func (s setter) Lower_snake_case_test(value string) {
	env.lower_snake_case_test = value
	return
}
func (g getter) Lower_snake_case_test_json() Lower_snake_case_test_json {
	return env.lower_snake_case_test_json
}
func (s setter) Lower_snake_case_test_json(value Lower_snake_case_test_json) {
	env.lower_snake_case_test_json = value
	return
}
func (g getter) Lowercasetest() string {
	return env.lowercasetest
}
func (s setter) Lowercasetest(value string) {
	env.lowercasetest = value
	return
}

type environment struct {
	BOOL_TEST1                 bool
	BOOL_TEST2                 bool
	BOOL_TEST3                 bool
	CamelCaseTest              string
	FLOAT64_TEST1              float64
	FLOAT64_TEST2              float64
	FLOAT64_TEST3              float64
	FORCE_BOOL                 bool
	FORCE_FLOAT32              float32
	FORCE_FLOAT64              float64
	FORCE_INT                  int
	FORCE_INT16                int16
	FORCE_INT32                int32
	FORCE_INT64                int64
	FORCE_INT8                 int8
	FORCE_INTERFACE            interface{}
	FORCE_S_BOOL               []bool
	FORCE_S_FLOAT32            []float32
	FORCE_S_FLOAT64            []float64
	FORCE_S_INT                []int
	FORCE_S_INT16              []int16
	FORCE_S_INT32              []int32
	FORCE_S_INT64              []int64
	FORCE_S_INT8               []int8
	FORCE_S_INTERFACE          []interface{}
	FORCE_S_UINT16             []uint16
	FORCE_S_UINT32             []uint32
	FORCE_S_UINT64             []uint64
	FORCE_S_UINT8              []uint8
	FORCE_UINT16               uint16
	FORCE_UINT32               uint32
	FORCE_UINT64               uint64
	FORCE_UINT8                uint8
	INT64_TEST1                int64
	INT64_TEST2                int64
	INT64_TEST3                int64
	INVALID_JSON_TEST          string
	JSON_TEST                  JSON_TEST
	JSON_TEST2                 JSON_TEST2
	JSON_TEST3                 JSON_TEST3
	STRING_TEST1               string
	STRING_TEST2               string
	STRING_TEST3               string
	S_BOOL_TEST                []bool
	S_FLOAT64_TEST             []float64
	S_INT64_TEST               []int64
	S_STRING_TEST              []string
	S_STRING_TEST2             []string
	UPPERCASETEST              string
	_TEST                      string
	lower_snake_case_test      string
	lower_snake_case_test_json Lower_snake_case_test_json
	lowercasetest              string
}

var env environment

func Load() error {
	var err error
	BOOL_TEST1 := false
	BOOL_TEST1__S := os.Getenv("BOOL_TEST1")
	if strings.ToLower(BOOL_TEST1__S) == "true" {
		BOOL_TEST1 = true
	}
	BOOL_TEST2 := false
	BOOL_TEST2__S := os.Getenv("BOOL_TEST2")
	if strings.ToLower(BOOL_TEST2__S) == "true" {
		BOOL_TEST2 = true
	}
	BOOL_TEST3 := false
	BOOL_TEST3__S := os.Getenv("BOOL_TEST3")
	if strings.ToLower(BOOL_TEST3__S) == "true" {
		BOOL_TEST3 = true
	}
	CamelCaseTest := os.Getenv("CamelCaseTest")
	FLOAT64_TEST1__S := os.Getenv("FLOAT64_TEST1")
	FLOAT64_TEST1__64, err := strconv.ParseFloat(FLOAT64_TEST1__S, 64)
	if err != nil {
		return err
	}
	FLOAT64_TEST1 := float64(FLOAT64_TEST1__64)
	FLOAT64_TEST2__S := os.Getenv("FLOAT64_TEST2")
	FLOAT64_TEST2__64, err := strconv.ParseFloat(FLOAT64_TEST2__S, 64)
	if err != nil {
		return err
	}
	FLOAT64_TEST2 := float64(FLOAT64_TEST2__64)
	FLOAT64_TEST3__S := os.Getenv("FLOAT64_TEST3")
	FLOAT64_TEST3__64, err := strconv.ParseFloat(FLOAT64_TEST3__S, 64)
	if err != nil {
		return err
	}
	FLOAT64_TEST3 := float64(FLOAT64_TEST3__64)
	FORCE_BOOL := false
	FORCE_BOOL__S := os.Getenv("FORCE_BOOL")
	if strings.ToLower(FORCE_BOOL__S) == "true" {
		FORCE_BOOL = true
	}
	FORCE_FLOAT32__S := os.Getenv("FORCE_FLOAT32")
	FORCE_FLOAT32__64, err := strconv.ParseFloat(FORCE_FLOAT32__S, 32)
	if err != nil {
		return err
	}
	FORCE_FLOAT32 := float32(FORCE_FLOAT32__64)
	FORCE_FLOAT64__S := os.Getenv("FORCE_FLOAT64")
	FORCE_FLOAT64__64, err := strconv.ParseFloat(FORCE_FLOAT64__S, 64)
	if err != nil {
		return err
	}
	FORCE_FLOAT64 := float64(FORCE_FLOAT64__64)
	FORCE_INT__S := os.Getenv("FORCE_INT")
	FORCE_INT, err := strconv.Atoi(FORCE_INT__S)
	if err != nil {
		return err
	}
	FORCE_INT16__S := os.Getenv("FORCE_INT16")
	FORCE_INT16__64, err := strconv.ParseInt(FORCE_INT16__S, 10, 16)
	if err != nil {
		return err
	}
	FORCE_INT16 := int16(FORCE_INT16__64)
	FORCE_INT32__S := os.Getenv("FORCE_INT32")
	FORCE_INT32__64, err := strconv.ParseInt(FORCE_INT32__S, 10, 32)
	if err != nil {
		return err
	}
	FORCE_INT32 := int32(FORCE_INT32__64)
	FORCE_INT64__S := os.Getenv("FORCE_INT64")
	FORCE_INT64__64, err := strconv.ParseInt(FORCE_INT64__S, 10, 64)
	if err != nil {
		return err
	}
	FORCE_INT64 := int64(FORCE_INT64__64)
	FORCE_INT8__S := os.Getenv("FORCE_INT8")
	FORCE_INT8__64, err := strconv.ParseInt(FORCE_INT8__S, 10, 8)
	if err != nil {
		return err
	}
	FORCE_INT8 := int8(FORCE_INT8__64)
	var FORCE_INTERFACE interface{} = os.Getenv("FORCE_INTERFACE")
	FORCE_S_BOOL__A := strings.Split(os.Getenv("FORCE_S_BOOL"), ",")
	var FORCE_S_BOOL []bool
	for _, v := range FORCE_S_BOOL__A {
		if strings.ToLower(v) == "true" {
			FORCE_S_BOOL = append(FORCE_S_BOOL, true)
		} else {
			FORCE_S_BOOL = append(FORCE_S_BOOL, false)
		}
	}
	FORCE_S_FLOAT32__A := strings.Split(os.Getenv("FORCE_S_FLOAT32"), ",")
	var FORCE_S_FLOAT32 []float32
	for _, v := range FORCE_S_FLOAT32__A {
		i, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return err
		}
		FORCE_S_FLOAT32 = append(FORCE_S_FLOAT32, float32(i))
	}
	FORCE_S_FLOAT64__A := strings.Split(os.Getenv("FORCE_S_FLOAT64"), ",")
	var FORCE_S_FLOAT64 []float64
	for _, v := range FORCE_S_FLOAT64__A {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		FORCE_S_FLOAT64 = append(FORCE_S_FLOAT64, float64(i))
	}
	FORCE_S_INT__A := strings.Split(os.Getenv("FORCE_S_INT"), ",")
	var FORCE_S_INT []int
	for _, v := range FORCE_S_INT__A {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		FORCE_S_INT = append(FORCE_S_INT, i)
	}
	FORCE_S_INT16__A := strings.Split(os.Getenv("FORCE_S_INT16"), ",")
	var FORCE_S_INT16 []int16
	for _, v := range FORCE_S_INT16__A {
		i, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return err
		}
		FORCE_S_INT16 = append(FORCE_S_INT16, int16(i))
	}
	FORCE_S_INT32__A := strings.Split(os.Getenv("FORCE_S_INT32"), ",")
	var FORCE_S_INT32 []int32
	for _, v := range FORCE_S_INT32__A {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		FORCE_S_INT32 = append(FORCE_S_INT32, int32(i))
	}
	FORCE_S_INT64__A := strings.Split(os.Getenv("FORCE_S_INT64"), ",")
	var FORCE_S_INT64 []int64
	for _, v := range FORCE_S_INT64__A {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		FORCE_S_INT64 = append(FORCE_S_INT64, int64(i))
	}
	FORCE_S_INT8__A := strings.Split(os.Getenv("FORCE_S_INT8"), ",")
	var FORCE_S_INT8 []int8
	for _, v := range FORCE_S_INT8__A {
		i, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return err
		}
		FORCE_S_INT8 = append(FORCE_S_INT8, int8(i))
	}
	FORCE_S_INTERFACE__A := strings.Split(os.Getenv("FORCE_S_INTERFACE"), ",")
	var FORCE_S_INTERFACE []interface{}
	for _, v := range FORCE_S_INTERFACE__A {
		FORCE_S_INTERFACE = append(FORCE_S_INTERFACE, v)
	}
	FORCE_S_UINT16__A := strings.Split(os.Getenv("FORCE_S_UINT16"), ",")
	var FORCE_S_UINT16 []uint16
	for _, v := range FORCE_S_UINT16__A {
		i, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return err
		}
		FORCE_S_UINT16 = append(FORCE_S_UINT16, uint16(i))
	}
	FORCE_S_UINT32__A := strings.Split(os.Getenv("FORCE_S_UINT32"), ",")
	var FORCE_S_UINT32 []uint32
	for _, v := range FORCE_S_UINT32__A {
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return err
		}
		FORCE_S_UINT32 = append(FORCE_S_UINT32, uint32(i))
	}
	FORCE_S_UINT64__A := strings.Split(os.Getenv("FORCE_S_UINT64"), ",")
	var FORCE_S_UINT64 []uint64
	for _, v := range FORCE_S_UINT64__A {
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		FORCE_S_UINT64 = append(FORCE_S_UINT64, uint64(i))
	}
	FORCE_S_UINT8__A := strings.Split(os.Getenv("FORCE_S_UINT8"), ",")
	var FORCE_S_UINT8 []uint8
	for _, v := range FORCE_S_UINT8__A {
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return err
		}
		FORCE_S_UINT8 = append(FORCE_S_UINT8, uint8(i))
	}
	FORCE_UINT16__S := os.Getenv("FORCE_UINT16")
	FORCE_UINT16__64, err := strconv.ParseUint(FORCE_UINT16__S, 10, 16)
	if err != nil {
		return err
	}
	FORCE_UINT16 := uint16(FORCE_UINT16__64)
	FORCE_UINT32__S := os.Getenv("FORCE_UINT32")
	FORCE_UINT32__64, err := strconv.ParseUint(FORCE_UINT32__S, 10, 32)
	if err != nil {
		return err
	}
	FORCE_UINT32 := uint32(FORCE_UINT32__64)
	FORCE_UINT64__S := os.Getenv("FORCE_UINT64")
	FORCE_UINT64__64, err := strconv.ParseUint(FORCE_UINT64__S, 10, 64)
	if err != nil {
		return err
	}
	FORCE_UINT64 := uint64(FORCE_UINT64__64)
	FORCE_UINT8__S := os.Getenv("FORCE_UINT8")
	FORCE_UINT8__64, err := strconv.ParseUint(FORCE_UINT8__S, 10, 8)
	if err != nil {
		return err
	}
	FORCE_UINT8 := uint8(FORCE_UINT8__64)
	INT64_TEST1__S := os.Getenv("INT64_TEST1")
	INT64_TEST1__64, err := strconv.ParseInt(INT64_TEST1__S, 10, 64)
	if err != nil {
		return err
	}
	INT64_TEST1 := int64(INT64_TEST1__64)
	INT64_TEST2__S := os.Getenv("INT64_TEST2")
	INT64_TEST2__64, err := strconv.ParseInt(INT64_TEST2__S, 10, 64)
	if err != nil {
		return err
	}
	INT64_TEST2 := int64(INT64_TEST2__64)
	INT64_TEST3__S := os.Getenv("INT64_TEST3")
	INT64_TEST3__64, err := strconv.ParseInt(INT64_TEST3__S, 10, 64)
	if err != nil {
		return err
	}
	INT64_TEST3 := int64(INT64_TEST3__64)
	INVALID_JSON_TEST := os.Getenv("INVALID_JSON_TEST")
	JSON_TEST__S := os.Getenv("JSON_TEST")
	var JSON_TEST JSON_TEST
	err = json.Unmarshal([]byte(JSON_TEST__S), &JSON_TEST)
	JSON_TEST2__S := os.Getenv("JSON_TEST2")
	var JSON_TEST2 JSON_TEST2
	err = json.Unmarshal([]byte(JSON_TEST2__S), &JSON_TEST2)
	JSON_TEST3__S := os.Getenv("JSON_TEST3")
	var JSON_TEST3 JSON_TEST3
	err = json.Unmarshal([]byte(JSON_TEST3__S), &JSON_TEST3)
	STRING_TEST1 := os.Getenv("STRING_TEST1")
	STRING_TEST2 := os.Getenv("STRING_TEST2")
	STRING_TEST3 := os.Getenv("STRING_TEST3")
	S_BOOL_TEST__A := strings.Split(os.Getenv("S_BOOL_TEST"), ",")
	var S_BOOL_TEST []bool
	for _, v := range S_BOOL_TEST__A {
		if strings.ToLower(v) == "true" {
			S_BOOL_TEST = append(S_BOOL_TEST, true)
		} else {
			S_BOOL_TEST = append(S_BOOL_TEST, false)
		}
	}
	S_FLOAT64_TEST__A := strings.Split(os.Getenv("S_FLOAT64_TEST"), ",")
	var S_FLOAT64_TEST []float64
	for _, v := range S_FLOAT64_TEST__A {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		S_FLOAT64_TEST = append(S_FLOAT64_TEST, float64(i))
	}
	S_INT64_TEST__A := strings.Split(os.Getenv("S_INT64_TEST"), ",")
	var S_INT64_TEST []int64
	for _, v := range S_INT64_TEST__A {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		S_INT64_TEST = append(S_INT64_TEST, int64(i))
	}
	S_STRING_TEST := strings.Split(os.Getenv("S_STRING_TEST"), ",")
	S_STRING_TEST2 := strings.Split(os.Getenv("S_STRING_TEST2"), ",")
	UPPERCASETEST := os.Getenv("UPPERCASETEST")
	_TEST := os.Getenv("_TEST")
	lower_snake_case_test := os.Getenv("lower_snake_case_test")
	lower_snake_case_test_json__S := os.Getenv("lower_snake_case_test_json")
	var lower_snake_case_test_json Lower_snake_case_test_json
	err = json.Unmarshal([]byte(lower_snake_case_test_json__S), &lower_snake_case_test_json)
	lowercasetest := os.Getenv("lowercasetest")
	env = environment{
		BOOL_TEST1:                 BOOL_TEST1,
		BOOL_TEST2:                 BOOL_TEST2,
		BOOL_TEST3:                 BOOL_TEST3,
		CamelCaseTest:              CamelCaseTest,
		FLOAT64_TEST1:              FLOAT64_TEST1,
		FLOAT64_TEST2:              FLOAT64_TEST2,
		FLOAT64_TEST3:              FLOAT64_TEST3,
		FORCE_BOOL:                 FORCE_BOOL,
		FORCE_FLOAT32:              FORCE_FLOAT32,
		FORCE_FLOAT64:              FORCE_FLOAT64,
		FORCE_INT:                  FORCE_INT,
		FORCE_INT16:                FORCE_INT16,
		FORCE_INT32:                FORCE_INT32,
		FORCE_INT64:                FORCE_INT64,
		FORCE_INT8:                 FORCE_INT8,
		FORCE_INTERFACE:            FORCE_INTERFACE,
		FORCE_S_BOOL:               FORCE_S_BOOL,
		FORCE_S_FLOAT32:            FORCE_S_FLOAT32,
		FORCE_S_FLOAT64:            FORCE_S_FLOAT64,
		FORCE_S_INT:                FORCE_S_INT,
		FORCE_S_INT16:              FORCE_S_INT16,
		FORCE_S_INT32:              FORCE_S_INT32,
		FORCE_S_INT64:              FORCE_S_INT64,
		FORCE_S_INT8:               FORCE_S_INT8,
		FORCE_S_INTERFACE:          FORCE_S_INTERFACE,
		FORCE_S_UINT16:             FORCE_S_UINT16,
		FORCE_S_UINT32:             FORCE_S_UINT32,
		FORCE_S_UINT64:             FORCE_S_UINT64,
		FORCE_S_UINT8:              FORCE_S_UINT8,
		FORCE_UINT16:               FORCE_UINT16,
		FORCE_UINT32:               FORCE_UINT32,
		FORCE_UINT64:               FORCE_UINT64,
		FORCE_UINT8:                FORCE_UINT8,
		INT64_TEST1:                INT64_TEST1,
		INT64_TEST2:                INT64_TEST2,
		INT64_TEST3:                INT64_TEST3,
		INVALID_JSON_TEST:          INVALID_JSON_TEST,
		JSON_TEST:                  JSON_TEST,
		JSON_TEST2:                 JSON_TEST2,
		JSON_TEST3:                 JSON_TEST3,
		STRING_TEST1:               STRING_TEST1,
		STRING_TEST2:               STRING_TEST2,
		STRING_TEST3:               STRING_TEST3,
		S_BOOL_TEST:                S_BOOL_TEST,
		S_FLOAT64_TEST:             S_FLOAT64_TEST,
		S_INT64_TEST:               S_INT64_TEST,
		S_STRING_TEST:              S_STRING_TEST,
		S_STRING_TEST2:             S_STRING_TEST2,
		UPPERCASETEST:              UPPERCASETEST,
		_TEST:                      _TEST,
		lower_snake_case_test:      lower_snake_case_test,
		lower_snake_case_test_json: lower_snake_case_test_json,
		lowercasetest:              lowercasetest,
	}
	return err
}

type getterInterface interface {
	BOOL_TEST1() bool
	BOOL_TEST2() bool
	BOOL_TEST3() bool
	CamelCaseTest() string
	FLOAT64_TEST1() float64
	FLOAT64_TEST2() float64
	FLOAT64_TEST3() float64
	FORCE_BOOL() bool
	FORCE_FLOAT32() float32
	FORCE_FLOAT64() float64
	FORCE_INT() int
	FORCE_INT16() int16
	FORCE_INT32() int32
	FORCE_INT64() int64
	FORCE_INT8() int8
	FORCE_INTERFACE() interface{}
	FORCE_S_BOOL() []bool
	FORCE_S_FLOAT32() []float32
	FORCE_S_FLOAT64() []float64
	FORCE_S_INT() []int
	FORCE_S_INT16() []int16
	FORCE_S_INT32() []int32
	FORCE_S_INT64() []int64
	FORCE_S_INT8() []int8
	FORCE_S_INTERFACE() []interface{}
	FORCE_S_UINT16() []uint16
	FORCE_S_UINT32() []uint32
	FORCE_S_UINT64() []uint64
	FORCE_S_UINT8() []uint8
	FORCE_UINT16() uint16
	FORCE_UINT32() uint32
	FORCE_UINT64() uint64
	FORCE_UINT8() uint8
	INT64_TEST1() int64
	INT64_TEST2() int64
	INT64_TEST3() int64
	INVALID_JSON_TEST() string
	JSON_TEST() JSON_TEST
	JSON_TEST2() JSON_TEST2
	JSON_TEST3() JSON_TEST3
	STRING_TEST1() string
	STRING_TEST2() string
	STRING_TEST3() string
	S_BOOL_TEST() []bool
	S_FLOAT64_TEST() []float64
	S_INT64_TEST() []int64
	S_STRING_TEST() []string
	S_STRING_TEST2() []string
	UPPERCASETEST() string
	UNDERLINE_TEST() string
	Lower_snake_case_test() string
	Lower_snake_case_test_json() Lower_snake_case_test_json
	Lowercasetest() string
}
type getter struct {
	getterInterface
}

func Get() getter {
	return getter{}
}

type setterInterface interface {
	BOOL_TEST1() bool
	BOOL_TEST2() bool
	BOOL_TEST3() bool
	CamelCaseTest() string
	FLOAT64_TEST1() float64
	FLOAT64_TEST2() float64
	FLOAT64_TEST3() float64
	FORCE_BOOL() bool
	FORCE_FLOAT32() float32
	FORCE_FLOAT64() float64
	FORCE_INT() int
	FORCE_INT16() int16
	FORCE_INT32() int32
	FORCE_INT64() int64
	FORCE_INT8() int8
	FORCE_INTERFACE() interface{}
	FORCE_S_BOOL() []bool
	FORCE_S_FLOAT32() []float32
	FORCE_S_FLOAT64() []float64
	FORCE_S_INT() []int
	FORCE_S_INT16() []int16
	FORCE_S_INT32() []int32
	FORCE_S_INT64() []int64
	FORCE_S_INT8() []int8
	FORCE_S_INTERFACE() []interface{}
	FORCE_S_UINT16() []uint16
	FORCE_S_UINT32() []uint32
	FORCE_S_UINT64() []uint64
	FORCE_S_UINT8() []uint8
	FORCE_UINT16() uint16
	FORCE_UINT32() uint32
	FORCE_UINT64() uint64
	FORCE_UINT8() uint8
	INT64_TEST1() int64
	INT64_TEST2() int64
	INT64_TEST3() int64
	INVALID_JSON_TEST() string
	JSON_TEST() JSON_TEST
	JSON_TEST2() JSON_TEST2
	JSON_TEST3() JSON_TEST3
	STRING_TEST1() string
	STRING_TEST2() string
	STRING_TEST3() string
	S_BOOL_TEST() []bool
	S_FLOAT64_TEST() []float64
	S_INT64_TEST() []int64
	S_STRING_TEST() []string
	S_STRING_TEST2() []string
	UPPERCASETEST() string
	UNDERLINE_TEST() string
	Lower_snake_case_test() string
	Lower_snake_case_test_json() Lower_snake_case_test_json
	Lowercasetest() string
}
type setter struct {
	setterInterface
}

func Set() setter {
	return setter{}
}
