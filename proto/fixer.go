package proto

import (
	"github.com/golang/protobuf/ptypes"
	"reflect"
	"fmt"
	"strings"
	"github.com/google/uuid"
)

func Fix(v interface{}) error {
	// Get the type and value of the argument we were passed.
	ptyp := reflect.TypeOf(v)
	pval := reflect.ValueOf(v)

	var typ reflect.Type
	var val reflect.Value

	if ptyp.Kind() == reflect.Ptr {
		typ = ptyp.Elem()
		val = pval.Elem()
	} else {
		typ = ptyp
		val = pval
	}

	// Make sure we now have a struct
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("can not reflect on none struct %v", v)
	}

	for i := 0; i < typ.NumField(); i++ {
		// Get the type of the field from the type of the struct. For a struct,
		// you always get a StructField.
		sfld := typ.Field(i)

		// Get the type of the StructField, which is the type actually stored
		// in that field of the struct.
		tfld := sfld.Type

		// Get the Kind of that type, which will be the underlying base type
		// used to define the type in question.
		kind := tfld.Kind()

		// Get the value of the field from the value of the struct.
		vfld := val.Field(i)

		// automatically assign UUID if it is empty
		if kind == reflect.String && vfld.CanSet() && sfld.Name == "ID" && strings.Trim(vfld.String(), " ") == "" {
			vfld.SetString(uuid.New().String())
		}

		// automatically assign updated at
		if kind == reflect.Ptr && vfld.CanSet() && sfld.Name == "UpdatedAt" {
			vfld.Set(reflect.ValueOf(ptypes.TimestampNow()))
		}

		// automatically assign created at
		if kind == reflect.Ptr && vfld.CanSet() && sfld.Name == "CreatedAt" && vfld.IsNil() {
			vfld.Set(reflect.ValueOf(ptypes.TimestampNow()))
		}
	}

	return nil
}
