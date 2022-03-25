package metricstruct

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus"
)

func Register(store prometheus.Registerer, str interface{}) error {
	strVal := reflect.ValueOf(str)
	if strVal.Kind() != reflect.Ptr {
		return errors.New("metricstruct.Register: argument must be a pointer to struct")
	}

	strElements := strVal.Elem()
	if strElements.Kind() != reflect.Struct {
		return errors.New("metricstruct.Register: argument must be a pointer to struct")
	}

	for i := 0; i < strElements.NumField(); i++ {
		field := strElements.Field(i)
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		if c, ok := field.Interface().(prometheus.Collector); ok {
			store.MustRegister(c)
		}
	}
	return nil
}
