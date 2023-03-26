package module

import (
	"fmt"
	"reflect"
	"strconv"
)

// Attempt to find the string type from the file presented
func (Opt *Options) String(p ...string) string {
	item, err := Opt.Get(reflect.String, p...)
	if err != nil {
		return ""
	}

	return fmt.Sprint(item)
}

// Attempt to find the int type from the file presented
func (Opt *Options) Integer(p ...string) int {
	var vectors []reflect.Kind = []reflect.Kind{
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	}

	collection, err := Opt.GetFromVectors(vectors, p...)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	period, err := strconv.Atoi(fmt.Sprint(collection))
	if err != nil {
		return 0
	}

	return period
}

// Attempt to find the boolean type from the file presented
func (options *Options) Boolean(p ...string) bool {
	item, err := options.Get(reflect.Bool, p...)
	if err != nil {
		return false
	}

	return item.(bool)
}

// Attempt to find the array type from the file presented
func (options *Options) Array(p ...string) []any {
	item, err := options.Get(reflect.Array, p...)
	if err != nil {
		return make([]any, 0)
	}

	return item.([]any)
}
