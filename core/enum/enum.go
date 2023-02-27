package enum

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

//stringEnumRecord  string to int reflection
type stringEnumRecord map[string]uint8

//intEnumRecord  int to string reflection
type intEnumRecord map[uint8]string

//recordPairs stringEnumRecord and intEnumRecord pair, key is a empty  struct
type recordPairs map[any][]any

//allRecords  record all the EnumRecord  pair
var allRecords = make(recordPairs)

//MapObject one enum object
type MapObject struct {
	Integer uint8
	String  string
}

//MapProperty enum object record
type MapProperty interface {
	string | ~uint8
}

//set record the enum
func set[T any](e *T) {
	key := *new(T)
	maps, ok := allRecords[key]
	if !ok {
		maps = make([]any, 2)
		maps[0] = make(stringEnumRecord)
		maps[1] = make(intEnumRecord)
		allRecords[key] = maps
	}
	a := (*MapObject)(unsafe.Pointer(e))
	(maps[0].(stringEnumRecord))[a.String] = a.Integer
	(maps[1].(intEnumRecord))[a.Integer] = a.String
}

//load fix the EnumObject from recordPairs
func load[T any](e *T) {
	key := *new(T)
	maps, ok := allRecords[key]
	if !ok {
		return
	}
	a := (*MapObject)(unsafe.Pointer(e))
	if a.String == "" {
		a.String = (maps[1].(intEnumRecord))[a.Integer]
	}
	if a.Integer < 0 {
		a.Integer = (maps[0].(stringEnumRecord))[a.String]
	}
}

//ToString  find the string value of 'i'
func ToString[T any](i uint8) string {
	key := *new(T)
	maps, ok := allRecords[key]
	if !ok {
		log.Panicf("invalid enum struct %v", reflect.TypeOf(new(T)))
	}
	return (maps[1].(intEnumRecord))[i]
}

//ToInteger find the uint8 value of 's'
func ToInteger[T any](s string) uint8 {
	key := *new(T)
	maps, ok := allRecords[key]
	if !ok {
		log.Panicf("invalid enum struct %v", reflect.TypeOf(new(T)))
	}
	return (maps[0].(stringEnumRecord))[s]
}

//New generate EnumClass
func New[T any](i uint8, s string) *T {
	e := MapObject{
		Integer: i,
		String:  s,
	}
	p := (*T)(unsafe.Pointer(&e))
	set(p)
	return p
}

// Get return EnumObject by value 'ev'
func Get[T any, P MapProperty](ev P) *T {
	p := new(T)
	e := (*MapObject)(unsafe.Pointer(p))

	var evi interface{}
	evi = ev
	switch evi.(type) {
	case string:
		e.String = fmt.Sprintf("%v", ev)
	default:
		e.Integer = *(*uint8)(unsafe.Pointer(&ev))
	}
	load(p)
	return p
}
