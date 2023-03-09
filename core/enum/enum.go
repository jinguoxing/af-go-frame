package enum

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

//IntegerType 整形数据 int8，已经满足很多场景的枚举需求了，其他的数据类型纯属为了兼容
type IntegerType int8

func (ip IntegerType) Int8() int8   { return int8(ip) }
func (ip IntegerType) Int16() int16 { return int16(ip) }
func (ip IntegerType) Int32() int32 { return int32(ip) }
func (ip IntegerType) Int() int     { return int(ip) }

//下面几种方法不支持有符号数据，有符号数据会转成无符号的数据类型

func (ip IntegerType) Uint8() uint8   { return uint8(ip) }
func (ip IntegerType) Uint16() uint16 { return uint16(ip) }
func (ip IntegerType) Uint32() uint32 { return uint32(ip) }
func (ip IntegerType) Uint() uint     { return uint(ip) }

type recordPair struct {
	//stringEnumRecord  string to int reflection
	stringEnumRecord map[string]IntegerType
	//intEnumRecord  int to string reflection
	intEnumRecord map[IntegerType]string
}

//recordPairs stringEnumRecord and intEnumRecord pair, key is a empty  struct
type recordPairs map[any]recordPair

//allRecords  record all the EnumRecord  pair
var allRecords recordPairs = make(map[any]recordPair)

//Object one enum object
type Object struct {
	Integer *IntegerType
	String  string
}

//enumProperty enum object record
type enumProperty interface {
	string | ~uint | ~uint32 | ~uint16 | ~uint8 | ~int | ~int32 | ~int16 | ~int8
}

//enumInteger enum integer type
type enumInteger interface {
	~uint | ~uint32 | ~uint16 | ~uint8 | ~int | ~int32 | ~int16 | ~int8
}

//set record the enum
func set[T any](e *T) {
	key := *new(T)
	maps, ok := allRecords[key]
	if !ok {
		maps = recordPair{
			stringEnumRecord: make(map[string]IntegerType),
			intEnumRecord:    make(map[IntegerType]string),
		}
		allRecords[key] = maps
	}
	a := (*Object)(unsafe.Pointer(e))
	maps.stringEnumRecord[a.String] = *a.Integer
	maps.intEnumRecord[*a.Integer] = a.String
}

//load fix the Object from recordPairs
func load[T any](e *T) {
	maps, ok := allRecords[*new(T)]
	if !ok {
		log.Panicf("invalid enum struct %v", reflect.TypeOf(new(T)))
	}
	a := (*Object)(unsafe.Pointer(e))
	if a.String == "" {
		a.String = maps.intEnumRecord[*a.Integer]
	}
	if a.Integer == nil {
		v := maps.stringEnumRecord[a.String]
		a.Integer = &v
	}
}

//ToString  find the string value of 'i'
func ToString[T any, V enumInteger](i V) string {
	maps, ok := allRecords[*new(T)]
	if !ok {
		log.Panicf("invalid enum struct %v", reflect.TypeOf(new(T)))
	}
	iv := IntegerType(i)
	return maps.intEnumRecord[iv]
}

//ToInteger find the uint8 value of 's'
func ToInteger[T any](s string) IntegerType {
	maps, ok := allRecords[*new(T)]
	if !ok {
		log.Panicf("invalid enum struct %v", reflect.TypeOf(new(T)))
	}
	return maps.stringEnumRecord[s]
}

//New generate EnumClass
func New[T any](i IntegerType, s string) *T {
	e := Object{
		Integer: &i,
		String:  s,
	}
	p := (*T)(unsafe.Pointer(&e))
	set[T](p)
	return p
}

//Is  check whether value v is a valid enum value
func Is[T any, P enumProperty](v P) bool {
	eObj := new(T)
	maps, ok := allRecords[*eObj]
	if !ok {
		log.Panicf("invalid enum struct %v", reflect.TypeOf(new(T)))
	}
	var ev interface{}
	ev = v
	switch ev.(type) {
	case string:
		_, ok := maps.stringEnumRecord[ev.(string)]
		return ok
	default:
		pv, _ := strconv.Atoi(fmt.Sprintf("%v", ev))
		_, ok := maps.intEnumRecord[(IntegerType)(pv)]
		return ok
	}
	return false
}
